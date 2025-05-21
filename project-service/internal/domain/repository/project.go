package repository

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
)

type ProjectRepo interface {
	GetByMemberId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*model.ProjectAndMember, int64, error)
	GetCollectByMemId(ctx context.Context, memberId int64, page int64, size int64) ([]*model.ProjectAndMember, int64, error)
}
