package service

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/application/dto"
)

//type MenuService interface {
//	GetMenus(ctx context.Context) ([]*model.ProjectMenuChild, error)
//}

type MenuService interface {
	GetMenus(ctx context.Context) (*dto.MenuResponse, error)
}
