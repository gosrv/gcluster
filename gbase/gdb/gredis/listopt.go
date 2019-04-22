package gredis

import (
	"github.com/go-redis/redis"
	"time"
)

type ListOperation struct {
	wrapKeyFunc WrapKeyFunc
	cmdable     redis.Cmdable
}

func NewListOperation(wrapKeyFunc WrapKeyFunc, cmdable redis.Cmdable) *ListOperation {
	return &ListOperation{wrapKeyFunc: wrapKeyFunc, cmdable: cmdable}
}
func (this *ListOperation) Expire(key string, duration time.Duration) (bool, error) {
	return this.cmdable.Expire(this.wrapKeyFunc(key), duration).Result()
}
func (this *ListOperation) LRange(key string, start int64, end int64) ([]string, error) {
	return this.cmdable.LRange(this.wrapKeyFunc(key), start, end).Result()
}
func (this *ListOperation) LTrim(key string, start int64, end int64) (string, error) {
	return this.cmdable.LTrim(this.wrapKeyFunc(key), start, end).Result()
}

func (this *ListOperation) LLen(key string) (int64, error) {
	return this.cmdable.LLen(this.wrapKeyFunc(key)).Result()
}
func (this *ListOperation) LPush(key string, value ...string) (int64, error) {
	return this.cmdable.LPush(this.wrapKeyFunc(key), value).Result()
}
func (this *ListOperation) LPushX(key string, value string) (int64, error) {
	return this.cmdable.LPushX(this.wrapKeyFunc(key), value).Result()
}
func (this *ListOperation) LInsert(key, op string, pivot, value string) (int64, error) {
	return this.cmdable.LInsert(this.wrapKeyFunc(key), op, pivot, value).Result()
}
func (this *ListOperation) RPush(key string, value ...interface{}) (int64, error) {
	return this.cmdable.RPush(this.wrapKeyFunc(key), value...).Result()
}
func (this *ListOperation) LSet(key string, index int64, value string) (string, error) {
	return this.cmdable.LSet(this.wrapKeyFunc(key), index, value).Result()
}
func (this *ListOperation) LRem(key string, count int64, value string) (int64, error) {
	return this.cmdable.LRem(this.wrapKeyFunc(key), count, value).Result()
}
func (this *ListOperation) LIndex(key string, index int64) (string, error) {
	return this.cmdable.LIndex(this.wrapKeyFunc(key), index).Result()
}
func (this *ListOperation) LPop(key string) (string, error) {
	return this.cmdable.LPop(this.wrapKeyFunc(key)).Result()
}
func (this *ListOperation) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(wkeys); i++ {
		wkeys[i] = this.wrapKeyFunc(keys[i])
	}
	return this.cmdable.BLPop(timeout, wkeys...).Result()
}
func (this *ListOperation) RPop(key string) (string, error) {
	return this.cmdable.RPop(this.wrapKeyFunc(key)).Result()
}
func (this *ListOperation) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(wkeys); i++ {
		wkeys[i] = this.wrapKeyFunc(keys[i])
	}
	return this.cmdable.BRPop(timeout, wkeys...).Result()
}
func (this *ListOperation) RPopLPush(sourceKey string, destinationKey string) (string, error) {
	return this.cmdable.RPopLPush(this.wrapKeyFunc(sourceKey), this.wrapKeyFunc(destinationKey)).Result()
}
func (this *ListOperation) BRPopLPush(source, destination string, timeout time.Duration) (string, error) {
	return this.cmdable.BRPopLPush(this.wrapKeyFunc(source), this.wrapKeyFunc(destination), timeout).Result()
}

type BoundListOperation struct {
	boundKey string
	cmdable  redis.Cmdable
}

func NewBoundListOperation(boundKey string, cmdable redis.Cmdable) *BoundListOperation {
	return &BoundListOperation{boundKey: boundKey, cmdable: cmdable}
}
func (this *BoundListOperation) Expire(key string, duration time.Duration) (bool, error) {
	return this.cmdable.Expire(this.boundKey, duration).Result()
}
func (this *BoundListOperation) LRange(start int64, end int64) ([]string, error) {
	return this.cmdable.LRange(this.boundKey, start, end).Result()
}
func (this *BoundListOperation) LTrim(start int64, end int64) (string, error) {
	return this.cmdable.LTrim(this.boundKey, start, end).Result()
}

func (this *BoundListOperation) LLen() (int64, error) {
	return this.cmdable.LLen(this.boundKey).Result()
}
func (this *BoundListOperation) LPush(value ...string) (int64, error) {
	return this.cmdable.LPush(this.boundKey, value).Result()
}
func (this *BoundListOperation) LPushX(value string) (int64, error) {
	return this.cmdable.LPushX(this.boundKey, value).Result()
}
func (this *BoundListOperation) LInsert(op string, pivot, value string) (int64, error) {
	return this.cmdable.LInsert(this.boundKey, op, pivot, value).Result()
}
func (this *BoundListOperation) RPush(value ...interface{}) (int64, error) {
	return this.cmdable.RPush(this.boundKey, value...).Result()
}
func (this *BoundListOperation) LSet(index int64, value string) (string, error) {
	return this.cmdable.LSet(this.boundKey, index, value).Result()
}
func (this *BoundListOperation) LRem(count int64, value string) (int64, error) {
	return this.cmdable.LRem(this.boundKey, count, value).Result()
}
func (this *BoundListOperation) LIndex(index int64) (string, error) {
	return this.cmdable.LIndex(this.boundKey, index).Result()
}
func (this *BoundListOperation) LPop() (string, error) {
	return this.cmdable.LPop(this.boundKey).Result()
}
func (this *BoundListOperation) RPop() (string, error) {
	return this.cmdable.RPop(this.boundKey).Result()
}
