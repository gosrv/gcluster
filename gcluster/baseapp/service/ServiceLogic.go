package service

import (
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gbase/gproto"
	"github.com/gosrv/gcluster/gbase/tcpnet"
	"github.com/sirupsen/logrus"
)

type ServiceLogic struct {
	log *logrus.Logger       `log:"app"`
	net *tcpnet.TcpNetServer `bean:""`
}

func (this *ServiceLogic) BeanInit() {
	eventRoute := this.net.GetEventRoute()
	eventRoute.Connect(gnet.NetEventConnect, func(from interface{}, key interface{}, data interface{}) interface{} {
		ctx := from.(gnet.ISessionCtx)
		netChannel := ctx.Get(gproto.INetChannelType).(gproto.INetChannel)
		this.log.Debugf("net connect event %v->%v", netChannel.RemoteAddr(), netChannel.LocalAddr())
		return nil
	})
	eventRoute.Connect(gnet.NetEventDisconnect, func(from interface{}, key interface{}, data interface{}) interface{} {
		ctx := from.(gnet.ISessionCtx)
		netChannel := ctx.Get(gproto.INetChannelType).(gproto.INetChannel)
		this.log.Debugf("net disconnect event %v->%v", netChannel.RemoteAddr(), netChannel.LocalAddr())
		return nil
	})
}

func (this *ServiceLogic) BeanUninit() {
}

func NewServiceLogic() *ServiceLogic {
	return &ServiceLogic{}
}
