package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/cache"
	impl "github.com/Wafer233/msproject-be/user-service/internal/infrastructure/repository"
)

func ProvideCachedCaptchaRepo(redisClient *cache.RedisCache) repository.CaptchaRepository {
	return impl.NewCachedCaptchaRepository(redisClient)
}
