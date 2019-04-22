package gredis

import (
	"github.com/go-redis/redis"
	"time"
)

type ValueOperation struct {
	wrapKeyFunc WrapKeyFunc
	cmdable     redis.Cmdable
}

func NewValueOperation(wrapKeyFunc WrapKeyFunc, cmdable redis.Cmdable) *ValueOperation {
	return &ValueOperation{wrapKeyFunc: wrapKeyFunc, cmdable: cmdable}
}

func (this *ValueOperation) Set(key string, value string) (string, error) {
	return this.cmdable.Set(this.wrapKeyFunc(key), value, 0).Result()
}

func (this *ValueOperation) Expire(key string, duration time.Duration) (bool, error) {
	return this.cmdable.Expire(this.wrapKeyFunc(key), duration).Result()
}

func (this *ValueOperation) SetTimeout(key string, value string, timeout time.Duration) (string, error) {
	return this.cmdable.Set(this.wrapKeyFunc(key), value, timeout).Result()
}
func (this *ValueOperation) SetNX(key string, value string) (bool, error) {
	return this.cmdable.SetNX(this.wrapKeyFunc(key), value, 0).Result()
}
func (this *ValueOperation) SetNXTimeout(key string, value string, timeout time.Duration) (bool, error) {
	return this.cmdable.SetNX(this.wrapKeyFunc(key), value, timeout).Result()
}
func (this *ValueOperation) MSet(kvalues map[string]string) (string, error) {
	pairs := make([]interface{}, len(kvalues)*2, len(kvalues)*2)
	idx := 0
	for k, v := range kvalues {
		pairs[idx] = this.wrapKeyFunc(k)
		pairs[idx+1] = v
		idx += 2
	}
	return this.cmdable.MSet(pairs...).Result()
}
func (this *ValueOperation) MSetNX(kvalues map[string]string) (bool, error) {
	pairs := make([]interface{}, len(kvalues)*2, len(kvalues)*2)
	idx := 0
	for k, v := range kvalues {
		pairs[idx] = this.wrapKeyFunc(k)
		pairs[idx+1] = v
		idx += 2
	}
	return this.cmdable.MSetNX(pairs...).Result()
}
func (this *ValueOperation) Get(key string) (string, error) {
	return this.cmdable.Get(this.wrapKeyFunc(key)).Result()
}
func (this *ValueOperation) GetSet(key string, value string) (string, error) {
	return this.cmdable.GetSet(this.wrapKeyFunc(key), value).Result()
}
func (this *ValueOperation) MGet(keys ...string) ([]string, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = this.wrapKeyFunc(keys[i])
	}
	vals, err := this.cmdable.MGet(wkeys...).Result()
	if len(vals) <= 0 {
		return nil, err
	}
	svals := make([]string, len(vals), len(vals))
	for i := 0; i < len(vals); i++ {
		svals[i] = vals[i].(string)
	}
	return svals, err
}
func (this *ValueOperation) Incr(key string) (int64, error) {
	return this.cmdable.Incr(this.wrapKeyFunc(key)).Result()
}
func (this *ValueOperation) IncrBy(key string, delta int64) (int64, error) {
	return this.cmdable.IncrBy(this.wrapKeyFunc(key), delta).Result()
}
func (this *ValueOperation) IncrByFloat(key string, delta float64) (float64, error) {
	return this.cmdable.IncrByFloat(this.wrapKeyFunc(key), delta).Result()
}
func (this *ValueOperation) Decr(key string) (int64, error) {
	return this.cmdable.Decr(this.wrapKeyFunc(key)).Result()
}
func (this *ValueOperation) DecrBy(key string, delta int64) (int64, error) {
	return this.cmdable.DecrBy(this.wrapKeyFunc(key), delta).Result()
}
func (this *ValueOperation) Append(key string, value string) (int64, error) {
	return this.cmdable.Append(this.wrapKeyFunc(key), value).Result()
}
func (this *ValueOperation) GetRange(key string, start int64, end int64) (string, error) {
	return this.cmdable.GetRange(this.wrapKeyFunc(key), start, end).Result()
}
func (this *ValueOperation) SetRange(key string, value string, offset int64) (int64, error) {
	return this.cmdable.SetRange(this.wrapKeyFunc(key), offset, value).Result()
}
func (this *ValueOperation) StrLen(key string) (int64, error) {
	return this.cmdable.StrLen(this.wrapKeyFunc(key)).Result()
}
func (this *ValueOperation) SetBit(key string, offset int64, value int) (int64, error) {
	return this.cmdable.SetBit(this.wrapKeyFunc(key), offset, value).Result()
}
func (this *ValueOperation) GetBit(key string, offset int64) (int64, error) {
	return this.cmdable.GetBit(this.wrapKeyFunc(key), offset).Result()
}

