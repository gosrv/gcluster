package glog

import (
	"github.com/gosrv/goioc/util"
	"log"
)

type IGLog interface {
	Panic(format string, v ...interface{})
	Error(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Trace(format string, v ...interface{})
}

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

var GLog IGLog = &gLog{}

func Panic(format string, v ...interface{}) {
	GLog.Panic(format, v...)
}

func Error(format string, v ...interface{}) {
	GLog.Error(format, v...)
}

func Warn(format string, v ...interface{}) {
	GLog.Warn(format, v...)
}

func Info(format string, v ...interface{}) {
	GLog.Info(format, v...)
}

func Debug(format string, v ...interface{}) {
	GLog.Debug(format, v...)
}

func Trace(format string, v ...interface{}) {
	GLog.Trace(format, v...)
}
