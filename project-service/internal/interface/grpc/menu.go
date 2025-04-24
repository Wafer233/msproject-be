package grpc

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/application/dto"
	"github.com/Wafer233/msproject-be/project-service/internal/application/service"
	pb "github.com/Wafer233/msproject-be/project-service/proto/menu"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type MenuServiceServer struct {
	pb.UnimplementedMenuServiceServer
	ms service.MenuService
}

func NewMenuServiceServer(ms service.MenuService) *MenuServiceServer {
	return &MenuServiceServer{
		ms: ms,
	}
}

// Index 处理菜单导航请求
func (s *MenuServiceServer) Index(ctx context.Context, req *pb.IndexMessage) (*pb.IndexResponse, error) {
	// 获取菜单层次结构（现在返回的是DTO）
	menuResponse, err := s.ms.GetMenus(ctx)
	if err != nil {
		zap.L().Error("获取菜单失败", zap.Error(err))
		return nil, err
	}

	// 转换为protobuf格式
	var pbMenus []*pb.MenuMessage
	err = s.convertMenuTree(menuResponse.Menus, &pbMenus)
	if err != nil {
		zap.L().Error("转换菜单失败", zap.Error(err))
		return nil, err
	}

	return &pb.IndexResponse{
		Menus: pbMenus,
	}, nil
}

// convertMenuTree 递归转换菜单树为protobuf格式
func (s *MenuServiceServer) convertMenuTree(menuDTOs []*dto.MenuDTO, pbMenus *[]*pb.MenuMessage) error {
	for _, menuDTO := range menuDTOs {
		pbMenu := &pb.MenuMessage{}
		err := copier.Copy(pbMenu, menuDTO)
		if err != nil {
			return err
		}

		// 递归处理子菜单
		if len(menuDTO.Children) > 0 {
			var children []*pb.MenuMessage
			err = s.convertMenuTree(menuDTO.Children, &children)
			if err != nil {
				return err
			}
			pbMenu.Children = children
		}

		*pbMenus = append(*pbMenus, pbMenu)
	}

	return nil
}
