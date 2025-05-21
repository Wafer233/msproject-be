package dao

import (
	"context"
	"fmt"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
	repo "github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/project-service/internal/infrastructure/persistence/entity"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GORMMenuDAO struct {
	db *gorm.DB
}

func (dao *GORMMenuDAO) GetAll(ctx context.Context) ([]*model.ProjectMenu, error) {

	var domainMenus []*model.ProjectMenu
	var entityMenus []*entity.ProjectMenu

	dao.db.Model(&entity.ProjectMenu{}).
		Find(&entityMenus)

	err := copier.Copy(&domainMenus, &entityMenus)
	if err != nil {
		zap.L().Warn("menu的领域与实体复制失败")
	}
	fmt.Println(domainMenus)
	return domainMenus, err
}

func NewGORMMenuRepo(db *gorm.DB) repo.MenuRepo {
	return &GORMMenuDAO{db: db}
}
