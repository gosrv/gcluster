package glog

import (
	"errors"
	"fmt"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
	"github.com/sirupsen/logrus"
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

func (this *gLog) Redirect(base, name string, loader gioc.IConfigLoader) error {
	return redirect(base, name, loader)
}

type logrusLog struct {
	logger *logrus.Logger
}

func (this *logrusLog) Panic(format string, v ...interface{}) {
	this.logger.Panicf(format, v...)
}

func (this *logrusLog) Error(format string, v ...interface{}) {
	this.logger.Errorf(format, v...)
}

func (this *logrusLog) Warn(format string, v ...interface{}) {
	this.logger.Warnf(format, v...)
}

func (this *logrusLog) Info(format string, v ...interface{}) {
	this.logger.Infof(format, v...)
}

func (this *logrusLog) Debug(format string, v ...interface{}) {
	this.logger.Debugf(format, v...)
}

func (this *logrusLog) Trace(format string, v ...interface{}) {
	this.logger.Printf(format, v...)
}

func (this *logrusLog) Redirect(base, name string, loader gioc.IConfigLoader) error {
	return redirect(base, name, loader)
}

func redirect(base, name string, loader gioc.IConfigLoader) error {
	cfg := &LogConfig{}
	err := loader.Config().Get(base, name).Scan(cfg)
	if err != nil {
		return err
	}

	logger := logrus.New()
	logWriter := NewLogOutputWriter()
	logger.SetOutput(logWriter)
	if cfg.TextFormatter != nil {
		logger.SetFormatter(cfg.TextFormatter)
	} else if cfg.JSONFormatter != nil {
		logger.SetFormatter(cfg.JSONFormatter)
	}

	lev, err := logrus.ParseLevel(cfg.Level)
	if err == nil {
		logger.SetLevel(lev)
	}
	for outputName, outputCfg := range cfg.Output {
		logConfig, ok := LogOutputConfigs[outputName]
		if !ok {
			return errors.New(fmt.Sprintf("no log output config [%v]", logConfig))
		}
		logConfig.ConfigLogOutput(logWriter, outputCfg)
	}
	GLog = &logrusLog{logger: logger}
	return nil
}
