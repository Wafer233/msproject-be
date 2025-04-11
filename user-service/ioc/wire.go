package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/cache"
	repo "github.com/Wafer233/msproject-be/user-service/internal/infrastructure/repository"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

// InitRedisProviderSet 提供Redis相关的依赖注入
var InitRedisProviderSet = wire.NewSet()

// ProvideRedisClient 提供Redis客户端
func ProvideRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.C.ReadRedisConfig().Host + ":" + config.C.Redis.Port,
		Password: config.C.Redis.Password,
		DB:       config.C.Redis.DB,
	})
}

// ProvideRedisCache 提供Redis缓存服务
func ProvideRedisCache(client *redis.Client) *cache.RedisClient {
	return cache.NewRedisClient(client)
}

// ProvideCaptchaRepository 提供验证码仓储
func ProvideCaptchaRepository(redisClient *cache.RedisClient) repository.CaptchaRepository {
	return repo.NewRedisCaptchaRepository(redisClient)
}
