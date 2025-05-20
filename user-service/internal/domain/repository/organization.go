package repository

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
)

type OrganizationRepo interface {
	Save(ctx context.Context, org *model.Organization) error
	FindByMemberId(ctx context.Context, memberId int64) ([]*model.Organization, error)
}
