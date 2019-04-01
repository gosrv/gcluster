package gredis

import "github.com/go-redis/redis"

type SetOperation struct {
	wrapKeyFunc WrapKeyFunc
	cmdable     redis.Cmdable
}

func NewSetOperation(wrapKeyFunc WrapKeyFunc, cmdable redis.Cmdable) *SetOperation {
	return &SetOperation{wrapKeyFunc: wrapKeyFunc, cmdable: cmdable}
}

func (this *SetOperation) Add(key string, values ...string) (int64, error) {
	ivalues := make([]interface{}, len(values), len(values))
	for i := 0; i < len(values); i++ {
		ivalues[i] = values[i]
	}
	return this.cmdable.SAdd(this.wrapKeyFunc(key), ivalues...).Result()
}
func (this *SetOperation) Remove(key string, values ...string) (int64, error) {
	ivalues := make([]interface{}, len(values), len(values))
	for i := 0; i < len(values); i++ {
		ivalues[i] = values[i]
	}
	return this.cmdable.SRem(this.wrapKeyFunc(key), ivalues...).Result()
}
func (this *SetOperation) Pop(key string) (string, error) {
	return this.cmdable.SPop(this.wrapKeyFunc(key)).Result()
}
func (this *SetOperation) PopN(key string, count int64) ([]string, error) {
	return this.cmdable.SPopN(this.wrapKeyFunc(key), count).Result()
}
func (this *SetOperation) Move(key string, value string, destKey string) (bool, error) {
	return this.cmdable.SMove(this.wrapKeyFunc(key), this.wrapKeyFunc(destKey), value).Result()
}
func (this *SetOperation) Size(key string) (int64, error) {
	return this.cmdable.SCard(this.wrapKeyFunc(key)).Result()
}
func (this *SetOperation) IsMember(key string, o string) (bool, error) {
	return this.cmdable.SIsMember(this.wrapKeyFunc(key), o).Result()
}
func (this *SetOperation) Intersect(keys ...string) ([]string, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.SInter(wkeys...).Result()
}
func (this *SetOperation) IntersectAndStore(destination string, keys ...string) (int64, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.SInterStore(this.wrapKeyFunc(destination), wkeys...).Result()
}
func (this *SetOperation) Union(keys ...string) ([]string, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.SUnion(wkeys...).Result()
}
func (this *SetOperation) UnionAndStore(destination string, keys ...string) (int64, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.SUnionStore(this.wrapKeyFunc(destination), wkeys...).Result()
}
func (this *SetOperation) Difference(keys ...string) ([]string, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.SDiff(wkeys...).Result()
}
func (this *SetOperation) DifferenceAndStore(destination string, keys ...string) (int64, error) {
	wkeys := make([]string, len(keys), len(keys))
	for i := 0; i < len(keys); i++ {
		wkeys[i] = keys[i]
	}
	return this.cmdable.SDiffStore(this.wrapKeyFunc(destination), wkeys...).Result()
}
func (this *SetOperation) Members(key string) ([]string, error) {
	return this.cmdable.SMembers(this.wrapKeyFunc(key)).Result()
}
func (this *SetOperation) RandomMember(key string) (string, error) {
	return this.cmdable.SRandMember(this.wrapKeyFunc(key)).Result()
}
func (this *SetOperation) RandomMembers(key string, count int64) ([]string, error) {
	return this.cmdable.SRandMemberN(this.wrapKeyFunc(key), count).Result()
}
func (this *SetOperation) Scan(key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	return this.cmdable.SScan(this.wrapKeyFunc(key), cursor, match, count)
}

type BoundSetOperation struct {
	boundKey string
	cmdable  redis.Cmdable
}

func NewBoundSetOperation(boundKey string, cmdable redis.Cmdable) *BoundSetOperation {
	return &BoundSetOperation{boundKey: boundKey, cmdable: cmdable}
}

func (this *BoundSetOperation) Add(values ...string) (int64, error) {
	ivalues := make([]interface{}, len(values), len(values))
	for i := 0; i < len(values); i++ {
		ivalues[i] = values[i]
	}
	return this.cmdable.SAdd(this.boundKey, ivalues...).Result()
}
func (this *BoundSetOperation) Remove(values ...string) (int64, error) {
	ivalues := make([]interface{}, len(values), len(values))
	for i := 0; i < len(values); i++ {
		ivalues[i] = values[i]
	}
	return this.cmdable.SRem(this.boundKey, ivalues...).Result()
}
func (this *BoundSetOperation) Pop() (string, error) {
	return this.cmdable.SPop(this.boundKey).Result()
}
func (this *BoundSetOperation) PopN(count int64) ([]string, error) {
	return this.cmdable.SPopN(this.boundKey, count).Result()
}
func (this *BoundSetOperation) Size() (int64, error) {
	return this.cmdable.SCard(this.boundKey).Result()
}
func (this *BoundSetOperation) IsMember(o string) (bool, error) {
	return this.cmdable.SIsMember(this.boundKey, o).Result()
}
func (this *BoundSetOperation) Members() ([]string, error) {
	return this.cmdable.SMembers(this.boundKey).Result()
}
func (this *BoundSetOperation) RandomMember() (string, error) {
	return this.cmdable.SRandMember(this.boundKey).Result()
}
func (this *BoundSetOperation) RandomMembers(count int64) ([]string, error) {
	return this.cmdable.SRandMemberN(this.boundKey, count).Result()
}
func (this *BoundSetOperation) Scan(cursor uint64, match string, count int64) *redis.ScanCmd {
	return this.cmdable.SScan(this.boundKey, cursor, match, count)
}
