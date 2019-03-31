package logic

import (
	"github.com/sirupsen/logrus"
)

type serviceLogOptDemo struct {
	// 日志属性配置文件设置
	// app日志的自动注入
	logApp *logrus.Logger `log:"app"`
	// engine日志的自动注入
	logEngine *logrus.Logger `log:"engine"`
}

func NewServiceLogOptDemo() *serviceLogOptDemo {
	return &serviceLogOptDemo{}
}

func (this *serviceLogOptDemo) BeanStart() {
	for i := 0; i < 10; i++ {
		this.logApp.WithField("app", "demo").Debugln("hello app log")
		this.logEngine.WithField("engine", "demo").Debugln("hello engine log")
	}
}

func (this *serviceLogOptDemo) BeanStop() {

}
