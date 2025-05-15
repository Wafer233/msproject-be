package service

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/application/dto"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"strconv"
)

type DefaultProjectService struct {
	projectRepo repository.ProjectRepository
}

func NewDefaultProjectService(projectRepo repository.ProjectRepository) ProjectService {
	return &DefaultProjectService{projectRepo: projectRepo}
}

func (s *DefaultProjectService) GetProjectsByMemberId(ctx context.Context, memberId int64, page, pageSize int64) (*dto.ProjectListResponse, error) {
	// 调用仓库获取项目数据
	projects, total, err := s.projectRepo.FindProjectsByMemberId(ctx, memberId, page, pageSize)
	if err != nil {
		zap.L().Error("获取用户项目失败", zap.Error(err))
		return nil, err
	}

	// 转换领域模型到DTO
	var projectDTOs []*dto.ProjectDTO
	for _, p := range projects {
		projectDTO := &dto.ProjectDTO{}
		if err := copier.Copy(projectDTO, p); err != nil {
			zap.L().Error("复制项目数据失败", zap.Error(err))
			return nil, err
		}

		// 重要: 简化处理，直接将ID转为字符串赋值给Code
		projectDTO.Code = strconv.FormatInt(p.Id, 10)

		projectDTOs = append(projectDTOs, projectDTO)
	}

	return &dto.ProjectListResponse{
		List:  projectDTOs,
		Total: total,
	}, nil
}
