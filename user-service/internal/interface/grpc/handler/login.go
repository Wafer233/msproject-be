package handler

import (
	"context"
	"errors"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	domainSvc "github.com/Wafer233/msproject-be/user-service/internal/domain/service"
	pb "github.com/Wafer233/msproject-be/user-service/proto/login"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"time"
)

type LoginGRPCHandler struct {
	pb.UnimplementedLoginServiceServer
	captchaSvc      service.CaptchaService
	loginSvc        service.LoginService
	registerSvc     service.RegisterService
	tokenSvc        domainSvc.TokenService
	organizationSvc service.OrganizationService
}

func (handler *LoginGRPCHandler) GetCaptcha(ctx context.Context, pbReq *pb.GetCaptchaRequest) (*pb.GetCaptchaResponse, error) {
	ctx = context.Background()
	//1.获取参数
	mobile := pbReq.Mobile
	//2.校验参数

	zap.L().Info("短信平台调用成功，发送短信")

	//redis 假设后续缓存可能存在mysql当中，也可能存在mongo当中 也可能存在memcache当中
	//5.存储验证码 redis当中 过期时间15分钟
	code, err := handler.captchaSvc.GenerateCaptcha(ctx, mobile)
	if err != nil {
		zap.L().Warn("验证码存入redis出错")
	}

	pbResp := &pb.GetCaptchaResponse{
		Code: code,
	}

	zap.L().Info("验证码存入redis成功")
	return pbResp, nil
}

func (handler *LoginGRPCHandler) Register(ctx context.Context, pbReq *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	ctx = context.Background()
	var mobile, captcha string
	mobile = pbReq.Mobile
	captcha = pbReq.Captcha
	ok, err := handler.captchaSvc.ValidateCaptcha(ctx, mobile, captcha)
	if err == redis.Nil {
		zap.L().Warn("验证码服务验证redis错误")
		return nil, errors.New("验证码服务redis错误")
	}
	if err != nil {
		zap.L().Warn("验证码服务验证失败")
		return nil, errors.New("验证码服务未知错误")
	}
	if !ok {
		zap.L().Warn("验证码服务验证失败")
		return nil, errors.New("验证码错误")
	}

	domainReq := &model.RegisterReq{}

	err = copier.Copy(domainReq, pbReq)
	if err != nil {
		zap.L().Warn("注册请求复制错误")
		return nil, errors.New("copy失败")
	}
	err = handler.registerSvc.Register(ctx, domainReq)

	if err != nil {
		zap.L().Warn("注册微服务调用服务失败")
		return nil, errors.New("注册服务失败")
	}

	zap.L().Info("注册微服务调用服务成功")
	return &pb.RegisterResponse{}, nil

}

func (handler *LoginGRPCHandler) Login(ctx context.Context, pbReq *pb.LoginRequest) (*pb.LoginResponse, error) {
	ctx = context.Background()

	domainReq := &model.LoginReq{}

	err := copier.Copy(domainReq, pbReq)
	if err != nil {
		return nil, errors.New("copy失败")
	}

	domainMem, domainOrganizations, er := handler.loginSvc.Login(ctx, domainReq)
	if er != nil {
		zap.L().Error("Login service error", zap.Error(err))
		return nil, errors.New("login 服务失败")
	}
	if domainMem == nil {
		return nil, errors.New("member为空")
	}

	// pbMember
	pbMember := &pb.Member{}
	err = copier.Copy(pbMember, domainMem)
	pbMember.Code, _ = domainSvc.EncryptInt64(domainMem.Id, model.AESKey)
	pbMember.LastLoginTime = time.UnixMilli(domainMem.LastLoginTime).Format("2006-01-02 15:04:05")
	pbMember.CreateTime = time.UnixMilli(domainMem.CreateTime).Format("2006-01-02 15:04:05")

	//pbOrganizations
	var pbOrganizations []*pb.Organization
	err = copier.Copy(&pbOrganizations, domainOrganizations)
	for _, v := range pbOrganizations {
		v.Code, _ = domainSvc.EncryptInt64(v.Id, model.AESKey)
		v.OwnerCode = pbMember.Code
		o := model.OrgToMap(domainOrganizations)[v.Id]
		v.CreateTime = time.UnixMilli(o.CreateTime).Format("2006-01-02 15:04:05")

	}
	if len(domainOrganizations) > 0 {
		pbMember.OrganizationCode, _ = domainSvc.EncryptInt64(domainOrganizations[0].Id, model.AESKey)
	}

	//pbToken
	domainToken, err := handler.tokenSvc.GenerateToken(ctx, domainReq, domainMem, domainOrganizations)
	if err != nil {
		zap.L().Error("token服务失败", zap.Error(err))
		return nil, errors.New("token服务失败")
	}

	//可以给token做加密处理 增加安全性
	pbToken := &pb.Token{
		AccessToken:    domainToken.AccessToken,
		RefreshToken:   domainToken.RefreshToken,
		AccessTokenExp: domainToken.AccessExp,
		TokenType:      "bearer",
	}

	return &pb.LoginResponse{
		Member:           pbMember,
		OrganizationList: pbOrganizations,
		TokenList:        pbToken,
	}, nil

}

