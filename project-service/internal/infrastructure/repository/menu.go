package impl

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/project-service/internal/infrastructure/persistence/entity"
	"gorm.io/gorm"
)

type GORMMenuRepository struct {
	db *gorm.DB
}

func NewGORMMenuRepository(db *gorm.DB) repository.MenuRepository {
	return &GORMMenuRepository{db: db}
}

// FindAll 检索所有菜单项
func (r *GORMMenuRepository) FindAll(ctx context.Context) ([]*model.ProjectMenu, error) {
	var menuEntities []entity.ProjectMenuEntity
	err := r.db.WithContext(ctx).Find(&menuEntities).Error
	if err != nil {
		return nil, err
	}

	// 将实体转换为领域模型
	var menus []*model.ProjectMenu
	for _, e := range menuEntities {
		menu := e.ToModel()
		menus = append(menus, menu)
	}

	return menus, nil
}
