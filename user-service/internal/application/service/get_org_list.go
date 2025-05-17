package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
)

type GetOrgListService interface {
	GetOrgList(ctx context.Context, memberId int64) ([]*dto.GetOrgListResponse, error)
}
