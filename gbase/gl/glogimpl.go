package gl

import (
	"github.com/gosrv/glog"
	"github.com/gosrv/goioc/util"
	"log"
)

type gLog struct {
}

func (this *gLog) Panic(format string, v ...interface{}) {
	util.Panic(format, v...)
}

func (this *gLog) Error(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (this *gLog) Warn(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (this *gLog) Info(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (this *gLog) Debug(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (this *gLog) Trace(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (this *gLog) Redirect(logger glog.IFieldLogger) error {
	return redirect(logger)
}

type ggLog struct {
	logger glog.IFieldLogger
}

func (this *ggLog) Panic(format string, v ...interface{}) {
	this.logger.Panic(format, v...)
}

func (this *ggLog) Error(format string, v ...interface{}) {
	this.logger.Error(format, v...)
}

func (this *ggLog) Warn(format string, v ...interface{}) {
	this.logger.Warn(format, v...)
}

func (this *ggLog) Info(format string, v ...interface{}) {
	this.logger.Info(format, v...)
}

func (this *ggLog) Debug(format string, v ...interface{}) {
	this.logger.Debug(format, v...)
}

func (this *ggLog) Trace(format string, v ...interface{}) {
	this.logger.Print(format, v...)
}

func (this *ggLog) Redirect(logger glog.IFieldLogger) error {
	return redirect(logger)
}

func redirect(logger glog.IFieldLogger) error {
	GLog = &ggLog{logger: logger}
	return nil
}
