package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
)

type UserService interface {
	// 根据会员ID获取组织列表
	GetOrganizationsByMemberId(ctx context.Context, memberId int64) ([]model.Organization, error)
}
