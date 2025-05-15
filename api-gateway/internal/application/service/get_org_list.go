package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	pb "github.com/Wafer233/msproject-be/api-gateway/proto/organization"
	"github.com/jinzhu/copier"
)

type GatewayGetOrgListService struct {
	client pb.OrganizationServiceClient
}

func NewGatewayGetOrgListService(client pb.OrganizationServiceClient) *GatewayGetOrgListService {
	return &GatewayGetOrgListService{
		client: client,
	}
}

func (s *GatewayGetOrgListService) GetOrgList(ctx context.Context, memberId int64) ([]*dto.GetOrgListResponse, error) {
	// 创建请求
	grpcMsg := &pb.GetOrgListMessage{
		MemId: memberId,
	}

	// 调用gRPC服务
	grpcResp, err := s.client.GetOrgList(ctx, grpcMsg)
	if err != nil {
		return nil, err
	}

	if grpcResp.OrganizationList == nil {
		return nil, nil
	}

	dtoResp := make([]*dto.GetOrgListResponse, 0)

	err = copier.Copy(&dtoResp, grpcResp.OrganizationList)
	if err != nil {
		return nil, err
	}

	return dtoResp, nil

}
