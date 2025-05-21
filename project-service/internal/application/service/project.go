package service

import (
	"context"
	"errors"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
	repo "github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
	"go.uber.org/zap"
)

type ProjectService interface {
	SelfProject(ctx context.Context, req *model.SelfProjectRequest) ([]*model.ProjectAndMember, int64, error)
}

type DefaultProjectService struct {
	projectRepo repo.ProjectRepo
}

func (service *DefaultProjectService) SelfProject(ctx context.Context, req *model.SelfProjectRequest) ([]*model.ProjectAndMember, int64, error) {
	memberId := req.MemberID
	page := req.Page
	pageSize := req.PageSize
	var projAndMems []*model.ProjectAndMember
	var cProjAndMems []*model.ProjectAndMember
	var total int64
	var err error

	switch req.SelectBy {
	case "", "my":
		condition := "deleted=0 "
		projAndMems, total, err = service.projectRepo.GetByMemberId(ctx, memberId, condition, page, pageSize)
		if err != nil {
			zap.L().Warn("通过id查找项目失败")
			return nil, 0, errors.New("通过id查找项目失败")
		}

	case "archive":
		condition := "archive=1 "
		projAndMems, total, err = service.projectRepo.GetByMemberId(ctx, memberId, condition, page, pageSize)
		if err != nil {
			zap.L().Warn("通过id查找项目失败")
			return nil, 0, errors.New("通过id查找项目失败")
		}

	case "deleted":
		condition := "deleted=1 "
		projAndMems, total, err = service.projectRepo.GetByMemberId(ctx, memberId, condition, page, pageSize)
		if err != nil {
			zap.L().Warn("通过id查找项目失败")
			return nil, 0, errors.New("通过id查找项目失败")
		}

	case "collect":
		projAndMems, total, err = service.projectRepo.GetCollectByMemId(ctx, memberId, page, pageSize)
		for _, projAndMem := range projAndMems {
			projAndMem.Collected = model.Collected
		}
		return projAndMems, total, err
	}

	cProjAndMems, _, er := service.projectRepo.GetCollectByMemId(ctx, memberId, page, pageSize)
	if er != nil {
		zap.L().Error("project FindProjectByMemId::FindCollectProjectByMemId error", zap.Error(err))
		return nil, 0, errors.New("通过id查找收集项目失败")
	}
	var cMap = make(map[int64]*model.ProjectAndMember)
	for _, cProjAndMem := range cProjAndMems {
		cMap[cProjAndMem.Id] = cProjAndMem
	}
	for _, projAndMem := range projAndMems {
		if cMap[projAndMem.ProjectCode] != nil {
			projAndMem.Collected = model.Collected
		}
	}

	if projAndMems == nil {
		return []*model.ProjectAndMember{}, 0, nil
	}

	return projAndMems, total, nil
}

func NewDefaultProjectService(projectRepo repo.ProjectRepo) ProjectService {
	return &DefaultProjectService{
		projectRepo: projectRepo,
	}
}
