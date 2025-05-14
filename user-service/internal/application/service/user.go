package service

import (
	"context"
	"errors"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
)

type DefaultUserService struct {
	or repository.OrganizationRepository
}

func NewDefaultUserService(or repository.OrganizationRepository) UserService {
	return &DefaultUserService{
		or: or,
	}
}

func (dus *DefaultUserService) GetOrganizationsByMemberId(ctx context.Context, memberId int64) ([]model.Organization, error) {
	// 从仓库获取组织
	organizations, err := dus.or.FindOrganizationsByMemberId(ctx, memberId)
	if err != nil {
		return nil, errors.New("获取组织信息失败")
	}

	return organizations, nil
}
