package ioc

import (
	"github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
	impl "github.com/Wafer233/msproject-be/project-service/internal/infrastructure/repository"
	"gorm.io/gorm"
)

func ProvideGORMMenuRepository(db *gorm.DB) repository.MenuRepository {
	return impl.NewGORMMenuRepository(db)
}

func ProvideGORMProjectRepository(db *gorm.DB) repository.ProjectRepository {
	return impl.NewGORMProjectRepository(db)
}
