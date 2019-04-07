package logic

import (
	"github.com/gosrv/glog"
)

type serviceLogOptDemo struct {
	// 日志属性配置文件设置
	// app日志的自动注入
	logApp glog.IFieldLogger `log:"app"`
	// engine日志的自动注入
	logEngine glog.IFieldLogger `log:"engine"`
}

func NewServiceLogOptDemo() *serviceLogOptDemo {
	return &serviceLogOptDemo{}
}

func (this *serviceLogOptDemo) BeanStart() {
	for i := 0; i < 10; i++ {
		this.logApp.WithField("app", "demo").Debug("hello app log")
		this.logEngine.WithField("engine", "demo").Debug("hello engine log")
	}
}

func (this *serviceLogOptDemo) BeanStop() {

}
