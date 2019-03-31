package main

import (
	"github.com/gosrv/gcluster/demo/http/ctl"
	"github.com/gosrv/gcluster/gbase/app"
	"github.com/gosrv/gcluster/gbase/ghttp"
	"github.com/gosrv/goioc"
)

func initServices(beanContainerBuilder gioc.IBeanContainerBuilder) {
	beanContainerBuilder.AddBean(
		// http服务启动，他会从配置文件的配置选项"demo.http"中加载配置数据
		ghttp.NewHttpServer("demo.http", nil),
		// http逻辑处理bean
		ctl.NewControllerHttp(),
	)
}

func main() {
	// 基础配置，大部分情况下可以保持不变
	application := app.NewApplication()
	// 命令行处理和配置加载，我们需要通过命令行指定一个配置文件
	// 大部分bean需要从配置文件读取配置数据，所以必须指定一个配置文件
	// start -c demo/conf/http.json
	configLoader := application.InitCli()
	// 创建bean容器
	builder := application.InitBuilder()
	// 添加基础服务bean
	application.InitBaseBeanBuilder(builder, configLoader)
	// 添加本应用专属bean
	initServices(builder)
	// 装配bean
	beanContainer := application.Build(builder)
	// 开始死循环
	application.Start(beanContainer)
}
