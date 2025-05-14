package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	orgpb "github.com/Wafer233/msproject-be/user-service/proto/user"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type OrganizationService interface {
	GetOrgList(ctx context.Context, memberId int64) ([]dto.OrganizationDTO, error)
}

type DefaultOrganizationService struct {
	client orgpb.OrganizationServiceClient
}

func NewOrganizationService(client orgpb.OrganizationServiceClient) OrganizationService {
	return &DefaultOrganizationService{
		client: client,
	}
}

func (s *DefaultOrganizationService) GetOrgList(ctx context.Context, memberId int64) ([]dto.OrganizationDTO, error) {
	// 创建请求
	req := &orgpb.OrgListRequest{
		MemberId: memberId,
	}

	// 调用 gRPC 服务
	resp, err := s.client.GetOrgList(ctx, req)
	if err != nil {
		zap.L().Error("调用组织服务失败", zap.Error(err))
		return nil, err
	}

	// 转换为 DTO
	var orgDTOs []dto.OrganizationDTO
	for _, pbOrg := range resp.OrganizationList {
		var orgDTO dto.OrganizationDTO
		if err := copier.Copy(&orgDTO, pbOrg); err != nil {
			zap.L().Error("组织数据转换失败", zap.Error(err))
			continue
		}
		orgDTOs = append(orgDTOs, orgDTO)
	}

	return orgDTOs, nil
}
