package gleveldb

import (
	"github.com/gosrv/goioc"
)

type AutoConfigLevelDB struct {
	appGroup string `cfg:"app.group"`
	// 启动条件
	gioc.IBeanCondition
	gioc.IConfigBase
	gioc.IBeanName
	url string `cfg.d:"leveldb.url"`
	ILevelDBDriver
}

func NewAutoConfigLevelDB(cfgBase, name string) *AutoConfigLevelDB {
	return &AutoConfigLevelDB{
		IBeanCondition: gioc.NewConditionOnValue(cfgBase, true),
		IConfigBase:    gioc.NewConfigBase(cfgBase),
		IBeanName:      gioc.NewBeanName(name),
	}
}

func (this *AutoConfigLevelDB) BeanAfterTagProcess(tagProcessor gioc.ITagProcessor, beanContainer gioc.IBeanContainer) {
	if tagProcessor.TagProcessorName() != gioc.ConfigTagProcessor {
		return
	}
	this.ILevelDBDriver = NewLevelDBDriver(this.url, this.appGroup)
}
