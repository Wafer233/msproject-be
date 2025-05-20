package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	repo "github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	domainSvc "github.com/Wafer233/msproject-be/user-service/internal/domain/service"
	"go.uber.org/zap"
)

type LoginService interface {
	Login(ctx context.Context, req *model.LoginReq) (*model.Member, []*model.Organization, error)
}

type DefaultLoginService struct {
	memberRepo       repo.MemberRepo
	organizationRepo repo.OrganizationRepo
}

func (service *DefaultLoginService) Login(ctx context.Context, req *model.LoginReq) (*model.Member, []*model.Organization, error) {

	account := req.Account
	password := domainSvc.EncryptPassword(req.Password)
	domainMember, err := service.memberRepo.GetByCredentials(ctx, account, password)
	if err != nil {
		zap.L().Warn("通过账号密码查找用户失败")
		return nil, nil, err
	}

	domainOrganizations, err := service.organizationRepo.FindByMemberId(ctx, domainMember.Id)
	if err != nil {
		zap.L().Warn("通过用户id查找组织失败")
		return nil, nil, err
	}

	return domainMember, domainOrganizations, nil

}

func NewDefaultLoginService(memberRepo repo.MemberRepo, organizationRepo repo.OrganizationRepo) LoginService {
	return &DefaultLoginService{
		memberRepo:       memberRepo,
		organizationRepo: organizationRepo,
	}
}
