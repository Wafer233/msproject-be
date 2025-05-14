package repository

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
)

type ProjectRepository interface {
	FindProjectsByMemberId(ctx context.Context, memberId int64, page, pageSize int64) ([]*model.ProjectWithMember, int64, error)
}
