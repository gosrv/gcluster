package gredis

import (
	"github.com/go-redis/redis"
	"time"
)

type ZSetOperation struct {
	wrapKeyFunc WrapKeyFunc
	cmdable     redis.Cmdable
}

func NewZSetOperation(wrapKeyFunc WrapKeyFunc, cmdable redis.Cmdable) *ZSetOperation {
	return &ZSetOperation{wrapKeyFunc: wrapKeyFunc, cmdable: cmdable}
}

func (this *ZSetOperation) BZPopMax(timeout time.Duration, keys ...string) (redis.ZWithKey, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.BZPopMax(timeout, wkeys...).Result()
}
func (this *ZSetOperation) BZPopMin(timeout time.Duration, keys ...string) (redis.ZWithKey, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.BZPopMin(timeout, wkeys...).Result()
}
func (this *ZSetOperation) Add(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAdd(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) ZAddNX(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddNX(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) ZAddXX(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddXX(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) ZAddCh(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddCh(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) ZAddNXCh(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddNXCh(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) ZAddXXCh(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddXXCh(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) ZIncr(key string, member redis.Z) (float64, error) {
	return this.cmdable.ZIncr(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) ZIncrNX(key string, member redis.Z) (float64, error) {
	return this.cmdable.ZIncrNX(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) ZIncrXX(key string, member redis.Z) (float64, error) {
	return this.cmdable.ZIncrXX(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) ZCard(key string) (int64, error) {
	return this.cmdable.ZCard(this.wrapKeyFunc(key)).Result()
}
func (this *ZSetOperation) ZCount(key, min, max string) (int64, error) {
	return this.cmdable.ZCount(this.wrapKeyFunc(key), min, max).Result()
}
func (this *ZSetOperation) ZLexCount(key, min, max string) (int64, error) {
	return this.cmdable.ZLexCount(this.wrapKeyFunc(key), min, max).Result()
}
func (this *ZSetOperation) ZIncrBy(key string, increment float64, member string) (float64, error) {
	return this.cmdable.ZIncrBy(this.wrapKeyFunc(key), increment, member).Result()
}
func (this *ZSetOperation) ZInterStore(destination string, store redis.ZStore, keys ...string) (int64, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.ZInterStore(this.wrapKeyFunc(destination), store, wkeys...).Result()
}
func (this *ZSetOperation) ZPopMax(key string, count ...int64) ([]redis.Z, error) {
	return this.cmdable.ZPopMax(this.wrapKeyFunc(key), count...).Result()
}
func (this *ZSetOperation) ZPopMin(key string, count ...int64) ([]redis.Z, error) {
	return this.cmdable.ZPopMin(this.wrapKeyFunc(key), count...).Result()
}
func (this *ZSetOperation) ZRange(key string, start, stop int64) ([]string, error) {
	return this.cmdable.ZRange(this.wrapKeyFunc(key), start, stop).Result()
}
func (this *ZSetOperation) ZRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return this.cmdable.ZRangeWithScores(this.wrapKeyFunc(key), start, stop).Result()
}
func (this *ZSetOperation) ZRangeByScore(key string, opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRangeByScore(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) ZRangeByLex(key string, opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRangeByLex(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) ([]redis.Z, error) {
	return this.cmdable.ZRangeByScoreWithScores(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) ZRank(key, member string) (int64, error) {
	return this.cmdable.ZRank(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) ZRem(key string, members ...interface{}) (int64, error) {
	return this.cmdable.ZRem(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	return this.cmdable.ZRemRangeByRank(this.wrapKeyFunc(key), start, stop).Result()
}
func (this *ZSetOperation) ZRemRangeByScore(key, min, max string) (int64, error) {
	return this.cmdable.ZRemRangeByScore(this.wrapKeyFunc(key), min, max).Result()
}
func (this *ZSetOperation) ZRemRangeByLex(key, min, max string) (int64, error) {
	return this.cmdable.ZRemRangeByLex(this.wrapKeyFunc(key), min, max).Result()
}
func (this *ZSetOperation) ZRevRange(key string, start, stop int64) ([]string, error) {
	return this.cmdable.ZRevRange(this.wrapKeyFunc(key), start, stop).Result()
}
func (this *ZSetOperation) ZRevRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return this.cmdable.ZRevRangeWithScores(this.wrapKeyFunc(key), start, stop).Result()
}
func (this *ZSetOperation) ZRevRangeByScore(key string, opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRevRangeByScore(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) ZRevRangeByLex(key string, opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRevRangeByLex(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) ZRevRangeByScoreWithScores(key string, opt redis.ZRangeBy) ([]redis.Z, error) {
	return this.cmdable.ZRevRangeByScoreWithScores(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) ZRevRank(key, member string) (int64, error) {
	return this.cmdable.ZRevRank(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) ZScore(key, member string) (float64, error) {
	return this.cmdable.ZScore(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) ZUnionStore(dest string, store redis.ZStore, keys ...string) (int64, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.ZUnionStore(this.wrapKeyFunc(dest), store, wkeys...).Result()
}

type BoundZSetOperation struct {
	boundKey string
	cmdable  redis.Cmdable
}

func NewBoundZSetOperation(boundKey string, cmdable redis.Cmdable) *BoundZSetOperation {
	return &BoundZSetOperation{boundKey: boundKey, cmdable: cmdable}
}

func (this *BoundZSetOperation) ZAdd(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAdd(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) ZAddNX(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddNX(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) ZAddXX(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddXX(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) ZAddCh(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddCh(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) ZAddNXCh(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddNXCh(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) ZAddXXCh(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddXXCh(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) ZIncr(member redis.Z) (float64, error) {
	return this.cmdable.ZIncr(this.boundKey, member).Result()
}
func (this *BoundZSetOperation) ZIncrNX(member redis.Z) (float64, error) {
	return this.cmdable.ZIncrNX(this.boundKey, member).Result()
}
func (this *BoundZSetOperation) ZIncrXX(member redis.Z) (float64, error) {
	return this.cmdable.ZIncrXX(this.boundKey, member).Result()
}
func (this *BoundZSetOperation) ZCard() (int64, error) {
	return this.cmdable.ZCard(this.boundKey).Result()
}
func (this *BoundZSetOperation) ZCount(min, max string) (int64, error) {
	return this.cmdable.ZCount(this.boundKey, min, max).Result()
}
func (this *BoundZSetOperation) ZLexCount(min, max string) (int64, error) {
	return this.cmdable.ZLexCount(this.boundKey, min, max).Result()
}
func (this *BoundZSetOperation) ZIncrBy(increment float64, member string) (float64, error) {
	return this.cmdable.ZIncrBy(this.boundKey, increment, member).Result()
}
func (this *BoundZSetOperation) ZPopMax(count ...int64) ([]redis.Z, error) {
	return this.cmdable.ZPopMax(this.boundKey, count...).Result()
}
func (this *BoundZSetOperation) ZPopMin(count ...int64) ([]redis.Z, error) {
	return this.cmdable.ZPopMin(this.boundKey, count...).Result()
}
func (this *BoundZSetOperation) ZRange(start, stop int64) ([]string, error) {
	return this.cmdable.ZRange(this.boundKey, start, stop).Result()
}
func (this *BoundZSetOperation) ZRangeWithScores(start, stop int64) ([]redis.Z, error) {
	return this.cmdable.ZRangeWithScores(this.boundKey, start, stop).Result()
}
func (this *BoundZSetOperation) ZRangeByScore(opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRangeByScore(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) ZRangeByLex(opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRangeByLex(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) ZRangeByScoreWithScores(opt redis.ZRangeBy) ([]redis.Z, error) {
	return this.cmdable.ZRangeByScoreWithScores(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) ZRank(member string) (int64, error) {
	return this.cmdable.ZRank(this.boundKey, member).Result()
}
func (this *BoundZSetOperation) ZRem(members ...interface{}) (int64, error) {
	return this.cmdable.ZRem(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) ZRemRangeByRank(start, stop int64) (int64, error) {
	return this.cmdable.ZRemRangeByRank(this.boundKey, start, stop).Result()
}
func (this *BoundZSetOperation) ZRemRangeByScore(min, max string) (int64, error) {
	return this.cmdable.ZRemRangeByScore(this.boundKey, min, max).Result()
}
func (this *BoundZSetOperation) ZRemRangeByLex(min, max string) (int64, error) {
	return this.cmdable.ZRemRangeByLex(this.boundKey, min, max).Result()
}
func (this *BoundZSetOperation) ZRevRange(start, stop int64) ([]string, error) {
	return this.cmdable.ZRevRange(this.boundKey, start, stop).Result()
}
func (this *BoundZSetOperation) ZRevRangeWithScores(start, stop int64) ([]redis.Z, error) {
	return this.cmdable.ZRevRangeWithScores(this.boundKey, start, stop).Result()
}
func (this *BoundZSetOperation) ZRevRangeByScore(opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRevRangeByScore(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) ZRevRangeByLex(opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRevRangeByLex(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) ZRevRangeByScoreWithScores(opt redis.ZRangeBy) ([]redis.Z, error) {
	return this.cmdable.ZRevRangeByScoreWithScores(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) ZRevRank(member string) (int64, error) {
	return this.cmdable.ZRevRank(this.boundKey, member).Result()
}
func (this *BoundZSetOperation) ZScore(member string) (float64, error) {
	return this.cmdable.ZScore(this.boundKey, member).Result()
}
