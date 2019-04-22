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

func (this *HashOperation) HDel(key string, hashKeys ...string) (int64, error) {
	return this.cmdable.HDel(this.wrapKeyFunc(key), hashKeys...).Result()
}

func (this *HashOperation) HExists(key string, hashKey string) (bool, error) {
	return this.cmdable.HExists(this.wrapKeyFunc(key), hashKey).Result()
}

func (this *HashOperation) HGet(key string, hashKey string) (string, error) {
	return this.cmdable.HGet(this.wrapKeyFunc(key), hashKey).Result()
}

func (this *HashOperation) HMGet(key string, hashKeys ...string) ([]interface{}, error) {
	return this.cmdable.HMGet(this.wrapKeyFunc(key), hashKeys...).Result()
}

func (this *HashOperation) HIncrBy(key string, hashKey string, delta int64) (int64, error) {
	return this.cmdable.HIncrBy(this.wrapKeyFunc(key), hashKey, delta).Result()
}

func (this *HashOperation) HIncrByFloat(key string, hashKey string, delta float64) (float64, error) {
	return this.cmdable.HIncrByFloat(this.wrapKeyFunc(key), hashKey, delta).Result()
}

func (this *HashOperation) HKeys(key string) ([]string, error) {
	return this.cmdable.HKeys(this.wrapKeyFunc(key)).Result()
}

func (this *HashOperation) HLen(key string) (int64, error) {
	return this.cmdable.HLen(this.wrapKeyFunc(key)).Result()
}

func (this *HashOperation) HMSet(key string, kv map[string]interface{}) (string, error) {
	return this.cmdable.HMSet(this.wrapKeyFunc(key), kv).Result()
}

func (this *HashOperation) HSet(key string, hashKey string, value string) (bool, error) {
	return this.cmdable.HSet(this.wrapKeyFunc(key), hashKey, value).Result()
}

func (this *HashOperation) HSetNX(key string, hashKey string, value string) (bool, error) {
	return this.cmdable.HSetNX(this.wrapKeyFunc(key), hashKey, value).Result()
}

func (this *HashOperation) HGetAll(key string) (map[string]string, error) {
	return this.cmdable.HGetAll(this.wrapKeyFunc(key)).Result()
}

func (this *HashOperation) HScan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
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

func (this *BoundHashOperation) HDel(hashKeys ...string) (int64, error) {
	return this.cmdable.HDel(this.boundKey, hashKeys...).Result()
}

func (this *BoundHashOperation) HExists(hashKey string) (bool, error) {
	return this.cmdable.HExists(this.boundKey, hashKey).Result()
}

func (this *BoundHashOperation) HGet(hashKey string) (string, error) {
	return this.cmdable.HGet(this.boundKey, hashKey).Result()
}

func (this *BoundHashOperation) Expire(duration time.Duration) (bool, error) {
	return this.cmdable.Expire(this.boundKey, duration).Result()
}

func (this *BoundHashOperation) HMGet(hashKeys ...string) ([]interface{}, error) {
	return this.cmdable.HMGet(this.boundKey, hashKeys...).Result()
}

func (this *BoundHashOperation) HIncrBy(hashKey string, delta int64) (int64, error) {
	return this.cmdable.HIncrBy(this.boundKey, hashKey, delta).Result()
}

func (this *BoundHashOperation) HIncrByFloat(hashKey string, delta float64) (float64, error) {
	return this.cmdable.HIncrByFloat(this.boundKey, hashKey, delta).Result()
}

func (this *BoundHashOperation) HKeys() ([]string, error) {
	return this.cmdable.HKeys(this.boundKey).Result()
}

func (this *BoundHashOperation) HLen() (int64, error) {
	return this.cmdable.HLen(this.boundKey).Result()
}

func (this *BoundHashOperation) HMSet(kv map[string]interface{}) (string, error) {
	return this.cmdable.HMSet(this.boundKey, kv).Result()
}

func (this *BoundHashOperation) HSet(hashKey string, value string) (bool, error) {
	return this.cmdable.HSet(this.boundKey, hashKey, value).Result()
}

func (this *BoundHashOperation) HSetNX(hashKey string, value string) (bool, error) {
	return this.cmdable.HSetNX(this.boundKey, hashKey, value).Result()
}

func (this *BoundHashOperation) HGetAll() (map[string]string, error) {
	return this.cmdable.HGetAll(this.boundKey).Result()
}

func (this *BoundHashOperation) HScan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return this.cmdable.HScan(this.boundKey, cursor, match, count)
}
