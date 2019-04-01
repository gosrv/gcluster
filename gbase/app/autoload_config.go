package app

import (
	"github.com/gosrv/gcluster/gbase/gutil"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
)

/**
监控配置文件变化，变化后重新注入所有的配置数据到所有的bean中
*/
type autoLoadConfig struct {
	configLoader  gioc.IConfigLoader   `bean`
	beanContainer gioc.IBeanContainer  `bean`
	tagParser     gioc.ITagParser      `bean`
	tagProcessors []gioc.ITagProcessor `bean`
	stop          bool
}

func NewAutoLoadConfig() *autoLoadConfig {
	return &autoLoadConfig{
		stop: false,
	}
}

func (this *autoLoadConfig) BeanStart() {
	var configTagProcessor gioc.ITagProcessor
	for _, tagProcessor := range this.tagProcessors {
		if tagProcessor.TagProcessorName() == gioc.ConfigTagProcessor {
			configTagProcessor = tagProcessor
			break
		}
	}
	util.VerifyNotNull(configTagProcessor)

	gutil.RecoverGo(func() {
		for !this.stop {
			this.configLoader.AutoLoad(func() {
				for _, bean := range this.beanContainer.GetAllBeans() {
					gioc.TagProcessorHelper.BeanTagProcess(bean, this.tagParser, configTagProcessor)
				}
			})
		}
	})
}

func (this *autoLoadConfig) BeanStop() {
	this.stop = true
}
