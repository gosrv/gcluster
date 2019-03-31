package main

import (
	"github.com/gosrv/gcluster/gbase/app"
	"github.com/gosrv/gcluster/gbase/codec"
	"github.com/gosrv/gcluster/gbase/tcpnet"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/testclient/logic"
	"github.com/gosrv/goioc"
)

func initBaseNet(builder gioc.IBeanContainerBuilder) {
	idtype := entity.NewLogicMsgIds()
	encoder := codec.NewNetMsgFixLenProtobufEncoder(idtype)
	decoder := codec.NewNetMsgFixLenProtobufDecoder(idtype)

	net := tcpnet.NewTcpNetServer("pcluster.basenet", "", encoder, decoder, nil, nil)
	builder.AddBean(net)
}

func initServices(beanContainerBuilder gioc.IBeanContainerBuilder) {
	beanContainerBuilder.AddBean(
		logic.NewControllerLogic(),
		logic.NewServiceLogic(),
	)
}

func main() {
	application := app.NewApplication()
	configLoader := application.InitCli()
	builder := application.InitBuilder()
	application.InitBaseBeanBuilder(builder, configLoader)

	initServices(builder)
	initBaseNet(builder)
	beanContainer := application.Build(builder)

	application.Start(beanContainer)
}
