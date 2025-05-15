// auth.go
package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	authpb "github.com/Wafer233/msproject-be/api-gateway/proto/auth"
	"github.com/jinzhu/copier"
)

type AuthService struct {
	client authpb.AuthServiceClient
}

func NewAuthService(client authpb.AuthServiceClient) *AuthService {
	return &AuthService{
		client: client,
	}
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) error {
	// 转换到gRPC请求
	grpcReq := &authpb.RegisterMessage{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		Mobile:   req.Mobile,
		Captcha:  req.Captcha,
	}

	// 调用gRPC服务
	_, err := s.client.Register(ctx, grpcReq)
	if err != nil {
		return err
	}

	// 检查响应
	//if !resp.Success {
	//	return errors.New(resp.Message)
	//}

	return nil
}

func (s *AuthService) Login(ctx context.Context, req *dto.LoginReq) (*dto.LoginRsp, error) {
	// 转换到gRPC请求
	grpcReq := &authpb.LoginMessage{
		Account:  req.Account,
		Password: req.Password,
	}

	// 调用gRPC服务
	grpcResp, err := s.client.Login(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 转换到DTO
	resp := &dto.LoginRsp{}

	// 转换用户信息
	resp.Member = dto.Member{
		Id:     grpcResp.Member.Id,
		Name:   grpcResp.Member.Name,
		Mobile: grpcResp.Member.Mobile,
		Status: int(grpcResp.Member.Status),
	}

	// 转换令牌
	resp.TokenList = dto.TokenList{
		AccessToken:    grpcResp.TokenList.AccessToken,
		RefreshToken:   grpcResp.TokenList.RefreshToken,
		TokenType:      grpcResp.TokenList.TokenType,
		AccessTokenExp: grpcResp.TokenList.AccessTokenExp,
	}

	// 转换组织列表
	resp.OrganizationList = make([]dto.OrganizationList, 0, len(grpcResp.OrganizationList))
	for _, org := range grpcResp.OrganizationList {
		var orgDTO dto.OrganizationList
		_ = copier.Copy(&orgDTO, org)
		resp.OrganizationList = append(resp.OrganizationList, orgDTO)
	}

	return resp, nil
}