type BoundValueOperation struct {
	boundKey string
	cmdable  redis.Cmdable
}

func NewBoundValueOperation(boundKey string, cmdable redis.Cmdable) *BoundValueOperation {
	return &BoundValueOperation{boundKey: boundKey, cmdable: cmdable}
}

func (this *BoundValueOperation) Set(value string) (string, error) {
	return this.cmdable.Set(this.boundKey, value, 0).Result()
}
func (this *BoundValueOperation) Expire(key string, duration time.Duration) (bool, error) {
	return this.cmdable.Expire(this.boundKey, duration).Result()
}
func (this *BoundValueOperation) SetTimeout(value string, timeout time.Duration) (string, error) {
	return this.cmdable.Set(this.boundKey, value, timeout).Result()
}
func (this *BoundValueOperation) SetNX(value string) (bool, error) {
	return this.cmdable.SetNX(this.boundKey, value, 0).Result()
}
func (this *BoundValueOperation) SetNXTimeout(value string, timeout time.Duration) (bool, error) {
	return this.cmdable.SetNX(this.boundKey, value, timeout).Result()
}

func (this *BoundValueOperation) Get() (string, error) {
	return this.cmdable.Get(this.boundKey).Result()
}
func (this *BoundValueOperation) GetSet(value string) (string, error) {
	return this.cmdable.GetSet(this.boundKey, value).Result()
}

func (this *BoundValueOperation) Incr() (int64, error) {
	return this.cmdable.Incr(this.boundKey).Result()
}
func (this *BoundValueOperation) IncrBy(delta int64) (int64, error) {
	return this.cmdable.IncrBy(this.boundKey, delta).Result()
}
func (this *BoundValueOperation) IncrByFloat(delta float64) (float64, error) {
	return this.cmdable.IncrByFloat(this.boundKey, delta).Result()
}
func (this *BoundValueOperation) Decr() (int64, error) {
	return this.cmdable.Decr(this.boundKey).Result()
}
func (this *BoundValueOperation) DecrBy(delta int64) (int64, error) {
	return this.cmdable.DecrBy(this.boundKey, delta).Result()
}
func (this *BoundValueOperation) Append(value string) (int64, error) {
	return this.cmdable.Append(this.boundKey, value).Result()
}
func (this *BoundValueOperation) GetRange(start int64, end int64) (string, error) {
	return this.cmdable.GetRange(this.boundKey, start, end).Result()
}
func (this *BoundValueOperation) SetRange(value string, offset int64) (int64, error) {
	return this.cmdable.SetRange(this.boundKey, offset, value).Result()
}
func (this *BoundValueOperation) StrLen() (int64, error) {
	return this.cmdable.StrLen(this.boundKey).Result()
}
func (this *BoundValueOperation) SetBit(offset int64, value int) (int64, error) {
	return this.cmdable.SetBit(this.boundKey, offset, value).Result()
}
func (this *BoundValueOperation) GetBit(offset int64) (int64, error) {
	return this.cmdable.GetBit(this.boundKey, offset).Result()
}
