package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	projpb "github.com/Wafer233/msproject-be/api-gateway/proto/project"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"strconv"
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
	// 创建gRPC请求
	req := &projpb.ProjectListRequest{
		MemberId: memberId,
		Page:     page,
		PageSize: pageSize,
	}

	// 调用gRPC服务
	resp, err := s.client.GetProjectsByMemberId(ctx, req)
	if err != nil {
		zap.L().Error("调用项目服务失败", zap.Error(err))
		return nil, err
	}

	// 转换Proto消息到DTO
	var projectDTOs []*dto.ProjectDTO
	for _, pbProject := range resp.List {
		projectDTO := &dto.ProjectDTO{}
		if err := copier.Copy(projectDTO, pbProject); err != nil {
			zap.L().Error("复制项目数据失败", zap.Error(err))
			return nil, err
		}

		// 重要: 简化处理，直接将ID转为字符串赋值给Code
		projectDTO.Code = strconv.FormatInt(pbProject.Id, 10)

		projectDTOs = append(projectDTOs, projectDTO)
	}

	return &dto.ProjectListResponse{
		List:  projectDTOs,
		Total: resp.Total,
	}, nil
}
