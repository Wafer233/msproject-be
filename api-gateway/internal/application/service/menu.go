package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	menupb "github.com/Wafer233/msproject-be/api-gateway/proto/menu"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

// ProjectService 处理项目相关的服务
type MenuService struct {
	client menupb.MenuServiceClient
}

// NewProjectService 创建一个新的项目服务
func NewMenuService(client menupb.MenuServiceClient) *MenuService {
	return &MenuService{
		client: client,
	}
}

func (s *MenuService) GetMenus(ctx context.Context, token string) (*dto.MenuResponse, error) {
	// 调用gRPC服务
	resp, err := s.client.Index(ctx, &menupb.IndexMessage{
		Token: token,
	})

	if err != nil {
		zap.L().Error("调用gRPC服务,获取菜单导航失败", zap.Error(err))
		return nil, err
	}

	// 转换为DTO
	menuResponse := &dto.MenuResponse{}
	var menus []*dto.MenuDTO

	// 递归转换菜单树
	for _, menu := range resp.Menus {
		menuDTO := &dto.MenuDTO{}
		er := copier.Copy(menuDTO, menu)
		if er != nil {
			zap.L().Error("递归转换菜单树失败", zap.Error(er))
			return nil, er
		}

		// 递归处理子菜单
		if len(menu.Children) > 0 {
			var children []*dto.MenuDTO
			err = s.ConvertMenuTree(menu.Children, &children)
			if err != nil {
				zap.L().Error("递归处理子菜单失败", zap.Error(err))
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
		err := copier.Copy(menuDTO, pbMenu)
		if err != nil {
			return err
		}

		// 递归处理子菜单
		if len(pbMenu.Children) > 0 {
			var children []*dto.MenuDTO
			err = s.ConvertMenuTree(pbMenu.Children, &children)
			if err != nil {
				return err
			}
			menuDTO.Children = children
		}

		*menuDTOs = append(*menuDTOs, menuDTO)
	}

	return nil
}
