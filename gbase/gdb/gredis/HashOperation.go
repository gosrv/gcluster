package gredis

import (
	"github.com/go-redis/redis"
	"time"
)

const (
	CAS_SCRIPT = "local l=1;\n local rawVal = redis.call('hget', KEYS[1], ARGV[1])\n" +
		"if not rawVal then rawVal='' end\n" +
		"if rawVal == ARGV[2] then\n" +
		"   redis.call('hset', KEYS[1], ARGV[1], ARGV[3])\n" +
		"   return true\n" +
		"else\n" +
		"   return false\n" +
		"end\n"
)

type HashOperation struct {
	wrapKeyFunc WrapKeyFunc
	cmdable     redis.Cmdable
}

func NewHashOperation(wrapKeyFunc WrapKeyFunc, cmdable redis.Cmdable) *HashOperation {
	return &HashOperation{wrapKeyFunc: wrapKeyFunc, cmdable: cmdable}
}

func (this *HashOperation) Cas(key string, hashKey string, old string, new string) (bool, error) {
	return this.cmdable.Eval(CAS_SCRIPT, []string{this.wrapKeyFunc(key)}, hashKey, old, new).Bool()
}

func (this *HashOperation) Expire(key string, duration time.Duration) (bool, error) {
	return this.cmdable.Expire(this.wrapKeyFunc(key), duration).Result()
}

func (this *HashOperation) Delete(key string, hashKeys ...string) (int64, error) {
	return this.cmdable.HDel(this.wrapKeyFunc(key), hashKeys...).Result()
}

func (this *HashOperation) HasKey(key string, hashKey string) (bool, error) {
	return this.cmdable.HExists(this.wrapKeyFunc(key), hashKey).Result()
}

func (this *HashOperation) Get(key string, hashKey string) (string, error) {
	return this.cmdable.HGet(this.wrapKeyFunc(key), hashKey).Result()
}

func (this *HashOperation) MultiGet(key string, hashKeys ...string) ([]interface{}, error) {
	return this.cmdable.HMGet(this.wrapKeyFunc(key), hashKeys...).Result()
}

func (this *HashOperation) Increment(key string, hashKey string, delta int64) (int64, error) {
	return this.cmdable.HIncrBy(this.wrapKeyFunc(key), hashKey, delta).Result()
}

func (this *HashOperation) IncrementByFloat(key string, hashKey string, delta float64) (float64, error) {
	return this.cmdable.HIncrByFloat(this.wrapKeyFunc(key), hashKey, delta).Result()
}

func (this *HashOperation) Keys(key string) ([]string, error) {
	return this.cmdable.HKeys(this.wrapKeyFunc(key)).Result()
}

func (this *HashOperation) Size(key string) (int64, error) {
	return this.cmdable.HLen(this.wrapKeyFunc(key)).Result()
}

func (this *HashOperation) PutAll(key string, kv map[string]interface{}) (string, error) {
	return this.cmdable.HMSet(this.wrapKeyFunc(key), kv).Result()
}

func (this *HashOperation) Put(key string, hashKey string, value string) (bool, error) {
	return this.cmdable.HSet(this.wrapKeyFunc(key), hashKey, value).Result()
}

func (this *HashOperation) PutIfAbsent(key string, hashKey string, value string) (bool, error) {
	return this.cmdable.HSetNX(this.wrapKeyFunc(key), hashKey, value).Result()
}

func (this *HashOperation) Entries(key string) (map[string]string, error) {
	return this.cmdable.HGetAll(this.wrapKeyFunc(key)).Result()
}

func (this *HashOperation) Scan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return this.cmdable.HScan(this.wrapKeyFunc(key), cursor, match, count)
}

type BoundHashOperation struct {
	boundKey string
	cmdable  redis.Cmdable
}

func NewBoundHashOperation(boundKey string, cmdable redis.Cmdable) *BoundHashOperation {
	return &BoundHashOperation{boundKey: boundKey, cmdable: cmdable}
}

func (this *BoundHashOperation) Cas(hashKey string, old string, new string) (bool, error) {
	return this.cmdable.Eval(CAS_SCRIPT, []string{this.boundKey}, hashKey, old, new).Bool()
}

func (this *BoundHashOperation) Delete(hashKeys ...string) (int64, error) {
	return this.cmdable.HDel(this.boundKey, hashKeys...).Result()
}

func (this *BoundHashOperation) HasKey(hashKey string) (bool, error) {
	return this.cmdable.HExists(this.boundKey, hashKey).Result()
}

func (this *BoundHashOperation) Get(hashKey string) (string, error) {
	return this.cmdable.HGet(this.boundKey, hashKey).Result()
}

func (this *BoundHashOperation) Expire(duration time.Duration) (bool, error) {
	return this.cmdable.Expire(this.boundKey, duration).Result()
}

func (this *BoundHashOperation) MultiGet(hashKeys ...string) ([]interface{}, error) {
	return this.cmdable.HMGet(this.boundKey, hashKeys...).Result()
}

func (this *BoundHashOperation) Increment(hashKey string, delta int64) (int64, error) {
	return this.cmdable.HIncrBy(this.boundKey, hashKey, delta).Result()
}

func (this *BoundHashOperation) IncrementByFloat(hashKey string, delta float64) (float64, error) {
	return this.cmdable.HIncrByFloat(this.boundKey, hashKey, delta).Result()
}

func (this *BoundHashOperation) Keys() ([]string, error) {
	return this.cmdable.HKeys(this.boundKey).Result()
}

func (this *BoundHashOperation) Size() (int64, error) {
	return this.cmdable.HLen(this.boundKey).Result()
}

func (this *BoundHashOperation) PutAll(kv map[string]interface{}) (string, error) {
	return this.cmdable.HMSet(this.boundKey, kv).Result()
}

func (this *BoundHashOperation) Put(hashKey string, value string) (bool, error) {
	return this.cmdable.HSet(this.boundKey, hashKey, value).Result()
}

func (this *BoundHashOperation) PutIfAbsent(hashKey string, value string) (bool, error) {
	return this.cmdable.HSetNX(this.boundKey, hashKey, value).Result()
}

func (this *BoundHashOperation) Entries() (map[string]string, error) {
	return this.cmdable.HGetAll(this.boundKey).Result()
}

func (this *BoundHashOperation) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return this.cmdable.HScan(this.boundKey, cursor, match, count)
}
