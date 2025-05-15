package service

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/application/dto"
	"github.com/Wafer233/msproject-be/project-service/internal/application/dto/convert"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
	"go.uber.org/zap"
)

type DefaultMenuService struct {
	menuRepo repository.MenuRepository
}

func NewDefaultMenuService(menuRepo repository.MenuRepository) *DefaultMenuService {
	return &DefaultMenuService{
		menuRepo: menuRepo,
	}
}

// GetMenus 检索所有菜单并转换为树状结构
func (s *DefaultMenuService) GetMenus(ctx context.Context) (*dto.MenuResponse, error) {
	// 调用仓库获取领域模型
	menus, err := s.menuRepo.FindAll(ctx)
	if err != nil {
		zap.L().Error("获取菜单失败", zap.Error(err))
		return nil, err
	}

	// 使用领域逻辑转换为树状结构
	menuTree := model.ConvertToMenuTree(menus)

	// 将领域模型转换为DTO
	menuDTOs := convert.ToMenuDTOs(menuTree)

	// 返回包装在响应DTO中的结果
	return &dto.MenuResponse{
		Menus: menuDTOs,
	}, nil
}
