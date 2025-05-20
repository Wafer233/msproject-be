package ioc

import (
	repo "github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	dao "github.com/Wafer233/msproject-be/user-service/internal/infrastructure/repository"
	"gorm.io/gorm"
)

func ProvideGORMMemberRepository(db *gorm.DB) repo.MemberRepo {
	return dao.NewGORMMemberRepository(db)
}

func ProvideGORMOrganizationRepository(db *gorm.DB) repo.OrganizationRepo {
	return dao.NewGORMOrganizationRepository(db)
}
