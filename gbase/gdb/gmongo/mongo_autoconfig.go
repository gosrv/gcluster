package gmongo

import (
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
	"reflect"
	"strings"
)

const (
	mongoConfigUrlValue = "mongo.url"
)

type AutoConfigMongo struct {
	appGroup string `cfg:"app.group"`
	// 启动条件
	gioc.IBeanCondition
	gioc.IConfigBase
	url string `cfg.d:"mongo.url"`
	*mongoDBDriver
	tagProcessor gioc.ITagProcessor
	domain       string
}

var _ gioc.ITagProcessor = (*AutoConfigMongo)(nil)

func (this *AutoConfigMongo) TagProcessorName() string {
	return this.tagProcessor.TagProcessorName()
}

func (this *AutoConfigMongo) TagProcess(bean interface{}, fType reflect.StructField, fValue reflect.Value, tags map[string]string) {
	this.tagProcessor.TagProcess(bean, fType, fValue, tags)
}

func NewAutoConfigMongo(cfgBase, domain string) *AutoConfigMongo {
	return &AutoConfigMongo{
		IBeanCondition: gioc.NewConditionOnValue(cfgBase, true),
		IConfigBase:    gioc.NewConfigBase(cfgBase),
		domain:         domain,
	}
}

func (this *AutoConfigMongo) PrepareProcess() {
	util.Assert(this.mongoDBDriver == nil, "")

	this.mongoDBDriver = NewMongoDriver(this.domain, this.url, strings.Replace(this.appGroup, ".", "_", -1))
	this.tagProcessor = NewMongoTagProcessor(this.mongoDBDriver)
}
