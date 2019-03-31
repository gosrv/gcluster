package main

import (
	"github.com/gosrv/gcluster/gbase/app"
	"github.com/gosrv/gcluster/gbase/gdb/gmongo"
	"github.com/gosrv/gcluster/gbase/gdb/gredis"
	"github.com/gosrv/gcluster/gbase/ghttp"
	"github.com/gosrv/gcluster/gcluster/loginapp/logic"
	"github.com/gosrv/goioc"
)

func initServices(beanContainerBuilder gioc.IBeanContainerBuilder) {
	beanContainerBuilder.AddBean(
		// redis 自动配置
		gredis.NewAutoConfigReids("pcluster.redis", ""),
		// mongo 自动配置
		gmongo.NewAutoConfigMongo("pcluster.mongo", ""),
		// http 自动配置
		ghttp.NewHttpServer("pcluster.http", nil),
		logic.NewControllerLogin(),
		logic.NewServiceLogin(),
	)
}

func main() {
	application := app.NewApplication()
	configLoader := application.InitCli()
	builder := application.InitBuilder()
	application.InitBaseBeanBuilder(builder, configLoader)

	initServices(builder)
	beanContainer := application.Build(builder)

	application.Start(beanContainer)
}
