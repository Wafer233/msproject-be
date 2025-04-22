package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/cache"
	impl "github.com/Wafer233/msproject-be/user-service/internal/infrastructure/repository"
	"gorm.io/gorm"
)

func ProvideCachedCaptchaRepo(redisClient *cache.RedisCache) repository.CaptchaRepository {
	return impl.NewCachedCaptchaRepository(redisClient)
}

func ProvideCachedMemberRepository(db *gorm.DB) repository.MemberRepository {
	return impl.NewCachedMemberRepository(db)
}

func ProvideCachedOrganizationRepository(db *gorm.DB) repository.OrganizationRepository {
	return impl.NewCachedOrganizationRepository(db)
}
