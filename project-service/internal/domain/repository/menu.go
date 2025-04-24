package repository

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
)

// MenuRepository 定义了处理ProjectMenu实体的操作
type MenuRepository interface {
	// FindAll 检索所有菜单项
	FindAll(ctx context.Context) ([]*model.ProjectMenu, error)
}
