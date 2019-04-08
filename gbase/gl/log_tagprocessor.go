package gl

import (
	"github.com/gosrv/glog"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
	"reflect"
)

const (
	LogTag    = "log"
	LogDomain = "log.domain"
)

type LogTagProcessor struct {
	domain     string
	logFactory glog.ILogFactory
}

func (this *LogTagProcessor) PrepareProcess() {

}

var _ gioc.ITagProcessor = (*LogTagProcessor)(nil)

func NewLogTagProcessor(domain string, logFactory glog.ILogFactory) *LogTagProcessor {
	return &LogTagProcessor{
		domain:     domain,
		logFactory: logFactory,
	}
}

func (this *LogTagProcessor) TagProcessorName() string {
	return "log"
}

func (this *LogTagProcessor) TagProcess(bean interface{}, fType reflect.StructField, fValue reflect.Value, tags map[string]string) {
	logName, logOk := tags[LogTag]
	if !logOk {
		return
	}

	if tags[LogDomain] != this.domain {
		return
	}
	logger := this.logFactory.GetLogger(logName)
	if logger == nil {
		util.Panic("logger %v:%v not exist", this.domain, logName)
	}
	fValue.Set(reflect.ValueOf(logger))
}
