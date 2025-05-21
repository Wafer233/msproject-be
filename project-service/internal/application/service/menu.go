package service

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
	repo "github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
	"go.uber.org/zap"
)

type IndexService interface {
	Index(ctx context.Context) ([]*model.ProjectMenu, error)
}

type DefaultIndexService struct {
	menuRepo repo.MenuRepo
}

func (service *DefaultIndexService) Index(ctx context.Context) ([]*model.ProjectMenu, error) {

	projMenus, err := service.menuRepo.GetAll(ctx)
	if err != nil {
		zap.L().Warn("memu仓库服务错误")
	}

	zap.L().Info("获取menu成功")
	return projMenus, nil
}

func NewDefaultIndexService(menuRepo repo.MenuRepo) IndexService {
	return &DefaultIndexService{
		menuRepo: menuRepo,
	}
}
