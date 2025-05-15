package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	orgpb "github.com/Wafer233/msproject-be/user-service/proto/user"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"strconv"
)

type OrganizationService struct {
	client orgpb.OrganizationServiceClient
}

func NewOrganizationService(client orgpb.OrganizationServiceClient) *OrganizationService {
	return &OrganizationService{
		client: client,
	}
}

func (s *OrganizationService) GetOrgList(ctx context.Context, memberId int64) ([]dto.OrganizationList, error) {
	// 创建请求
	req := &orgpb.OrgListRequest{
		MemberId: memberId,
	}

	// 调用gRPC服务
	resp, err := s.client.GetOrgList(ctx, req)
	if err != nil {
		zap.L().Error("调用组织服务失败", zap.Error(err))
		return nil, err
	}

	// 检查是否为空
	if resp.OrganizationList == nil {
		return []dto.OrganizationList{}, nil // 返回空数组而不是nil
	}

	// 确保初始化数组
	orgDTOs := make([]dto.OrganizationList, 0, len(resp.OrganizationList))

	// 转换为DTO
	for _, pbOrg := range resp.OrganizationList {
		var orgDTO dto.OrganizationList
		if err := copier.Copy(&orgDTO, pbOrg); err != nil {
			zap.L().Error("复制组织数据失败", zap.Error(err))
			continue
		}

		// 重要: 简化处理，直接将ID转为字符串赋值给Code
		orgDTO.Code = strconv.FormatInt(pbOrg.Id, 10)

		orgDTOs = append(orgDTOs, orgDTO)
	}

	return orgDTOs, nil
}
