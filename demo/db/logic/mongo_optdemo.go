package logic

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gosrv/gbase/gdb/gmongo"
	"github.com/gosrv/goioc"
)

type serviceMongoOptDemo struct {
	gioc.IBeanCondition
	coltest *mgo.Collection `mongo.c:"test"`
}

func NewServiceMongoOptDemo() *serviceMongoOptDemo {
	return &serviceMongoOptDemo{
		// 依赖配置，只有配置了gleveldb.IMongoDBDriverType之后，这个bean才会生效
		IBeanCondition: gioc.NewConditionOnBeanType(gmongo.IMongoDBDriverType, true),
	}
}

func (this *serviceMongoOptDemo) BeanStart() {
	_, _ = this.coltest.Upsert(bson.M{"_id": "123"}, bson.M{"$set": bson.M{"abc1": 1223}})
	res := make(map[string]interface{})
	_ = this.coltest.Find(bson.M{"_id": "123"}).One(&res)
}

func (this *serviceMongoOptDemo) BeanStop() {

}
