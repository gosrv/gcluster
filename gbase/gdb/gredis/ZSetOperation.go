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

func (this *ZSetOperation) BPopMax(timeout time.Duration, keys ...string) (redis.ZWithKey, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.BZPopMax(timeout, wkeys...).Result()
}
func (this *ZSetOperation) BPopMin(timeout time.Duration, keys ...string) (redis.ZWithKey, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.BZPopMin(timeout, wkeys...).Result()
}
func (this *ZSetOperation) Add(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAdd(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) AddNX(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddNX(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) AddXX(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddXX(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) AddCh(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddCh(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) AddNXCh(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddNXCh(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) AddXXCh(key string, members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddXXCh(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) Incr(key string, member redis.Z) (float64, error) {
	return this.cmdable.ZIncr(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) IncrNX(key string, member redis.Z) (float64, error) {
	return this.cmdable.ZIncrNX(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) IncrXX(key string, member redis.Z) (float64, error) {
	return this.cmdable.ZIncrXX(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) Card(key string) (int64, error) {
	return this.cmdable.ZCard(this.wrapKeyFunc(key)).Result()
}
func (this *ZSetOperation) Count(key, min, max string) (int64, error) {
	return this.cmdable.ZCount(this.wrapKeyFunc(key), min, max).Result()
}
func (this *ZSetOperation) LexCount(key, min, max string) (int64, error) {
	return this.cmdable.ZLexCount(this.wrapKeyFunc(key), min, max).Result()
}
func (this *ZSetOperation) IncrBy(key string, increment float64, member string) (float64, error) {
	return this.cmdable.ZIncrBy(this.wrapKeyFunc(key), increment, member).Result()
}
func (this *ZSetOperation) InterStore(destination string, store redis.ZStore, keys ...string) (int64, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.ZInterStore(this.wrapKeyFunc(destination), store, wkeys...).Result()
}
func (this *ZSetOperation) PopMax(key string, count ...int64) ([]redis.Z, error) {
	return this.cmdable.ZPopMax(this.wrapKeyFunc(key), count...).Result()
}
func (this *ZSetOperation) PopMin(key string, count ...int64) ([]redis.Z, error) {
	return this.cmdable.ZPopMin(this.wrapKeyFunc(key), count...).Result()
}
func (this *ZSetOperation) Range(key string, start, stop int64) ([]string, error) {
	return this.cmdable.ZRange(this.wrapKeyFunc(key), start, stop).Result()
}
func (this *ZSetOperation) RangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return this.cmdable.ZRangeWithScores(this.wrapKeyFunc(key), start, stop).Result()
}
func (this *ZSetOperation) RangeByScore(key string, opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRangeByScore(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) RangeByLex(key string, opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRangeByLex(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) RangeByScoreWithScores(key string, opt redis.ZRangeBy) ([]redis.Z, error) {
	return this.cmdable.ZRangeByScoreWithScores(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) Rank(key, member string) (int64, error) {
	return this.cmdable.ZRank(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) Rem(key string, members ...interface{}) (int64, error) {
	return this.cmdable.ZRem(this.wrapKeyFunc(key), members...).Result()
}
func (this *ZSetOperation) RemRangeByRank(key string, start, stop int64) (int64, error) {
	return this.cmdable.ZRemRangeByRank(this.wrapKeyFunc(key), start, stop).Result()
}
func (this *ZSetOperation) RemRangeByScore(key, min, max string) (int64, error) {
	return this.cmdable.ZRemRangeByScore(this.wrapKeyFunc(key), min, max).Result()
}
func (this *ZSetOperation) RemRangeByLex(key, min, max string) (int64, error) {
	return this.cmdable.ZRemRangeByLex(this.wrapKeyFunc(key), min, max).Result()
}
func (this *ZSetOperation) RevRange(key string, start, stop int64) ([]string, error) {
	return this.cmdable.ZRevRange(this.wrapKeyFunc(key), start, stop).Result()
}
func (this *ZSetOperation) RevRangeWithScores(key string, start, stop int64) ([]redis.Z, error) {
	return this.cmdable.ZRevRangeWithScores(this.wrapKeyFunc(key), start, stop).Result()
}
func (this *ZSetOperation) RevRangeByScore(key string, opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRevRangeByScore(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) RevRangeByLex(key string, opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRevRangeByLex(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) RevRangeByScoreWithScores(key string, opt redis.ZRangeBy) ([]redis.Z, error) {
	return this.cmdable.ZRevRangeByScoreWithScores(this.wrapKeyFunc(key), opt).Result()
}
func (this *ZSetOperation) RevRank(key, member string) (int64, error) {
	return this.cmdable.ZRevRank(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) Score(key, member string) (float64, error) {
	return this.cmdable.ZScore(this.wrapKeyFunc(key), member).Result()
}
func (this *ZSetOperation) UnionStore(dest string, store redis.ZStore, keys ...string) (int64, error) {
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

func (this *BoundZSetOperation) Add(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAdd(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) AddNX(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddNX(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) AddXX(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddXX(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) AddCh(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddCh(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) AddNXCh(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddNXCh(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) AddXXCh(members ...redis.Z) (int64, error) {
	return this.cmdable.ZAddXXCh(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) Incr(member redis.Z) (float64, error) {
	return this.cmdable.ZIncr(this.boundKey, member).Result()
}
func (this *BoundZSetOperation) IncrNX(member redis.Z) (float64, error) {
	return this.cmdable.ZIncrNX(this.boundKey, member).Result()
}
func (this *BoundZSetOperation) IncrXX(member redis.Z) (float64, error) {
	return this.cmdable.ZIncrXX(this.boundKey, member).Result()
}
func (this *BoundZSetOperation) Card() (int64, error) {
	return this.cmdable.ZCard(this.boundKey).Result()
}
func (this *BoundZSetOperation) Count(min, max string) (int64, error) {
	return this.cmdable.ZCount(this.boundKey, min, max).Result()
}
func (this *BoundZSetOperation) LexCount(min, max string) (int64, error) {
	return this.cmdable.ZLexCount(this.boundKey, min, max).Result()
}
func (this *BoundZSetOperation) IncrBy(increment float64, member string) (float64, error) {
	return this.cmdable.ZIncrBy(this.boundKey, increment, member).Result()
}
func (this *BoundZSetOperation) PopMax(count ...int64) ([]redis.Z, error) {
	return this.cmdable.ZPopMax(this.boundKey, count...).Result()
}
func (this *BoundZSetOperation) PopMin(count ...int64) ([]redis.Z, error) {
	return this.cmdable.ZPopMin(this.boundKey, count...).Result()
}
func (this *BoundZSetOperation) Range(start, stop int64) ([]string, error) {
	return this.cmdable.ZRange(this.boundKey, start, stop).Result()
}
func (this *BoundZSetOperation) RangeWithScores(start, stop int64) ([]redis.Z, error) {
	return this.cmdable.ZRangeWithScores(this.boundKey, start, stop).Result()
}
func (this *BoundZSetOperation) RangeByScore(opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRangeByScore(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) RangeByLex(opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRangeByLex(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) RangeByScoreWithScores(opt redis.ZRangeBy) ([]redis.Z, error) {
	return this.cmdable.ZRangeByScoreWithScores(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) Rank(member string) (int64, error) {
	return this.cmdable.ZRank(this.boundKey, member).Result()
}
func (this *BoundZSetOperation) Rem(members ...interface{}) (int64, error) {
	return this.cmdable.ZRem(this.boundKey, members...).Result()
}
func (this *BoundZSetOperation) RemRangeByRank(start, stop int64) (int64, error) {
	return this.cmdable.ZRemRangeByRank(this.boundKey, start, stop).Result()
}
func (this *BoundZSetOperation) RemRangeByScore(min, max string) (int64, error) {
	return this.cmdable.ZRemRangeByScore(this.boundKey, min, max).Result()
}
func (this *BoundZSetOperation) RemRangeByLex(min, max string) (int64, error) {
	return this.cmdable.ZRemRangeByLex(this.boundKey, min, max).Result()
}
func (this *BoundZSetOperation) RevRange(start, stop int64) ([]string, error) {
	return this.cmdable.ZRevRange(this.boundKey, start, stop).Result()
}
func (this *BoundZSetOperation) RevRangeWithScores(start, stop int64) ([]redis.Z, error) {
	return this.cmdable.ZRevRangeWithScores(this.boundKey, start, stop).Result()
}
func (this *BoundZSetOperation) RevRangeByScore(opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRevRangeByScore(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) RevRangeByLex(opt redis.ZRangeBy) ([]string, error) {
	return this.cmdable.ZRevRangeByLex(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) RevRangeByScoreWithScores(opt redis.ZRangeBy) ([]redis.Z, error) {
	return this.cmdable.ZRevRangeByScoreWithScores(this.boundKey, opt).Result()
}
func (this *BoundZSetOperation) RevRank(member string) (int64, error) {
	return this.cmdable.ZRevRank(this.boundKey, member).Result()
}
func (this *BoundZSetOperation) Score(member string) (float64, error) {
	return this.cmdable.ZScore(this.boundKey, member).Result()
}
