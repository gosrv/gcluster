package glog

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type LogConfig struct {
	Level         string
	Output        map[string]map[string]string
	TextFormatter *logrus.TextFormatter
	JSONFormatter *logrus.JSONFormatter
}

type LogDriver struct {
	logConfigs map[string]*LogConfig
	loggers    map[string]*logrus.Logger
	lock       sync.Mutex
}

func NewLogDriver(logConfigs map[string]*LogConfig) *LogDriver {
	driver := &LogDriver{
		logConfigs: logConfigs,
		loggers:    make(map[string]*logrus.Logger),
	}

	for name, cfg := range logConfigs {
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
				Panic("no log output config [%v]", logConfig)
			}
			logConfig.ConfigLogOutput(logWriter, outputCfg)
		}
		driver.loggers[name] = logger
	}
	return driver
}

func (this *LogDriver) GetLogger(name string) *logrus.Logger {
	logger, ok := this.loggers[name]
	if ok {
		return logger
	}

	this.lock.Lock()
	defer this.lock.Unlock()

	newLoggers := make(map[string]*logrus.Logger)
	for k, v := range this.loggers {
		newLoggers[k] = v
	}
	logger = logrus.New()
	newLoggers[name] = logger

	this.loggers = newLoggers
	return logger
}
