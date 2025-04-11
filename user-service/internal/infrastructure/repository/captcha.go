package impl

import (
	"context"
	"fmt"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/cache"
	"time"
)

type CachedCaptchaRepository struct {
	RedisCache *cache.RedisCache
}

func NewCachedCaptchaRepository(redisClient *cache.RedisCache) repository.CaptchaRepository {
	return &CachedCaptchaRepository{
		RedisCache: redisClient,
	}
}

func (r *CachedCaptchaRepository) SaveCaptcha(ctx context.Context, mobile string, code string, expiration time.Duration) error {
	key := fmt.Sprintf("REGISTER_%s", mobile)
	return r.RedisCache.Put(ctx, key, code, expiration)
}

func (r *CachedCaptchaRepository) GetCaptcha(ctx context.Context, mobile string) (string, error) {
	key := fmt.Sprintf("REGISTER_%s", mobile)
	return r.RedisCache.Get(ctx, key)
}
