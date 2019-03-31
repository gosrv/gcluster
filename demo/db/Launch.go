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
		// redis 自动配置
		gredis.NewAutoConfigReids("demo.redis", ""),
		// mongo 自动配置
		gmongo.NewAutoConfigMongo("demo.mongo", ""),
		// leveldb 自动配置
		gleveldb.NewAutoConfigLevelDB("demo.leveldb", ""),
		// 测试
		logic.NewServiceRedisOptDemo(),
		logic.NewServiceMongoOptDemo(),
		logic.NewServiceLevelDBOptDemo(),
		logic.NewServiceLogOptDemo(),
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
