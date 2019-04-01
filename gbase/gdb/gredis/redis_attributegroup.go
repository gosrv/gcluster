package gredis

import "time"

type RedisAttributeGroup struct {
	opt *BoundHashOperation
}

func NewRedisAttributeGroup(opt *BoundHashOperation) *RedisAttributeGroup {
	return &RedisAttributeGroup{opt: opt}
}

func (this *RedisAttributeGroup) SetExpireDuration(duration time.Duration) error {
	ok, err := this.opt.Expire(duration)
	if ok {
		return nil
	}
	return err
}

func (this *RedisAttributeGroup) CasSetAttribute(key string, oldValue string, newValue string) bool {
	success, error := this.opt.Cas(key, oldValue, newValue)
	return success && error == nil
}

func (this *RedisAttributeGroup) GetAttribute(key string) (string, error) {
	return this.opt.Get(key)
}

func (this *RedisAttributeGroup) SetAttribute(key string, value string) error {
	_, err := this.opt.Put(key, value)
	return err
}

func (this *RedisAttributeGroup) SetAttributes(values map[string]interface{}) error {
	_, err := this.opt.PutAll(values)
	return err
}
