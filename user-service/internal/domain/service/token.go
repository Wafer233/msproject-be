package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

	zap.L().Info("生成token成功")
	return domainMember, domainOrganizations, nil
}

func (service *JWTTokenService) ParseToken(ctx context.Context, req *model.LoginReq, tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(service.accessSecret), nil
	})
	if err != nil {
		zap.L().Warn("jWT官方解析token失败")
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		val := claims["token"].(string)
		exp := int64(claims["exp"].(float64))
		if exp <= time.Now().Unix() {
			zap.L().Warn("token过期了")
			return "", errors.New("token过期了")
		}
		if claims["ip"] != req.Ip {
			zap.L().Warn("ip不合法")
			return "", errors.New("ip不合法")
		}
		return val, nil
	} else {
		return "", err
	}

}

func (service *JWTTokenService) GenerateToken(
	ctx context.Context,
	req *model.LoginReq,
	member *model.Member,
	organizations []*model.Organization,
) (*model.TokenPair, error) {

	memberIdStr := strconv.FormatInt(member.Id, 10)

	// Create access token
	accessExpTime := time.Duration(service.accessExp*3600*24) * time.Second
	accessExp := time.Now().Add(accessExpTime).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": memberIdStr,
		"exp":   accessExp,
		"ip":    req.Ip,
	})
	accessTokenStr, err := accessToken.SignedString([]byte(service.accessSecret))
	if err != nil {
		zap.L().Warn("调用token进入密钥失败")
		return nil, err
	}

	// Create refresh token
	refreshExpTime := time.Duration(service.refreshExp*3600*24) * time.Second
	refreshExp := time.Now().Add(refreshExpTime).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": memberIdStr,
		"exp":   refreshExp,
	})
	refreshTokenStr, err := refreshToken.SignedString([]byte(service.refreshSecret))
	if err != nil {
		zap.L().Warn("调用token更新密钥失败")
		return nil, err
	}

	// Async cache write
	go func() {
		if memberJSON, err := json.Marshal(member); err == nil {
			er := service.tokenRepo.Put(ctx, model.KeyMember+"::"+memberIdStr, string(memberJSON), accessExpTime)
			if er != nil {
				zap.L().Error("token缓存用户Id失败")
			}
		}
		if orgJSON, err := json.Marshal(organizations); err == nil {
			er := service.tokenRepo.Put(ctx, model.KeyOrganization+"::"+memberIdStr, string(orgJSON), accessExpTime)
			if er != nil {
				zap.L().Error("token缓存组织失败")
			}
		}
	}()

	zap.L().Info("生成token成功")
	return &model.TokenPair{
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
		AccessExp:    accessExp,
		RefreshExp:   refreshExp,
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
