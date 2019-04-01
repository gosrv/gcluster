package glog

import "github.com/gosrv/goioc"

type IGLog interface {
	Panic(format string, v ...interface{})
	Error(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Trace(format string, v ...interface{})
	Redirect(base, name string, loader gioc.IConfigLoader) error
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

func Redirect(base, name string, loader gioc.IConfigLoader) error {
	return GLog.Redirect(base, name, loader)
}
