package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	userpb "github.com/Wafer233/msproject-be/user-service/proto/user"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type UserService struct {
	client userpb.UserServiceClient
}

func NewUserService(client userpb.UserServiceClient) *UserService {
	return &UserService{
		client: client,
	}
}

func (s *UserService) GetOrgList(ctx context.Context, memberId int64) ([]dto.OrganizationDTO, error) {
	// 调用用户服务获取该成员的组织列表
	resp, err := s.client.GetOrganizationsByMemberId(ctx, &userpb.GetOrgListRequest{
		MemberId: memberId,
	})

	if err != nil {
		zap.L().Error("无法从服务获取组织", zap.Error(err))
		return nil, err
	}

	// 转换响应为DTO
	var orgDTOs []dto.OrganizationDTO
	for _, pbOrg := range resp.OrganizationList {
		orgDTO := dto.OrganizationDTO{}
		err := copier.Copy(&orgDTO, pbOrg)
		if err != nil {
			zap.L().Error("复制对象失败", zap.Error(err))
		}
		orgDTOs = append(orgDTOs, orgDTO)
	}

	return orgDTOs, nil
}
