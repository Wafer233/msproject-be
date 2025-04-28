package service

import (
	"context"
	"errors"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	domainService "github.com/Wafer233/msproject-be/user-service/internal/domain/service"
	"github.com/jinzhu/copier"
	"time"
)

type DefaultAuthService struct {
	mr repository.MemberRepository
	or repository.OrganizationRepository
	ps *domainService.PasswordService
	cr repository.CaptchaRepository
	ts domainService.TokenService
}

// NewAuthService 创建认证服务
func NewDefaultAuthService(
	mr repository.MemberRepository,
	or repository.OrganizationRepository,
	ps *domainService.PasswordService,
	cr repository.CaptchaRepository,
	ts domainService.TokenService,
) AuthService {
	return &DefaultAuthService{
		mr: mr,
		or: or,
		ps: ps,
		cr: cr,
		ts: ts,
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

func (das *DefaultAuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	// 加密密码
	pwd := das.ps.EncryptPassword(req.Password)

	// 查找用户
	member, err := das.mr.FindMember(ctx, req.Account, pwd)
	if err != nil {
		return nil, errors.New("账号或密码错误")
	}

	// 检查用户状态
	if !member.IsValid() {
		return nil, errors.New("账号已被禁用")
	}

	// 查找组织
	organizations, err := das.or.FindOrganizationsByMemberId(ctx, member.Id)
	if err != nil {
		return nil, errors.New("获取组织信息失败")
	}

	// 生成令牌
	//accessToken, refreshToken, accessExp := das.ts.GenerateTokens(strconv.FormatInt(member.Id, 10))

	// 构建响应
	response := &dto.LoginResponse{
		Member: dto.MemberDTO{
			Id:            member.Id,
			Account:       member.Account,
			Name:          member.Name,
			Mobile:        member.Mobile,
			Status:        member.Status,
			LastLoginTime: member.LastLoginTime,
			Email:         member.Email,
			Avatar:        member.Avatar,
		},
		//TokenList: dto.TokenDTO{
		//	AccessToken:    accessToken,
		//	RefreshToken:   refreshToken,
		//	TokenType:      "bearer",
		//	AccessTokenExp: accessExp,
		//},
	}

	// 转换组织列表
	var orgDTOs []dto.OrganizationDTO
	for _, org := range organizations {
		var orgDTO dto.OrganizationDTO
		_ = copier.Copy(&orgDTO, org)
		orgDTOs = append(orgDTOs, orgDTO)
	}
	response.OrganizationList = orgDTOs

	return response, nil
}
