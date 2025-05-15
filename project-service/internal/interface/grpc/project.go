package grpc

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/application/service"
	pb "github.com/Wafer233/msproject-be/project-service/proto/project"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"strconv"
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
	// 调用应用服务
	response, err := s.projectService.GetProjectsByMemberId(ctx, req.MemberId, req.Page, req.PageSize)
	if err != nil {
		zap.L().Error("获取用户项目失败", zap.Error(err))
		return nil, err
	}

	// 转换DTO到Proto
	var pbProjects []*pb.ProjectMessage
	for _, dto := range response.List {
		pbProject := &pb.ProjectMessage{}
		if err := copier.Copy(pbProject, dto); err != nil {
			zap.L().Error("复制项目数据失败", zap.Error(err))
			return nil, err
		}

		// 重要: 简化处理，直接将ID转为字符串赋值给Code
		pbProject.Code = strconv.FormatInt(pbProject.Id, 10)

		pbProjects = append(pbProjects, pbProject)
	}

	return &pb.ProjectListResponse{
		List:  pbProjects,
		Total: response.Total,
	}, nil
}
