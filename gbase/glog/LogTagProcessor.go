package glog

import (
	"reflect"
)

const (
	LogTag    = "log"
	LogDomain = "log.domain"
)

type LogTagProcessor struct {
	domain string
	driver *LogDriver
}

func NewLogTagProcessor(domain string, driver *LogDriver) *LogTagProcessor {
	return &LogTagProcessor{
		domain: domain,
		driver: driver,
	}
}

func (this *LogTagProcessor) TagProcessorName() string {
	return "log"
}

func (this *LogTagProcessor) TagProcess(bean interface{}, field reflect.Value, tags map[string]string) {
	logName, logOk := tags[LogTag]
	if !logOk {
		return
	}

	if tags[LogDomain] != this.domain {
		return
	}

	field.Set(reflect.ValueOf(this.driver.GetLogger(logName)))
}
