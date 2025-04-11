package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/go-redis/redis/v8"
)

func ProvideRedisClient(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
}
