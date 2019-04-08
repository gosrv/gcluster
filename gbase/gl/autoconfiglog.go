package gl

import (
	"github.com/gosrv/glog"
	"github.com/gosrv/goioc"
	"reflect"
)

type AutoConfigLog struct {
	gioc.IBeanCondition
	gioc.IConfigBase
	tagProcessor *LogTagProcessor
	logConfigs   *glog.ConfigLogRoot `cfg.d:"logger"`
	logFactory   glog.ILogFactory    `bean:""`
	domain       string
}

var _ gioc.ITagProcessor = (*AutoConfigLog)(nil)

func NewAutoConfigLog(cfgBase, domain string) *AutoConfigLog {
	return &AutoConfigLog{
		IBeanCondition: gioc.NewConditionOnValue(cfgBase, true),
		IConfigBase:    gioc.NewConfigBase(cfgBase),
		domain:         domain,
	}
}

func (this *AutoConfigLog) PrepareProcess() {
	this.tagProcessor = NewLogTagProcessor(this.domain, this.logFactory)
}

func (this *AutoConfigLog) TagProcessorName() string {
	return this.tagProcessor.TagProcessorName()
}

func (this *AutoConfigLog) TagProcess(bean interface{}, fType reflect.StructField, fValue reflect.Value, tags map[string]string) {
	this.tagProcessor.TagProcess(bean, fType, fValue, tags)
}
