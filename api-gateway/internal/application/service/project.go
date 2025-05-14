package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	projpb "github.com/Wafer233/msproject-be/api-gateway/proto/project"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type ProjectService struct {
	client projpb.ProjectServiceClient
}

func NewProjectService(client projpb.ProjectServiceClient) *ProjectService {
	return &ProjectService{
		client: client,
	}
}

func (s *ProjectService) GetMyProjects(ctx context.Context, memberId int64, page, pageSize int64) (*dto.ProjectListResponse, error) {
	// Create request
	req := &projpb.ProjectListRequest{
		MemberId: memberId,
		Page:     page,
		PageSize: pageSize,
	}

	// Call gRPC service
	resp, err := s.client.GetProjectsByMemberId(ctx, req)
	if err != nil {
		zap.L().Error("Failed to get projects from service", zap.Error(err))
		return nil, err
	}

	// Convert proto message to DTO
	var projectDTOs []*dto.ProjectDTO
	for _, pbProject := range resp.List {
		projectDTO := &dto.ProjectDTO{}
		if err := copier.Copy(projectDTO, pbProject); err != nil {
			zap.L().Error("Failed to copy proto message to DTO", zap.Error(err))
			return nil, err
		}
		projectDTOs = append(projectDTOs, projectDTO)
	}

	return &dto.ProjectListResponse{
		List:  projectDTOs,
		Total: resp.Total,
	}, nil
}
