package main

import (
	"github.com/gosrv/gcluster/demo/http/ctl"
	"github.com/gosrv/gcluster/gbase/app"
	"github.com/gosrv/gcluster/gbase/ghttp"
	"github.com/gosrv/goioc"
)

func initServices(beanContainerBuilder gioc.IBeanContainerBuilder) {
	beanContainerBuilder.AddBean(
		// http 自动配置
		ghttp.NewHttpServer("demo.http", nil),
		ctl.NewControllerHttp(),
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
