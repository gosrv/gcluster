package controller

import (
	"github.com/gosrv/gcluster/gbase/controller"
	"github.com/gosrv/gcluster/gbase/gnet"
	"github.com/gosrv/gcluster/gcluster/baseapp/service"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/sirupsen/logrus"
)

/**
逻辑消息控制器
*/
type ControllerLogic struct {
	log *logrus.Logger `log:"app"`
	// 控制器标记
	controller.IController
	// 逻辑处理
	serviceLogic *service.ServiceLogic `bean:""`
}

func NewControllerLogic() *ControllerLogic {
	return &ControllerLogic{
		// 路由收集器，它会收集这样的函数作为路由器：
		// 第一个变量是gnet.ISessionCtx，第二个是消息，可以返回一个一个消息，也可以不返回
		IController: controller.NewTypeController(""),
	}
}

// 心跳消息处理
func (this *ControllerLogic) Logic(ctx gnet.ISessionCtx, msg *netproto.CS_Tick) *netproto.SC_Tick {
	return &netproto.SC_Tick{}
}
