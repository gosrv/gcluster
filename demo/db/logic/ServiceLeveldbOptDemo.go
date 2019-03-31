package logic

import (
	"github.com/gosrv/gcluster/gbase/gdb/gleveldb"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
)

type serviceLevelDBOptDemo struct {
	gioc.IBeanCondition
	leveldb gleveldb.ILevelDBDriver `bean`
}

func NewServiceLevelDBOptDemo() *serviceLevelDBOptDemo {
	return &serviceLevelDBOptDemo{
		IBeanCondition: gioc.NewConditionOnBean(gleveldb.ILevelDBDriverType, true),
	}
}

func (this *serviceLevelDBOptDemo) BeanStart() {
	this.leveldb.Set("acc", "12345")
	data, _ := this.leveldb.Get("acc")
	util.Assert(data == "12345", "assert failed")
}

func (this *serviceLevelDBOptDemo) BeanStop() {

}
