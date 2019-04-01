package gredis

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisMessageQueue struct {
	queue  string
	cmdOpt redis.Cmdable
}

func NewRedisMessageQueue(queue string, cmdOpt redis.Cmdable) *RedisMessageQueue {
	return &RedisMessageQueue{queue: queue, cmdOpt: cmdOpt}
}

func (this *RedisMessageQueue) SetExpireDuration(duration time.Duration) error {
	ok, err := this.cmdOpt.Expire(this.queue, duration).Result()
	if ok {
		return nil
	}
	return err
}

func (this *RedisMessageQueue) Pop(num int) []string {
	val := this.cmdOpt.LPop(this.queue).Val()
	return []string{val}
}

func (this *RedisMessageQueue) Push(msg string) bool {
	return this.cmdOpt.RPush(this.queue, msg).Err() != nil
}

func (this *RedisMessageQueue) Name() string {
	return "redis"
}
