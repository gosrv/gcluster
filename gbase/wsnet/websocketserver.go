package wsnet

import (
	"github.com/gosrv/gcluster/gbase/controller"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/gproto"
	"github.com/gosrv/gcluster/gbase/route"
	"github.com/gosrv/glog"
	"github.com/gosrv/goioc"
	"net/http"
)

const (
	tcpnetConfigValue = "app.tcpnet.host"
)

/**
自动网络配置
*/
type WebsocketServer struct {
	// 启动条件
	gioc.IBeanCondition
	gioc.IConfigBase
	// 从配置文件中注入host
	host    string            `cfg.d:"ws.host"`
	wsentry string            `cfg.d:"ws.entry" cfg.default:"/ws"`
	msgType string            `cfg.d:"ws.msgtype" cfg.default:"string"`
	log     glog.IFieldLogger `log:"engine"`
	// 注入控制器
	controlPointCollector controller.IControlPointGroupMgr `bean`
	eventRoute            gproto.IRoute
	delegateDataRoute     gproto.IRouteDelegate
	encoder               gproto.IEncoder
	decoder               gproto.IDecoder
	ctlGroup              string
	net                   *netSystem
	handler               *http.ServeMux
}

func (this *WebsocketServer) GetEventRoute() gproto.IRoute {
	return this.eventRoute
}

// 启动网络
func (this *WebsocketServer) BeanStart() {
	if len(this.host) > 0 {
		this.net = GoStartWsServer(this.host, this.wsentry, this.msgType,
			this.createNetConfig(), this.handler)
		this.log.Debug("websocket listen on %v", this.host)
	}
}

func (this *WebsocketServer) BeanStop() {

}

func NewWebsocketServer(cfgBase, ctlGroup string, encoder gproto.IEncoder, decoder gproto.IDecoder) *WebsocketServer {
	return &WebsocketServer{
		// 启用条件：配置文件中存在配置tcpnetConfigValue
		IBeanCondition:    gioc.NewConditionOnValue(tcpnetConfigValue, true),
		IConfigBase:       gioc.NewConfigBase(cfgBase),
		encoder:           encoder,
		decoder:           decoder,
		eventRoute:        route.NewRouteMap(false, false),
		delegateDataRoute: route.NewRouteDelegate(false),
		ctlGroup:          ctlGroup,
		handler:           http.NewServeMux(),
	}
}

func (this *WebsocketServer) createNetConfig() *gnet.NetConfig {
	this.delegateDataRoute.SetDelegate(controller.NewControlPointRoute(
		this.controlPointCollector.GetControlPointGroup(this.ctlGroup)))
	return &gnet.NetConfig{
		// 网络消息编码器，4字节长度 + 2字节id + protobuf
		Encoder: this.encoder,
		// 网络消息解码器，4字节长度 + 2字节id + protobuf
		Decoder: this.decoder,
		// 事件路由器
		EventRoute: this.eventRoute,
		// 数据路由器，转发给控制器
		DataRoute:        this.delegateDataRoute,
		ReadBufSize:      16384,
		WriteChannelSize: 1024,
		HeartTickMs:      10000,
	}
}
