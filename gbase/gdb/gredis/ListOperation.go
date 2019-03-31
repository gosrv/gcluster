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
func (this *ListOperation) Range(key string, start int64, end int64) ([]string, error) {
	return this.cmdable.LRange(this.wrapKeyFunc(key), start, end).Result()
}
func (this *ListOperation) Trim(key string, start int64, end int64) (string, error) {
	return this.cmdable.LTrim(this.wrapKeyFunc(key), start, end).Result()
}

func (this *ListOperation) Size(key string) (int64, error) {
	return this.cmdable.LLen(this.wrapKeyFunc(key)).Result()
}
func (this *ListOperation) LeftPush(key string, value ...string) (int64, error) {
	return this.cmdable.LPush(this.wrapKeyFunc(key), value).Result()
}
func (this *ListOperation) LeftPushIfPresent(key string, value string) (int64, error) {
	return this.cmdable.LPushX(this.wrapKeyFunc(key), value).Result()
}
func (this *ListOperation) LeftInsert(key, op string, pivot, value string) (int64, error) {
	return this.cmdable.LInsert(this.wrapKeyFunc(key), op, pivot, value).Result()
}
func (this *ListOperation) RightPush(key string, value ...interface{}) (int64, error) {
	return this.cmdable.RPush(this.wrapKeyFunc(key), value...).Result()
}
func (this *ListOperation) Set(key string, index int64, value string) (string, error) {
	return this.cmdable.LSet(this.wrapKeyFunc(key), index, value).Result()
}
func (this *ListOperation) Remove(key string, count int64, value string) (int64, error) {
	return this.cmdable.LRem(this.wrapKeyFunc(key), count, value).Result()
}
func (this *ListOperation) Index(key string, index int64) (string, error) {
	return this.cmdable.LIndex(this.wrapKeyFunc(key), index).Result()
}
func (this *ListOperation) LeftPop(key string) (string, error) {
	return this.cmdable.LPop(this.wrapKeyFunc(key)).Result()
}
func (this *ListOperation) BLeftPop(timeout time.Duration, keys ...string) ([]string, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(wkeys); i++ {
		wkeys[i] = this.wrapKeyFunc(keys[i])
	}
	return this.cmdable.BLPop(timeout, wkeys...).Result()
}
func (this *ListOperation) RightPop(key string) (string, error) {
	return this.cmdable.RPop(this.wrapKeyFunc(key)).Result()
}
func (this *ListOperation) BRightPop(timeout time.Duration, keys ...string) ([]string, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(wkeys); i++ {
		wkeys[i] = this.wrapKeyFunc(keys[i])
	}
	return this.cmdable.BRPop(timeout, wkeys...).Result()
}
func (this *ListOperation) RightPopAndLeftPush(sourceKey string, destinationKey string) (string, error) {
	return this.cmdable.RPopLPush(this.wrapKeyFunc(sourceKey), this.wrapKeyFunc(destinationKey)).Result()
}
func (this *ListOperation) BRightPopAndLeftPush(source, destination string, timeout time.Duration) (string, error) {
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
func (this *BoundListOperation) Range(start int64, end int64) ([]string, error) {
	return this.cmdable.LRange(this.boundKey, start, end).Result()
}
func (this *BoundListOperation) Trim(start int64, end int64) (string, error) {
	return this.cmdable.LTrim(this.boundKey, start, end).Result()
}

func (this *BoundListOperation) Size() (int64, error) {
	return this.cmdable.LLen(this.boundKey).Result()
}
func (this *BoundListOperation) LeftPush(value ...string) (int64, error) {
	return this.cmdable.LPush(this.boundKey, value).Result()
}
func (this *BoundListOperation) LeftPushIfPresent(value string) (int64, error) {
	return this.cmdable.LPushX(this.boundKey, value).Result()
}
func (this *BoundListOperation) LeftInsert(op string, pivot, value string) (int64, error) {
	return this.cmdable.LInsert(this.boundKey, op, pivot, value).Result()
}
func (this *BoundListOperation) RightPush(value ...interface{}) (int64, error) {
	return this.cmdable.RPush(this.boundKey, value...).Result()
}
func (this *BoundListOperation) Set(index int64, value string) (string, error) {
	return this.cmdable.LSet(this.boundKey, index, value).Result()
}
func (this *BoundListOperation) Remove(count int64, value string) (int64, error) {
	return this.cmdable.LRem(this.boundKey, count, value).Result()
}
func (this *BoundListOperation) Index(index int64) (string, error) {
	return this.cmdable.LIndex(this.boundKey, index).Result()
}
func (this *BoundListOperation) LeftPop() (string, error) {
	return this.cmdable.LPop(this.boundKey).Result()
}
func (this *BoundListOperation) RightPop() (string, error) {
	return this.cmdable.RPop(this.boundKey).Result()
}
