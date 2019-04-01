package main

import (
	"github.com/gosrv/gcluster/gbase/app"
	"github.com/gosrv/gcluster/gbase/gdb/gmongo"
	"github.com/gosrv/gcluster/gbase/gdb/gredis"
	"github.com/gosrv/gcluster/gbase/ghttp"
	"github.com/gosrv/gcluster/gbase/glog"
	"github.com/gosrv/gcluster/gcluster/loginapp/logic"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
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
	// 重定向系统日志
	err := glog.Redirect("pcluster.log", "engine", configLoader)
	util.VerifyNoError(err)
	glog.Debug("application init...")

	builder := application.InitBuilder()
	application.InitBaseBeanBuilder(builder, configLoader)

	initServices(builder)
	beanContainer := application.Build(builder)

	glog.Debug("application start...")
	application.Start(beanContainer)
	glog.Debug("application finished...")
}
