package grpc

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/application/service"
	pb "github.com/Wafer233/msproject-be/project-service/proto/project"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type ProjectServiceServer struct {
	pb.UnimplementedProjectServiceServer
	projectService service.ProjectService
}

func NewProjectServiceServer(projectService service.ProjectService) *ProjectServiceServer {
	return &ProjectServiceServer{
		projectService: projectService,
	}
}

func (s *ProjectServiceServer) GetProjectsByMemberId(ctx context.Context, req *pb.ProjectListRequest) (*pb.ProjectListResponse, error) {
	// Call application service
	response, err := s.projectService.GetProjectsByMemberId(ctx, req.MemberId, req.Page, req.PageSize)
	if err != nil {
		zap.L().Error("Failed to get projects by member ID", zap.Error(err))
		return nil, err
	}

	// Convert DTO to proto message
	var pbProjects []*pb.ProjectMessage
	for _, dto := range response.List {
		pbProject := &pb.ProjectMessage{}
		if err := copier.Copy(pbProject, dto); err != nil {
			zap.L().Error("Failed to copy DTO to proto message", zap.Error(err))
			return nil, err
		}
		pbProjects = append(pbProjects, pbProject)
	}

	return &pb.ProjectListResponse{
		List:  pbProjects,
		Total: response.Total,
	}, nil
}
