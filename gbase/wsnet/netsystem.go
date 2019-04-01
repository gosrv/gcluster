package wsnet

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/gosrv/gcluster/gbase/glog"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/gproto"
	"github.com/gosrv/gcluster/gbase/gutil"
	"github.com/gosrv/goioc/util"
	"net/http"
	"reflect"
)

type netSystem struct {
	host    string
	wsentry string
	msgType int
	config  *gnet.NetConfig
	handler *http.ServeMux
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  32 * 1024,
	WriteBufferSize: 64 * 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 将客户端消息，转发到上游服务器
func (this *netSystem) goReadProcess(netChannel *wsNetChannel) error {
	// 通知上游打开连接
	for netChannel.IsActive() {
		messageType, msg, err := netChannel.conn.ReadMessage()
		util.Verify(messageType == this.msgType)
		if err != nil {
			return errors.New("net channel read closed")
		}
		netMsg := this.config.Decoder.Decode(msg)
		netChannel.msgUnprocessedChannel <- netMsg
	}
	return nil
}

// 将上游服务器的消息发送到客户端
func (this *netSystem) goWriteProcess(netChannel *wsNetChannel) error {
	eventRoute := this.config.EventRoute
	dataRoute := this.config.DataRoute

	eventRoute.Trigger(netChannel.ctx, gnet.NetEventConnect, nil)
	defer func() {
		netChannel.Close()
		for {
			// 把剩下的活干完就可以退出了
			if msg, ok := <-netChannel.msgUnprocessedChannel; ok {
				netChannel.ctx.Clear(gnet.ScopeRequest)
				netChannel.ctx.SetAttribute(gnet.ScopeRequest, reflect.TypeOf(msg), msg)
				_ = dataRoute.Trigger(netChannel.ctx, reflect.TypeOf(msg), msg)
			} else {
				break
			}
		}
		eventRoute.Trigger(netChannel.ctx, gnet.NetEventDisconnect, nil)
	}()

	for netChannel.IsActive() {
		select {
		case _, _ = <-netChannel.closeChannel:
			return nil
		case msg, ok := <-netChannel.msgProcessedChannel:
			{
				if !ok {
					// channel已被关闭
					return errors.New("gnet msg processed channel closed.")
				}
				err := netChannel.conn.WriteMessage(this.msgType, msg.([]byte))
				if err != nil {
					return err
				}
			}
		case msg, ok := <-netChannel.msgUnprocessedChannel:
			{
				if !ok {
					// channel已被关闭
					return errors.New("gnet msg unprocessed channel closed.")
				}
				netChannel.ctx.Clear(gnet.ScopeRequest)
				netChannel.ctx.SetAttribute(gnet.ScopeRequest, reflect.TypeOf(msg), msg)
				response := dataRoute.Trigger(netChannel.ctx, reflect.TypeOf(msg), msg)
				if response == nil || reflect.ValueOf(response).IsNil() {
					continue
				}
				netChannel.msgProcessedChannel <- response
			}
		case _, ok := <-netChannel.heartTicker.C:
			{
				if !ok {
					return errors.New("ticker stoped")
				}
				eventRoute.Trigger(netChannel.ctx, gnet.NetEventTick, nil)
			}
		}
	}
	return nil
}

func (this *netSystem) wsStart(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		glog.Warn("ws upgrade error %v", err)
		return
	}
	netChannel := NewWsNetChannel(conn, this.config.WriteChannelSize,
		this.config.WriteChannelSize, this.config.HeartTickMs)
	netChannel.ctx.SetAttribute(gnet.ScopeSession, gproto.INetChannelType, netChannel)
	netChannel.ctx.SetAttribute(gnet.ScopeSession, gnet.ISessionCtxType, netChannel.ctx)

	gutil.RecoverGo(func() {
		defer netChannel.Close()
		this.goWriteProcess(netChannel)
	})
	gutil.RecoverGo(func() {
		defer netChannel.Close()
		this.goWriteProcess(netChannel)
	})
}

const (
	MsgTypeString = "string"
	MsgTypeBinary = "bin"
)

func GoStartWsServer(host string, wsentry string, msgType string,
	config *gnet.NetConfig, handler *http.ServeMux) *netSystem {
	mt := 0
	switch msgType {
	case MsgTypeString:
		mt = websocket.TextMessage
	case MsgTypeBinary:
		mt = websocket.BinaryMessage
	default:
		glog.Panic("unknown websocket msg type")
		return nil
	}
	net := &netSystem{
		host:    host,
		wsentry: wsentry,
		msgType: mt,
		config:  config,
		handler: handler,
	}
	gutil.RecoverGo(func() {
		handler.HandleFunc(wsentry, func(writer http.ResponseWriter, request *http.Request) {
			net.wsStart(writer, request)
		})
		err := http.ListenAndServe(host, handler)
		if err != nil {
			glog.Panic("start websocket error %v", err)
		}
	})
	return nil
}
