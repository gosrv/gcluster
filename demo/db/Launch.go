package main

import (
	"github.com/gosrv/gcluster/demo/db/logic"
	"github.com/gosrv/gcluster/gbase/app"
	"github.com/gosrv/gcluster/gbase/gdb/gleveldb"
	"github.com/gosrv/gcluster/gbase/gdb/gmongo"
	"github.com/gosrv/gcluster/gbase/gdb/gredis"
	"github.com/gosrv/goioc"
)

func initServices(beanContainerBuilder gioc.IBeanContainerBuilder) {
	beanContainerBuilder.AddBean(
		// redis数据库配置
		gredis.NewAutoConfigReids("demo.redis", ""),
		// mongo数据库配置
		gmongo.NewAutoConfigMongo("demo.mongo", ""),
		// leveldb数据库配置
		gleveldb.NewAutoConfigLevelDB("demo.leveldb", ""),
		// redis驱动测试逻辑
		logic.NewServiceRedisOptDemo(),
		// mongo驱动测试逻辑
		logic.NewServiceMongoOptDemo(),
		// leveldb驱动测试逻辑
		logic.NewServiceLevelDBOptDemo(),
		// 日志驱动测试逻辑
		logic.NewServiceLogOptDemo(),
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
