package service

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/application/dto"
)

type ProjectService interface {
	GetProjectsByMemberId(ctx context.Context, memberId int64, page, pageSize int64) (*dto.ProjectListResponse, error)
}
