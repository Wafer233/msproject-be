package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	repo "github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

type TokenService interface {
	GenerateToken(ctx context.Context, req *model.LoginReq, member *model.Member, organizations []*model.Organization) (*model.TokenPair, error)
	ParseToken(ctx context.Context, req *model.LoginReq, tokenStr string) (string, error)
	ValidateToken(ctx context.Context, req *model.LoginReq) (*model.Member, []*model.Organization, error)
}

type JWTTokenService struct {
	accessExp     int64
	refreshExp    int64
	accessSecret  string
	refreshSecret string
	tokenRepo     repo.TokenRepo
}

func (service *JWTTokenService) ValidateToken(ctx context.Context, req *model.LoginReq) (*model.Member, []*model.Organization, error) {
	token := req.Token

	if strings.Contains(token, "bearer") {
		zap.L().Info("去除head的bearer成功")
		token = strings.ReplaceAll(token, "bearer ", "")
	}

	parseToken, err := service.ParseToken(ctx, req, token)

	if err != nil {
		zap.L().Warn("调用解析服务解析token失败")
		return nil, nil, err
	}

	memberJSON, err := service.tokenRepo.Get(context.Background(), model.KeyMember+"::"+parseToken)
	if err != nil {
		zap.L().Warn("调用解析获取token服务获取用户token失败")
		return nil, nil, err
	}
	if memberJSON == "" {
		zap.L().Warn("memberToken空")
		return nil, nil, errors.New("member空")
	}

	domainMember := &model.Member{}
	er := json.Unmarshal([]byte(memberJSON), domainMember)
	if er != nil {
		zap.L().Warn("JSON解析member失败")
		return nil, nil, er
	}
	orgsJson, err := service.tokenRepo.Get(context.Background(), model.KeyOrganization+"::"+parseToken)
	if err != nil {
		zap.L().Warn("调用解析获取token服务获取组织TOKEN失败")
		return domainMember, nil, err
	}
	if orgsJson == "" {
		zap.L().Info("组织空")
		return domainMember, nil, errors.New("组织空")
	}
	var domainOrganizations []*model.Organization
	er = json.Unmarshal([]byte(orgsJson), &domainOrganizations)
	if er != nil {
		zap.L().Warn("JSON解析组织失败")
		return nil, nil, er
	}

	zap.L().Info("验证token成功")
	return domainMember, domainOrganizations, nil
}

func (service *JWTTokenService) ParseToken(ctx context.Context, req *model.LoginReq, tokenStr string) (string, error) {
	claims := &model.CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.accessSecret), nil
	})

	if err != nil {
		zap.L().Warn("解析token失败", zap.Error(err))
		return "", err
	}

	if !token.Valid {
		zap.L().Warn("token无效")
		return "", errors.New("token无效")
	}

	zap.L().Info("token解析通过，签名&结构有效")

	// 校验 IP
	if claims.Ip != req.Ip {
		zap.L().Warn("ip不合法")
		return "", errors.New("ip不合法")
	}

	zap.L().Info("token全部校验通过，返回用户id")
	return claims.Token, nil
}

func (service *JWTTokenService) GenerateToken(
	ctx context.Context,
	req *model.LoginReq,
	member *model.Member,
	organizations []*model.Organization,
) (*model.TokenPair, error) {

	memberIdStr := strconv.FormatInt(member.Id, 10)

	// 计算 access token 的过期时间（精确到时间对象）
	accessExpDuration := time.Duration(service.accessExp*3600*24) * time.Second
	accessExpTime := time.Now().Add(accessExpDuration)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.CustomClaims{
		Token: memberIdStr,
		Ip:    req.Ip,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	accessTokenStr, err := accessToken.SignedString([]byte(service.accessSecret))
	if err != nil {
		zap.L().Warn("生成 accessToken 签名失败", zap.Error(err))
		return nil, err
	}

	// 计算 refresh token 的过期时间
	refreshExpDuration := time.Duration(service.refreshExp*3600*24) * time.Second
	refreshExpTime := time.Now().Add(refreshExpDuration)

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.CustomClaims{
		Token: memberIdStr,
		Ip:    req.Ip,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	refreshTokenStr, err := refreshToken.SignedString([]byte(service.refreshSecret))
	if err != nil {
		zap.L().Warn("生成 refreshToken 签名失败", zap.Error(err))
		return nil, err
	}

	// 异步写入缓存
	go func() {
		if memberJSON, err := json.Marshal(member); err == nil {
			er := service.tokenRepo.Put(ctx, model.KeyMember+"::"+memberIdStr, string(memberJSON), accessExpDuration)
			if er != nil {
				zap.L().Error("token缓存用户失败", zap.Error(er))
			}
		}
		if orgJSON, err := json.Marshal(organizations); err == nil {
			er := service.tokenRepo.Put(ctx, model.KeyOrganization+"::"+memberIdStr, string(orgJSON), accessExpDuration)
			if er != nil {
				zap.L().Error("token缓存组织失败", zap.Error(er))
			}
		}
	}()

	zap.L().Info("生成token成功")

	return &model.TokenPair{
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
		AccessExp:    accessExpTime.Unix(),
		RefreshExp:   refreshExpTime.Unix(),
	}, nil
}

func NewJWTTokenService(cfg *config.Config, tokenRepo repo.TokenRepo) TokenService {
	return &JWTTokenService{
		accessExp:     cfg.JWT.AccessExp,
		refreshExp:    cfg.JWT.RefreshExp,
		accessSecret:  cfg.JWT.AccessSecret,
		refreshSecret: cfg.JWT.RefreshSecret,
		tokenRepo:     tokenRepo,
	}
}
