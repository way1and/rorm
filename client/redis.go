package client

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

// Get 获取
func Get(client *redis.Client, key string) map[string]string {
	val, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil
	}
	return val
}

// GetEX 过期时间
func GetEX(client *redis.Client, key string) time.Duration {
	return client.TTL(ctx, key).Val()
}

// Exist 是否存在
func Exist(client *redis.Client, key string) bool {
	if client.Exists(ctx, key).Val() != 1 {
		return false
	}
	return true
}

// 设置
func set(client *redis.Client, key string, field string, v any) {
	client.HSet(ctx, key, field, v)
}

// SetEX 设置过期时间
func SetEX(client *redis.Client, key string, interval time.Duration) {
	client.Expire(ctx, key, interval)
}

func IncrBy(client *redis.Client, key string, field string, change int64) {
	client.HIncrBy(ctx, key, field, change)
}

func Sets(client *redis.Client, key string, v map[string]any) bool {
	err := client.HMSet(ctx, key, v).Err()
	if err != nil {
		return false
	}
	return true
}

func Del(client *redis.Client, key string) bool {
	err := client.Del(ctx, key).Err()
	if err != nil {
		return false
	}
	return true
}
