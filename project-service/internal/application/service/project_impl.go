package service

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/application/dto"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type DefaultProjectService struct {
	projectRepo repository.ProjectRepository
}

func NewDefaultProjectService(projectRepo repository.ProjectRepository) ProjectService {
	return &DefaultProjectService{projectRepo: projectRepo}
}

func (s *DefaultProjectService) GetProjectsByMemberId(ctx context.Context, memberId int64, page, pageSize int64) (*dto.ProjectListResponse, error) {
	// Call repository
	projects, total, err := s.projectRepo.FindProjectsByMemberId(ctx, memberId, page, pageSize)
	if err != nil {
		zap.L().Error("Failed to find projects by member ID", zap.Error(err))
		return nil, err
	}

	// Convert domain models to DTOs
	var projectDTOs []*dto.ProjectDTO
	for _, p := range projects {
		dto := &dto.ProjectDTO{}
		if err := copier.Copy(dto, p); err != nil {
			zap.L().Error("Failed to copy project to DTO", zap.Error(err))
			return nil, err
		}
		projectDTOs = append(projectDTOs, dto)
	}

	return &dto.ProjectListResponse{
		List:  projectDTOs,
		Total: total,
	}, nil
}
