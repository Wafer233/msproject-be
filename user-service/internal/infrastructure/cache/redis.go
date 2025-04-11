package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// RedisCache 提供Redis基本操作
type RedisCache struct {
	client *redis.Client
}

// NewRedisClient 创建Redis客户端
func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

// Put 存储键值对
func (rc *RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return rc.client.Set(ctx, key, value, expire).Err()
}

// Get 获取键值
func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return rc.client.Get(ctx, key).Result()
}
