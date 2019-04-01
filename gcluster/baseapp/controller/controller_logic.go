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
		IController: controller.NewTypeController(""),
	}
}

// 心跳消息处理
func (this *ControllerLogic) Logic(ctx gnet.ISessionCtx, msg *netproto.CS_Tick) *netproto.SC_Tick {
	return &netproto.SC_Tick{}
}
