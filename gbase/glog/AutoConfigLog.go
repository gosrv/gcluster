package glog

import (
	"github.com/gosrv/goioc"
	"reflect"
)

type AutoConfigLog struct {
	gioc.IBeanCondition
	gioc.IConfigBase
	tagProcessor *LogTagProcessor
	logConfigs   map[string]*LogConfig `cfg.d:""`
	*LogDriver
	domain string
}

var _ gioc.ITagProcessor = (*AutoConfigLog)(nil)

func NewAutoConfigLog(cfgBase, domain string) *AutoConfigLog {
	return &AutoConfigLog{
		IBeanCondition: gioc.NewConditionOnValue(cfgBase, true),
		IConfigBase:    gioc.NewConfigBase(cfgBase),
		domain:         domain,
	}
}
func (this *AutoConfigLog) BeanBeforeTagProcess(tagProcessor gioc.ITagProcessor, beanContainer gioc.IBeanContainer) {

}
func (this *AutoConfigLog) BeanAfterTagProcess(tagProcessor gioc.ITagProcessor, beanContainer gioc.IBeanContainer) {
	if tagProcessor.TagProcessorName() != gioc.ConfigTagProcessor {
		return
	}
	this.LogDriver = NewLogDriver(this.logConfigs)
	this.tagProcessor = NewLogTagProcessor(this.domain, this.LogDriver)
}

func (this *AutoConfigLog) TagProcessorName() string {
	return this.tagProcessor.TagProcessorName()
}

func (this *AutoConfigLog) TagProcess(bean interface{}, field reflect.Value, tags map[string]string) {
	this.tagProcessor.TagProcess(bean, field, tags)
}
