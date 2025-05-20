package cache

import (
	"context"
	repo "github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisTokenCache struct {
	client *redis.Client
}

func NewRedisTokenCache(client *redis.Client) repo.TokenRepo {
	return &RedisTokenCache{
		client: client,
	}
}

func (cache *RedisTokenCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	return cache.client.Set(ctx, key, value, expire).Err()
}

// Get 获取键值
func (cache *RedisTokenCache) Get(ctx context.Context, key string) (string, error) {
	return cache.client.Get(ctx, key).Result()
}
