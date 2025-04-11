package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/cache"
	"github.com/go-redis/redis/v8"
)

func ProvideRedisCache(client *redis.Client) *cache.RedisCache {
	return cache.NewRedisCache(client)
}
