package gredis

import (
	"github.com/gosrv/gcluster/gbase/gl"
	"github.com/gosrv/goioc"
	"reflect"
)

const (
	RedisObjTag = "redis.obj"
	RedisDomain = "redis.domain"
)

type redisTagProcessor struct {
	domain               string
	boundRedisObjMaker   map[reflect.Type]func(id string) interface{}
	unboundRedisObjMaker map[reflect.Type]func() interface{}
}

func (this *redisTagProcessor) PrepareProcess() {

}

var _ gioc.ITagProcessor = (*redisTagProcessor)(nil)

func NewRedisTagProcessor(driver IRedisDriver) *redisTagProcessor {
	processor := &redisTagProcessor{domain: driver.Domain()}

	processor.boundRedisObjMaker = make(map[reflect.Type]func(id string) interface{})
	processor.boundRedisObjMaker[reflect.TypeOf((*BoundValueOperation)(nil))] = func(id string) interface{} {
		return driver.GetBoundValueOperation(id)
	}
	processor.boundRedisObjMaker[reflect.TypeOf((*BoundHashOperation)(nil))] = func(id string) interface{} {
		return driver.GetBoundHashOperation(id)
	}
	processor.boundRedisObjMaker[reflect.TypeOf((*BoundListOperation)(nil))] = func(id string) interface{} {
		return driver.GetBoundListOperation(id)
	}
	processor.boundRedisObjMaker[reflect.TypeOf((*BoundSetOperation)(nil))] = func(id string) interface{} {
		return driver.GetBoundSetOperation(id)
	}
	processor.boundRedisObjMaker[reflect.TypeOf((*BoundZSetOperation)(nil))] = func(id string) interface{} {
		return driver.GetBoundZSetOperation(id)
	}

	processor.unboundRedisObjMaker = make(map[reflect.Type]func() interface{})
	processor.unboundRedisObjMaker[reflect.TypeOf((*ValueOperation)(nil))] = func() interface{} {
		return driver.GetValueOperation()
	}
	processor.unboundRedisObjMaker[reflect.TypeOf((*HashOperation)(nil))] = func() interface{} {
		return driver.GetHashOperation()
	}
	processor.unboundRedisObjMaker[reflect.TypeOf((*ListOperation)(nil))] = func() interface{} {
		return driver.GetListOperation()
	}
	processor.unboundRedisObjMaker[reflect.TypeOf((*SetOperation)(nil))] = func() interface{} {
		return driver.GetSetOperation()
	}
	processor.unboundRedisObjMaker[reflect.TypeOf((*ZSetOperation)(nil))] = func() interface{} {
		return driver.GetZSetOperation()
	}

	return processor
}

func (this *redisTagProcessor) TagProcessorName() string {
	return "redis"
}

func (this *redisTagProcessor) TagProcess(bean interface{}, fType reflect.StructField, fValue reflect.Value, tags map[string]string) {
	redisObjName, redisObjNameOk := tags[RedisObjTag]
	if !redisObjNameOk {
		return
	}
	if tags[RedisDomain] != this.domain {
		return
	}

	var redisObj interface{}
	if len(redisObjName) > 0 {
		// bound redis object
		maker, ok := this.boundRedisObjMaker[fValue.Type()]
		if !ok {
			gl.Panic("can not find bound redis object type %v with name %v, in bean %v", fValue.Type(),
				redisObjName, reflect.TypeOf(bean))
			return
		}
		redisObj = maker(redisObjName)
	} else {
		// unbound redis object
		maker, ok := this.unboundRedisObjMaker[fValue.Type()]
		if !ok {
			gl.Panic("can not find unbound redis object type %v, in bean %v", fValue.Type(),
				redisObjName, reflect.TypeOf(bean))
			return
		}
		redisObj = maker()
	}

	if redisObj == nil {
		gl.Panic("make redis object type %v with name %v, in bean %v failed", fValue.Type(),
			redisObjName, reflect.TypeOf(bean))
		return
	}

	fValue.Set(reflect.ValueOf(redisObj))
}
