package repository

import (
	"context"
	"fmt"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/cache"
	"time"
)

type RedisCaptchaRepository struct {
	redisClient *cache.RedisClient
}

func NewRedisCaptchaRepository(redisClient *cache.RedisClient) repository.CaptchaRepository {
	return &RedisCaptchaRepository{
		redisClient: redisClient,
	}
}

func (r *RedisCaptchaRepository) SaveCaptcha(ctx context.Context, mobile string, code string, expiration time.Duration) error {
	key := fmt.Sprintf("REGISTER_%s", mobile)
	return r.redisClient.Put(ctx, key, code, expiration)
}

func (r *RedisCaptchaRepository) GetCaptcha(ctx context.Context, mobile string) (string, error) {
	key := fmt.Sprintf("REGISTER_%s", mobile)
	return r.redisClient.Get(ctx, key)
}
