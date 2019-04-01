package app

import (
	"github.com/globalsign/mgo/bson"
	"github.com/gosrv/gcluster/gbase/controller"
	"github.com/gosrv/gcluster/gbase/glog"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
	"github.com/urfave/cli"
	"os"
	"reflect"
)

const (
	Shtudown = "shutdown"
)

type IApplication interface {
	TriggerEvent(event string)
	ConnectEvent(event string, processor func())
}

type Application struct {
	shutdown   bool
	eventChan  chan string
	eventFuncs map[string]func()
}

func (this *Application) TriggerEvent(event string) {
	this.eventChan <- event
}

func (this *Application) ConnectEvent(event string, processor func()) {
	_, exist := this.eventFuncs[event]
	if exist {
		glog.Panic("duplicate event connect %v", event)
	}
	this.eventFuncs[event] = processor
}

func NewApplication() *Application {
	app := &Application{
		shutdown:   false,
		eventChan:  make(chan string),
		eventFuncs: make(map[string]func()),
	}
	app.ConnectEvent(Shtudown, func() {
		app.shutdown = true
	})
	return app
}

func (this *Application) InitCli() gioc.IConfigLoader {
	app := cli.NewApp()

	startConfig := ""
	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"start"},
			Usage:   "start server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config, c",
					Usage:       "Load configuration from `FILE`",
					Destination: &startConfig,
				},
			},
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	util.VerifyNoError(err)

	configLoader := gioc.NewConfigLoader()
	err = configLoader.Load(startConfig)
	util.VerifyNoError(err)
	err = configLoader.RawConf().Sync()
	util.VerifyNoError(err)
	cfgs := make([]string, 0)
	configLoader.Config().Get("cfg.files").Scan(&cfgs)
	for _, cfg := range cfgs {
		err = configLoader.Load(cfg)
		util.VerifyNoError(err)
	}
	err = configLoader.RawConf().Sync()
	util.VerifyNoError(err)
	return configLoader
}

func (this *Application) InitBuilder() gioc.IBeanContainerBuilder {
	builder := gioc.NewBeanContainerBuilder()
	return builder
}

func (this *Application) InitBaseBeanBuilder(builder gioc.IBeanContainerBuilder, configLoader gioc.IConfigLoader) {
	builder.AddBean(configLoader)

	beanContainer := builder.GetBeanContainer()
	builder.AddBean(beanContainer)
	builder.AddBean(this)

	// 进程id，每次启动都不一样
	builder.AddNamedBean("app.id", bson.NewObjectId().Hex())
	builder.AddBean(gioc.NewBeanBeanConditionInjector())
	// tag解析器
	builder.AddBean(gioc.NewTagParser())
	// bean tag处理器
	builder.AddBean(gioc.NewBeanTagProcessor(beanContainer))
	// cfg tag处理器
	builder.AddBean(gioc.NewConfigValueTagProcessor(configLoader))
	// log tag处理器
	builder.AddBean(glog.NewAutoConfigLog("pcluster.log", ""))
	// 控制器
	builder.AddBean(controller.NewBeanControlPointGroupMgr())
	// bean init
	builder.AddBean(gioc.NewBeanInitDriver())
	builder.AddBean(NewAutoLoadConfig())
}

func (this *Application) Build(builder gioc.IBeanContainerBuilder) gioc.IBeanContainer {
	builder.Build()
	return builder.GetBeanContainer()
}

func (this *Application) enterEventLoop() {
	for !this.shutdown {
		event := <-this.eventChan
		processor, ok := this.eventFuncs[event]
		if !ok {
			continue
		}
		processor()
	}
}

func (this *Application) Start(container gioc.IBeanContainer) {
	beanInitDriver := container.GetBeanByType(reflect.TypeOf((*gioc.BeanInitDriver)(nil)))[0].(*gioc.BeanInitDriver)

	beanInitDriver.CallInit()
	beanInitDriver.CallStart()

	this.enterEventLoop()

	beanInitDriver.CallStop()
	beanInitDriver.CallUnInit()
}

func (this *Application) Shutdown() {
	this.TriggerEvent(Shtudown)
}
