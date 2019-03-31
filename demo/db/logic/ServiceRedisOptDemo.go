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
	// redis对象注入，redis对象名自动添加一个前缀，配置文件中的"app.group"，
	// 这是为了使多个不同的服务集群共享一个redis实例
	// 注入使用过标签redis.obj实现的，绑定对象还可以进一步指定特定的key
	// key - value对象的注入
	valueOperation *gredis.ValueOperation `redis.obj:""`
	// 绑定对象和非绑定对象的区别是，绑定对象的函数里不用再含有key，他就是针对特定key的操作
	boundValueOperation *gredis.BoundValueOperation `redis.obj:"test.boundval"`
	// hashmap对象的注入
	hashOperation      *gredis.HashOperation      `redis.obj:""`
	boundHashOperation *gredis.BoundHashOperation `redis.obj:"test.boundhash"`
	// list对象的注入
	listOperation      *gredis.ListOperation      `redis.obj:""`
	boundListOperation *gredis.BoundListOperation `redis.obj:"test.boundlist"`
	// set对象的注入
	setOperation      *gredis.SetOperation      `redis.obj:""`
	boundSetOperation *gredis.BoundSetOperation `redis.obj:"test.boundset"`
	// zset对象的注入
	zsetOperation      *gredis.ZSetOperation      `redis.obj:""`
	boundZSetOperation *gredis.BoundZSetOperation `redis.obj:"test.boundzset"`
	// 原始redis驱动的注入
	redisDriver gredis.IRedisDriver `bean`
}

func NewServiceRedisOptDemo() *serviceRedisOptDemo {
	return &serviceRedisOptDemo{
		// 依赖配置，只有配置了gleveldb.IRedisDriverType之后，这个bean才会生效
		IBeanCondition: gioc.NewConditionOnBeanType(gredis.IRedisDriverType, true),
	}
}

// 服务器启动前先调用BeanInit，然后调用BeanStart
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
