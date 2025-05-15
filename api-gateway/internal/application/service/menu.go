package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	menupb "github.com/Wafer233/msproject-be/api-gateway/proto/menu"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

// MenuService 处理菜单相关的应用服务
type MenuService struct {
	client menupb.MenuServiceClient
}

// NewMenuService 创建菜单服务
func NewMenuService(client menupb.MenuServiceClient) *MenuService {
	return &MenuService{
		client: client,
	}
}

// GetMenus 获取菜单数据
func (s *MenuService) GetMenus(ctx context.Context, token string) (*dto.MenuResponse, error) {
	// 调用gRPC服务
	resp, err := s.client.Index(ctx, &menupb.IndexMessage{
		Token: token,
	})

	if err != nil {
		zap.L().Error("调用gRPC服务获取菜单失败", zap.Error(err))
		return nil, err
	}

	// 转换为DTO
	menuResponse := &dto.MenuResponse{}
	var menus []*dto.MenuDTO

	// 递归转换菜单树
	for _, menu := range resp.Menus {
		menuDTO := &dto.MenuDTO{}
		if err := copier.Copy(menuDTO, menu); err != nil {
			zap.L().Error("复制菜单数据失败", zap.Error(err))
			return nil, err
		}

		// 递归处理子菜单
		if len(menu.Children) > 0 {
			var children []*dto.MenuDTO
			if err := s.ConvertMenuTree(menu.Children, &children); err != nil {
				return nil, err
			}
			menuDTO.Children = children
		}

		menus = append(menus, menuDTO)
	}

	menuResponse.Menus = menus
	return menuResponse, nil
}

// convertMenuTree 递归转换菜单树
func (s *MenuService) ConvertMenuTree(pbMenus []*menupb.MenuMessage, menuDTOs *[]*dto.MenuDTO) error {
	for _, pbMenu := range pbMenus {
		menuDTO := &dto.MenuDTO{}
		if err := copier.Copy(menuDTO, pbMenu); err != nil {
			return err
		}

		// 递归处理子菜单
		if len(pbMenu.Children) > 0 {
			var children []*dto.MenuDTO
			if err := s.ConvertMenuTree(pbMenu.Children, &children); err != nil {
				return err
			}
			menuDTO.Children = children
		}

		*menuDTOs = append(*menuDTOs, menuDTO)
	}

	return nil
}
