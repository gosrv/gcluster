package wsnet

import (
	"github.com/gorilla/websocket"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/gproto"
	"net"
	"sync"
	"time"
)

type wsNetChannel struct {
	encoder gproto.IEncoder
	conn    *websocket.Conn
	closed  bool
	// 已经处理过的消息，需要通过网络发出去
	msgProcessedChannel chan interface{}
	// 未处理消息，需要通过路由器进行处理
	msgUnprocessedChannel chan interface{}
	heartTicker           *time.Ticker
	closeChannel          chan int
	ctx                   gnet.ISessionCtx
	lock                  sync.Mutex
}

var _ gproto.INetChannel = (*wsNetChannel)(nil)

func NewWsNetChannel(conn *websocket.Conn,
	msgProcessedChannelSize int,
	msgUnprocessedChannelSize int,
	heartTickMs int,
) *wsNetChannel {
	return &wsNetChannel{
		conn:                  conn,
		msgProcessedChannel:   make(chan interface{}, msgProcessedChannelSize),
		msgUnprocessedChannel: make(chan interface{}, msgUnprocessedChannelSize),
		closed:                false,
		ctx:                   gnet.NewSessionCtx(),
		heartTicker:           time.NewTicker(time.Duration(heartTickMs) * time.Millisecond),
		closeChannel:          make(chan int, 1),
	}
}

func (this *wsNetChannel) IsActive() bool {
	return !this.closed
}

func (this *wsNetChannel) RemoteAddr() net.Addr {
	return this.conn.RemoteAddr()
}

func (this *wsNetChannel) LocalAddr() net.Addr {
	return this.conn.LocalAddr()
}

func (this *wsNetChannel) ProcessMsg(msg interface{}) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	this.msgUnprocessedChannel <- msg
	return false
}

func (this *wsNetChannel) Send(data interface{}) {
	if this.encoder != nil {
		this.msgProcessedChannel <- this.encoder.Encode(data)
	} else {
		this.msgProcessedChannel <- data
	}
}

func (this *wsNetChannel) Close() error {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.closed {
		return nil
	}
	this.closed = true
	this.heartTicker.Stop()
	close(this.closeChannel)
	close(this.msgUnprocessedChannel)
	return this.conn.Close()
}

func (this *wsNetChannel) SetEncoder(encoder gproto.IEncoder) {
	this.encoder = encoder
}
