package logic

import (
	"github.com/gosrv/gcluster/gbase/gdb/gredis"
	"github.com/gosrv/goioc"
)

/**
redis 使用示例
*/
type serviceRedisOptDemo struct {
	gioc.IBeanCondition
	// 动注入
	valueOperation      *gredis.ValueOperation      `redis.obj:""`
	boundValueOperation *gredis.BoundValueOperation `redis.obj:"test.boundval"`

	hashOperation      *gredis.HashOperation      `redis.obj:""`
	boundHashOperation *gredis.BoundHashOperation `redis.obj:"test.boundhash"`

	listOperation      *gredis.ListOperation      `redis.obj:""`
	boundListOperation *gredis.BoundListOperation `redis.obj:"test.boundlist"`

	setOperation      *gredis.SetOperation      `redis.obj:""`
	boundSetOperation *gredis.BoundSetOperation `redis.obj:"test.boundset"`

	zsetOperation      *gredis.ZSetOperation      `redis.obj:""`
	boundZSetOperation *gredis.BoundZSetOperation `redis.obj:"test.boundzset"`

	redisDriver gredis.IRedisDriver `bean`
}

func NewServiceRedisOptDemo() *serviceRedisOptDemo {
	return &serviceRedisOptDemo{
		IBeanCondition: gioc.NewConditionOnBean(gredis.IRedisDriverType, true),
	}
}

func (this *serviceRedisOptDemo) BeanStart() {
	_, _ = this.valueOperation.Set("test.value", "hello test.value")
	_, _ = this.valueOperation.Get("test.value")
	_, _ = this.boundValueOperation.Set("hello test.boundvalue")
	_, _ = this.boundValueOperation.Get()
	_, _ = this.hashOperation.Put("test.hash", "test.hashkey", "hello test.hash")
	_, _ = this.hashOperation.Get("test.hash", "test.hashkey")
	_, _ = this.boundHashOperation.Put("test.hashkey", "hello test.hash")
	_, _ = this.boundHashOperation.Get("test.hashkey")

	_, _ = this.listOperation.LeftPush("test.list", "a", "b", "c")
	_, _ = this.listOperation.RightPop("test.list")
	_, _ = this.boundListOperation.LeftPush("a", "b", "c")
	_, _ = this.boundListOperation.RightPop()

	_, _ = this.setOperation.Add("test.set", "a", "b", "c")
	_, _ = this.setOperation.Pop("test.set")
	_, _ = this.boundSetOperation.Add("a", "b", "c")
	_, _ = this.boundSetOperation.Pop()

	ag := this.redisDriver.GetAttributeGroup("player", "server")
	_ = ag.CasSetAttribute("gw", "", "123")
	_ = ag.CasSetAttribute("gw", "", "123")
	_ = ag.CasSetAttribute("gw", "123", "456")
	_ = ag.CasSetAttribute("gw", "456", "")
}

func (this *serviceRedisOptDemo) BeanStop() {

}
