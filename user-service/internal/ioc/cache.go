package ioc

import (
	repo "github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/cache"
	"github.com/go-redis/redis/v8"
)

func ProvideRedisTokenCache(client *redis.Client) repo.TokenRepo {
	return cache.NewRedisTokenCache(client)
}

func ProvideRedisCaptchaCache(client *redis.Client) repo.CaptchaRepo {
	return cache.NewRedisCaptchaCache(client)
}
