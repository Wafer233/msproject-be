package service

import (
	"context"
	"errors"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	domainService "github.com/Wafer233/msproject-be/user-service/internal/domain/service"
	"time"
)

type DefaultAuthService struct {
	mr repository.MemberRepository
	or repository.OrganizationRepository
	ps *domainService.PasswordService
	cr repository.CaptchaRepository
}

// NewAuthService 创建认证服务
func NewDefaultAuthService(
	mr repository.MemberRepository,
	or repository.OrganizationRepository,
	ps *domainService.PasswordService,
	cr repository.CaptchaRepository,
) AuthService {
	return &DefaultAuthService{
		mr: mr,
		or: or,
		ps: ps,
		cr: cr,
	}
}

// Register 用户注册
func (das *DefaultAuthService) Register(ctx context.Context, req dto.RegisterRequest) error {

	// 验证验证码
	code, err := das.cr.GetCaptcha(ctx, "REGISTER_"+req.Mobile)
	if err != nil {
		return errors.New("验证码获取失败")
	}
	if code != req.Captcha {
		return errors.New("验证码错误")
	}

	// 检查用户是否已存在
	exists, err := das.mr.FindMemberByAccount(ctx, req.Name)
	if err != nil {
		return errors.New("系统错误")
	}
	if exists {
		return errors.New("账号已存在")
	}

	// 创建新用户
	member := &model.Member{
		Account:       req.Name,
		Password:      das.ps.EncryptPassword(req.Password),
		Name:          req.Name,
		Mobile:        req.Mobile,
		Email:         req.Email,
		CreateTime:    time.Now().UnixMilli(),
		LastLoginTime: time.Now().UnixMilli(),
		Status:        1, // 默认启用
	}

	// 保存用户
	if err := das.mr.SaveMember(ctx, member); err != nil {
		return errors.New("注册失败")
	}

	// 创建个人组织
	org := &model.Organization{
		Name:       member.Name + "_organization",
		MemberId:   member.Id,
		CreateTime: time.Now().UnixMilli(),
		Personal:   1,
		Avatar:     "",
	}

	// 保存组织
	if err := das.or.SaveOrganization(ctx, org); err != nil {
		return errors.New("注册失败")
	}

	return nil
}
