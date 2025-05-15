package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	pb "github.com/Wafer233/msproject-be/api-gateway/proto/index"
	"github.com/jinzhu/copier"
)

// GatewayMenuService 处理菜单相关的应用服务
type GatewayIndexService struct {
	client pb.IndexServiceClient
}

// NewMenuService 创建菜单服务
func NewGatewayIndexService(client pb.IndexServiceClient) *GatewayIndexService {
	return &GatewayIndexService{
		client: client,
	}
}

// GetMenus 获取菜单数据
func (s *GatewayIndexService) GetMenus(ctx context.Context) ([]*dto.IndexResponse, error) {
	// 调用gRPC服务
	grpcMsg := &pb.IndexMessage{}
	grpcResp, err := s.client.Index(ctx, grpcMsg)

	if err != nil {
		return nil, err
	}

	var dtoResp []*dto.IndexResponse

	err = copier.Copy(dtoResp, grpcResp)
	if err != nil {
		return nil, err
	}

	return dtoResp, err
}
