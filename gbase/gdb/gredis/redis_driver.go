package gredis

import (
	"github.com/go-redis/redis"
	"github.com/gosrv/gcluster/gbase/gdb"
	"github.com/gosrv/goioc"
	"reflect"
)

type WrapKeyFunc func(key string) string

type IRedisDriver interface {
	Domain() string
	RedisDB()
	GetHashOperation() *HashOperation
	GetBoundHashOperation(key string) *BoundHashOperation
	GetListOperation() *ListOperation
	GetBoundListOperation(key string) *BoundListOperation
	GetValueOperation() *ValueOperation
	GetBoundValueOperation(key string) *BoundValueOperation
	GetSetOperation() *SetOperation
	GetBoundSetOperation(key string) *BoundSetOperation
	GetZSetOperation() *ZSetOperation
	GetBoundZSetOperation(key string) *BoundZSetOperation
	GetPriority() int
	GetMessageQueue(group, id string) gdb.IMessageQueue
	GetAttributeGroup(group, id string) gdb.IDBAttributeGroup
}

var IRedisDriverType = reflect.TypeOf((*IRedisDriver)(nil)).Elem()

type redisDriver struct {
	domain   string
	appGroup string
	keySep   string
	redis.Cmdable
}

func (this *redisDriver) Domain() string {
	return this.domain
}

func NewRedisDriverStandalone(domain string, appGroup string, keySep string, opt *redis.Options) *redisDriver {
	return &redisDriver{
		domain:   domain,
		appGroup: appGroup,
		keySep:   keySep,
		Cmdable:  redis.NewClient(opt),
	}
}

func NewRedisDriverCluster(domain string, appGroup string, keySep string, opt *redis.ClusterOptions) *redisDriver {
	client := redis.NewClusterClient(opt)
	return &redisDriver{
		domain:   domain,
		appGroup: appGroup,
		keySep:   keySep,
		Cmdable:  client,
	}
}

func (this *redisDriver) RedisDB() {
}

func (this *redisDriver) GetWrapKeyFunc() WrapKeyFunc {
	return func(key string) string {
		return this.appGroup + this.keySep + key
	}
}

func (this *redisDriver) GetHashOperation() *HashOperation {
	return NewHashOperation(this.GetWrapKeyFunc(), this.Cmdable)
}

func (this *redisDriver) GetBoundHashOperation(key string) *BoundHashOperation {
	return NewBoundHashOperation(this.GetWrapKeyFunc()(key), this.Cmdable)
}

func (this *redisDriver) GetListOperation() *ListOperation {
	return NewListOperation(this.GetWrapKeyFunc(), this.Cmdable)
}

func (this *redisDriver) GetBoundListOperation(key string) *BoundListOperation {
	return NewBoundListOperation(this.GetWrapKeyFunc()(key), this.Cmdable)
}

func (this *redisDriver) GetValueOperation() *ValueOperation {
	return NewValueOperation(this.GetWrapKeyFunc(), this.Cmdable)
}

func (this *redisDriver) GetBoundValueOperation(key string) *BoundValueOperation {
	return NewBoundValueOperation(this.GetWrapKeyFunc()(key), this.Cmdable)
}

func (this *redisDriver) GetSetOperation() *SetOperation {
	return NewSetOperation(this.GetWrapKeyFunc(), this.Cmdable)
}

func (this *redisDriver) GetBoundSetOperation(key string) *BoundSetOperation {
	return NewBoundSetOperation(this.GetWrapKeyFunc()(key), this.Cmdable)
}

func (this *redisDriver) GetZSetOperation() *ZSetOperation {
	return NewZSetOperation(this.GetWrapKeyFunc(), this.Cmdable)
}

func (this *redisDriver) GetBoundZSetOperation(key string) *BoundZSetOperation {
	return NewBoundZSetOperation(this.GetWrapKeyFunc()(key), this.Cmdable)
}

func (this *redisDriver) GetPriority() int {
	return gioc.PriorityHigh
}

func (this *redisDriver) GetMessageQueue(group, id string) gdb.IMessageQueue {
	return NewRedisMessageQueue(group+this.keySep+id, this.Cmdable)
}

func (this *redisDriver) GetAttributeGroup(group, id string) gdb.IDBAttributeGroup {
	return NewRedisAttributeGroup(this.GetBoundHashOperation(group + this.keySep + id))
}
