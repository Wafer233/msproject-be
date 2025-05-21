package ioc

import (
	repo "github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
	dao "github.com/Wafer233/msproject-be/project-service/internal/infrastructure/repository"
	"gorm.io/gorm"
)

func ProvideGORMMenuRepo(db *gorm.DB) repo.MenuRepo {
	return dao.NewGORMMenuRepo(db)
}

func ProvideGORMProjectRepo(db *gorm.DB) repo.ProjectRepo {
	return dao.NewGORMProjectRepo(db)
}
