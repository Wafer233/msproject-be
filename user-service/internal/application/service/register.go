package service

import (
	"context"
	"errors"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	repo "github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	domainSvc "github.com/Wafer233/msproject-be/user-service/internal/domain/service"
	"go.uber.org/zap"
	"time"
)

type RegisterService interface {
	Register(ctx context.Context, req *model.RegisterReq) error
}

type DefaultRegisterService struct {
	memberRepo       repo.MemberRepo
	organizationRepo repo.OrganizationRepo
}

func (service *DefaultRegisterService) Register(ctx context.Context, req *model.RegisterReq) error {
	//3.校验业务逻辑（邮箱是否被注册 账号是否被注册 手机号是否被注册）
	exist, err := service.memberRepo.ExistByEmail(ctx, req.Email)
	if err != nil {
		zap.L().Warn("从数据库查看邮箱存在失败")

		return errors.New("register db get error")
	}
	if exist {
		zap.L().Warn("有邮箱了")
		return errors.New("email already exists")
	}

	exist, err = service.memberRepo.ExistByAccount(ctx, req.Name)
	if err != nil {
		zap.L().Warn("从数据库查看用户名存在失败")
		return errors.New("db get error")
	}
	if exist {
		zap.L().Warn("有帐号了")
		return errors.New("account already exists")
	}

	exist, err = service.memberRepo.ExistByMobile(ctx, req.Mobile)
	if err != nil {
		zap.L().Warn("从数据库查看手机失败")
		return errors.New("db get error")
	}
	if exist {
		zap.L().Warn("有手机了")
		return errors.New("mobile already exists")
	}

	//4.执行业务 将数据存入member表 生成一个数据 存入组织表 organization

	domainMem := &model.Member{
		Account:       req.Name,
		Password:      domainSvc.EncryptPassword(req.Password),
		Name:          req.Name,
		Mobile:        req.Mobile,
		Email:         req.Email,
		CreateTime:    time.Now().UnixMilli(),
		LastLoginTime: time.Now().UnixMilli(),
		Status:        model.StatusNormal,
	}

	err = service.memberRepo.Save(ctx, domainMem)
	if err != nil {
		zap.L().Warn("保存用户失败")
		return errors.New("db SaveMember error")
	}
	//存入组织
	domainOrg := &model.Organization{
		Name:       domainMem.Name + "个人组织",
		MemberId:   domainMem.Id,
		CreateTime: time.Now().UnixMilli(),
		Personal:   model.StatusPersonal,
		Avatar:     "https://www.pixiv.net/artworks/128859426",
	}
	err = service.organizationRepo.Save(ctx, domainOrg)
	if err != nil {
		zap.L().Error("调用组织仓库保存失败")
		return errors.New("db SaveOrganization db err")
	}

	zap.L().Info("注册成功")
	return nil

}

func NewDefaultRegisterService(registerRepo repo.MemberRepo, organizationRepo repo.OrganizationRepo) RegisterService {
	return &DefaultRegisterService{
		memberRepo:       registerRepo,
		organizationRepo: organizationRepo,
	}
}
