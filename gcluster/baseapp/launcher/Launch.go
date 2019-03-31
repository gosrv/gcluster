package main

import (
	"github.com/gosrv/gcluster/gbase/app"
	"github.com/gosrv/gcluster/gbase/codec"
	"github.com/gosrv/gcluster/gbase/gdb/gmongo"
	"github.com/gosrv/gcluster/gbase/gdb/gredis"
	"github.com/gosrv/gcluster/gbase/ghttp"
	"github.com/gosrv/gcluster/gbase/tcpnet"
	"github.com/gosrv/gcluster/gcluster/baseapp/controller"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/baseapp/service"
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/goioc"
)

func initBaseNet(builder gioc.IBeanContainerBuilder) {
	idtype := entity.NewLogicMsgIds()
	encoder := codec.NewNetMsgFixLenProtobufEncoder(idtype)
	decoder := codec.NewNetMsgFixLenProtobufDecoder(idtype)

	net := tcpnet.NewTcpNetServer("pcluster.basenet", "",
		encoder, decoder, nil, entity.NewAutoSyncDataRoute())
	builder.AddBean(net)
}

func initClusterMsgCenter(builder gioc.IBeanContainerBuilder) {
	idtype := entity.NewLogicMsgIds()
	encoder := codec.NewIdProtobufEncoder(idtype)
	decoder := codec.NewIdProtobufDecoder(idtype)
	builder.AddBean(common.NewClusterMsgCenter(encoder, decoder))
}

func initServices(builder gioc.IBeanContainerBuilder) {
	builder.AddBean(
		// redis 自动配置
		gredis.NewAutoConfigReids("pcluster.redis", ""),
		// mongo 自动配置
		gmongo.NewAutoConfigMongo("pcluster.mongo", ""),
		ghttp.NewHttpServer("pcluster.http", nil),
		common.NewClusterNodeMgr(),

		controller.NewControllerLogin(),
		controller.NewControllerLogic(),
		service.NewPlayerMgr(),
		service.NewServiceDataAutoSync(),
		service.NewServiceLogic(),
		service.NewServiceLogin(),
		service.NewServicePlayerMsgQueue(),
	)
}

func main() {
	application := app.NewApplication()
	configLoader := application.InitCli()
	builder := application.InitBuilder()
	application.InitBaseBeanBuilder(builder, configLoader)

	initServices(builder)
	initBaseNet(builder)
	initClusterMsgCenter(builder)
	beanContainer := application.Build(builder)

	application.Start(beanContainer)
}
