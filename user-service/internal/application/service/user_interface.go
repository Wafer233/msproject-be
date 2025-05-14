package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
)

type UserService interface {
	// GetOrgList 获取用户所有的组织
	GetOrgList(ctx context.Context, memberId int64) ([]dto.OrganizationDTO, error)
}
