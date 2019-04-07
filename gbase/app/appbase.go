package app

import (
	"github.com/globalsign/mgo/bson"
	"github.com/gosrv/gcluster/gbase/controller"
	"github.com/gosrv/gcluster/gbase/gl"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
	"github.com/urfave/cli"
	"os"
	"reflect"
)

const (
	Shtudown = "shutdown"
)

type appEvent struct {
	name string
	data interface{}
}

func NewAppEvent(event string, data interface{}) *appEvent {
	return &appEvent{name: event, data: data}
}

// 应用程序
type IApplication interface {
	// 是否已关闭
	IsShutdown() bool
	// 触发事件
	TriggerEvent(event string, data interface{})
	// 设置事件处理函数，非线程安全，需要在bean初始化期间全部设置完成
	ConnectEvent(event string, processor func(string, interface{}))
}

type Application struct {
	shutdown   bool
	eventChan  chan *appEvent
	eventFuncs map[string][]func(string, interface{})
}

func (this *Application) IsShutdown() bool {
	return this.shutdown
}

func (this *Application) TriggerEvent(event string, data interface{}) {
	this.eventChan <- NewAppEvent(event, data)
}

func (this *Application) ConnectEvent(event string, processor func(string, interface{})) {
	this.eventFuncs[event] = append(this.eventFuncs[event], processor)
}

func NewApplication() *Application {
	app := &Application{
		shutdown:   false,
		eventChan:  make(chan *appEvent),
		eventFuncs: make(map[string][]func(string, interface{})),
	}
	app.ConnectEvent(Shtudown, func(string, interface{}) {
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
	builder.AddBean(gl.NewAutoConfigLog("pcluster.log", ""))
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
		processors, ok := this.eventFuncs[event.name]
		if !ok {
			continue
		}
		for _, pro := range processors {
			pro(event.name, event.data)
		}
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
	this.TriggerEvent(Shtudown, nil)
}
