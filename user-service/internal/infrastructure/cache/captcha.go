package cache

import (
	"context"
	repo "github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCaptchaCache struct {
	client *redis.Client
}

func NewRedisCaptchaCache(client *redis.Client) repo.CaptchaRepo {
	return &RedisCaptchaCache{
		client: client,
	}
}

func (cache *RedisCaptchaCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	return cache.client.Set(ctx, key, value, expire).Err()
}

// Get 获取键值
func (cache *RedisCaptchaCache) Get(ctx context.Context, key string) (string, error) {
	return cache.client.Get(ctx, key).Result()
}
