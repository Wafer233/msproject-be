package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	repo "github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
)

type OrganizationService interface {
	GetOrganizations(ctx context.Context, memId int64) ([]*model.Organization, error)
}

type DefaultOrganizationService struct {
	repo repo.OrganizationRepo
}

func (service *DefaultOrganizationService) GetOrganizations(ctx context.Context, memId int64) ([]*model.Organization, error) {
	domainOrganization, err := service.repo.FindByMemberId(ctx, memId)
	if err != nil {
		return nil, err
	}
	return domainOrganization, nil

}

func NewDefaultOrganizationService(repo repo.OrganizationRepo) OrganizationService {
	return &DefaultOrganizationService{
		repo: repo,
	}
}
