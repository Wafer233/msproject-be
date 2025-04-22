package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/cache"
	impl "github.com/Wafer233/msproject-be/user-service/internal/infrastructure/repository"
	"gorm.io/gorm"
)

func ProvideRedisCaptchaRepository(rc *cache.RedisCache) repository.CaptchaRepository {
	return impl.NewRedisCaptchaRepository(rc)
}

func ProvideGORMMemberRepository(db *gorm.DB) repository.MemberRepository {
	return impl.NewGORMMemberRepository(db)
}

func ProvideGORMOrganizationRepository(db *gorm.DB) repository.OrganizationRepository {
	return impl.NewGORMOrganizationRepository(db)
}
