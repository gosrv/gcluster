package logic

import (
	"github.com/sirupsen/logrus"
)

type serviceLogOptDemo struct {
	logApp    *logrus.Logger `log:"app"`
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
