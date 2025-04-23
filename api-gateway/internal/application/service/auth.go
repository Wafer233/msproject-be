package service

import (
	"context"
	"errors"
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
	// Convert to gRPC request
	grpcReq := &authpb.RegisterRequest{
		Email:     req.Email,
		Name:      req.Name,
		Password:  req.Password,
		Password2: req.Password2,
		Mobile:    req.Mobile,
		Captcha:   req.Captcha,
	}

	// Call gRPC service
	resp, err := s.client.Register(ctx, grpcReq)
	if err != nil {
		return err
	}

	// Check response
	if !resp.Success {
		return errors.New(resp.Message)
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	// Convert to gRPC request
	grpcReq := &authpb.LoginRequest{
		Account:  req.Account,
		Password: req.Password,
	}

	// Call gRPC service
	grpcResp, err := s.client.Login(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// Convert to DTO
	resp := &dto.LoginResponse{}

	// Convert member
	resp.Member = dto.MemberDTO{
		Id:            grpcResp.Member.Id,
		Account:       grpcResp.Member.Account,
		Name:          grpcResp.Member.Name,
		Mobile:        grpcResp.Member.Mobile,
		Status:        int(grpcResp.Member.Status),
		LastLoginTime: grpcResp.Member.LastLoginTime,
		Email:         grpcResp.Member.Email,
		Avatar:        grpcResp.Member.Avatar,
	}

	// Convert token
	resp.TokenList = dto.TokenDTO{
		AccessToken:    grpcResp.TokenList.AccessToken,
		RefreshToken:   grpcResp.TokenList.RefreshToken,
		TokenType:      grpcResp.TokenList.TokenType,
		AccessTokenExp: grpcResp.TokenList.AccessTokenExp,
	}

	// Convert organizations
	resp.OrganizationList = make([]dto.OrganizationDTO, 0, len(grpcResp.OrganizationList))
	for _, org := range grpcResp.OrganizationList {
		var orgDTO dto.OrganizationDTO
		_ = copier.Copy(&orgDTO, org)
		resp.OrganizationList = append(resp.OrganizationList, orgDTO)
	}

	return resp, nil
}
