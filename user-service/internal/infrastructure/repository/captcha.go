package impl

import (
	"context"
	"fmt"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/cache"
	"time"
)

type RedisCaptchaRepository struct {
	rc *cache.RedisCache
}

func NewRedisCaptchaRepository(rc *cache.RedisCache) repository.CaptchaRepository {
	return &RedisCaptchaRepository{
		rc: rc,
	}
}

func (rcr *RedisCaptchaRepository) SaveCaptcha(ctx context.Context, mobile string, code string, expiration time.Duration) error {
	key := fmt.Sprintf("REGISTER_%s", mobile)
	return rcr.rc.Put(ctx, key, code, expiration)
}

func (rcr *RedisCaptchaRepository) GetCaptcha(ctx context.Context, mobile string) (string, error) {
	key := fmt.Sprintf("REGISTER_%s", mobile)
	return rcr.rc.Get(ctx, key)
}