func (handler *LoginGRPCHandler) TokenVerify(ctx context.Context, pbReq *pb.TokenVerifyRequest) (*pb.TokenVerifyResponse, error) {

	domainReq := &model.LoginReq{}
	err := copier.Copy(domainReq, pbReq)
	if err != nil {
		return nil, errors.New("copy失败")
	}

	domainMember, domainOrganization, err := handler.tokenSvc.ValidateToken(ctx, domainReq)
	if err != nil {
		zap.L().Error("token服务失败", zap.Error(err))
		return nil, errors.New("token服务失败")
	}
	pbMember := &pb.Member{}

	err = copier.Copy(pbMember, domainMember)
	if err != nil {
		return nil, errors.New("copy失败")
	}

	pbMember.Code, _ = domainSvc.EncryptInt64(domainMember.Id, model.AESKey)

	if len(domainOrganization) > 0 {
		pbMember.OrganizationCode, _ = domainSvc.EncryptInt64(domainOrganization[0].Id, model.AESKey)
	}
	pbMember.CreateTime = time.UnixMilli(domainMember.CreateTime).Format("2006-01-02 15:04:05")
	pbResp := &pb.TokenVerifyResponse{
		Member: pbMember,
	}

	return pbResp, nil

}

func (handler *LoginGRPCHandler) GetOrgList(ctx context.Context, pbReq *pb.GetOrgListRequest) (*pb.GetOrgListResponse, error) {
	memberId := pbReq.MemberId

	organizations, err := handler.organizationSvc.GetOrganizations(ctx, memberId)
	if err != nil {
		zap.L().Error("获取组织列表失败", zap.Error(err))
		return nil, errors.New("获取组织列表失败")
	}

	var pbOrganizations []*pb.Organization
	err = copier.Copy(&pbOrganizations, organizations)
	if err != nil {
		zap.L().Error("copy失败", zap.Error(err))
		return nil, errors.New("copy失败")
	}
	for _, v := range pbOrganizations {
		v.Code, _ = domainSvc.EncryptInt64(v.Id, model.AESKey)
		o := model.OrgToMap(organizations)[v.Id]
		v.CreateTime = time.UnixMilli(o.CreateTime).Format("2006-01-02 15:04:05")
	}
	return &pb.GetOrgListResponse{
		OrganizationList: pbOrganizations,
	}, nil

}

func NewLoginGRPCHandler(
	captchaSvc service.CaptchaService,
	loginSvc service.LoginService,
	registerSvc service.RegisterService,
	tokenSvc domainSvc.TokenService,
	organizationSvc service.OrganizationService,
) *LoginGRPCHandler {
	return &LoginGRPCHandler{
		captchaSvc:      captchaSvc,
		loginSvc:        loginSvc,
		registerSvc:     registerSvc,
		tokenSvc:        tokenSvc,
		organizationSvc: organizationSvc,
	}
}
