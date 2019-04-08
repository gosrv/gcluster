package gmongo

import (
	"github.com/gosrv/gcluster/gbase/gl"
	"github.com/gosrv/goioc"
	"reflect"
)

const (
	MongoTag    = "mongo.c"
	MongoDomain = "mongo.domain"
)

type MongoTagProcessor struct {
	driver *mongoDBDriver
}

func (this *MongoTagProcessor) PrepareProcess() {

}

var _ gioc.ITagProcessor = (*MongoTagProcessor)(nil)

func NewMongoTagProcessor(driver *mongoDBDriver) *MongoTagProcessor {
	return &MongoTagProcessor{driver: driver}
}

func (this *MongoTagProcessor) TagProcessorName() string {
	return "mongo"
}

func (this *MongoTagProcessor) TagProcess(bean interface{}, fType reflect.StructField, fValue reflect.Value, tags map[string]string) {
	mongoCol, mongoColOk := tags[MongoTag]
	if !mongoColOk {
		return
	}
	if tags[MongoDomain] != this.driver.Name() {
		return
	}
	if len(mongoCol) == 0 {
		gl.Panic("mongo tag [%v] must has a value in [%v:%v]", MongoTag,
			reflect.TypeOf(bean), fValue.Type())
	}

	fValue.Set(reflect.ValueOf(this.driver.GetCollection(mongoCol)))
}
