package gmongo

import (
	"github.com/gosrv/gcluster/gbase/glog"
	"reflect"
)

const (
	MongoTag    = "mongo.c"
	MongoDomain = "mongo.domain"
)

type MongoTagProcessor struct {
	driver *mongoDBDriver
}

func NewMongoTagProcessor(driver *mongoDBDriver) *MongoTagProcessor {
	return &MongoTagProcessor{driver: driver}
}

func (this *MongoTagProcessor) TagProcessorName() string {
	return "mongo"
}

func (this *MongoTagProcessor) TagProcess(bean interface{}, field reflect.Value, tags map[string]string) {
	mongoCol, mongoColOk := tags[MongoTag]
	if !mongoColOk {
		return
	}
	if tags[MongoDomain] != this.driver.Name() {
		return
	}
	if len(mongoCol) == 0 {
		glog.Panic("mongo tag [%v] must has a value in [%v:%v]", MongoTag,
			reflect.TypeOf(bean), field.Type())
	}

	field.Set(reflect.ValueOf(this.driver.GetCollection(mongoCol)))
}
