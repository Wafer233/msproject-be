package grpc

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	pb "github.com/Wafer233/msproject-be/user-service/proto/auth"
	"github.com/jinzhu/copier"
)

type AuthServiceServer struct {
	pb.UnimplementedAuthServiceServer
	authService service.AuthService
}

func NewAuthServiceServer(authService service.AuthService) *AuthServiceServer {
	return &AuthServiceServer{
		authService: authService,
	}
}

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Convert proto request to DTO
	dtoReq := dto.RegisterRequest{
		Email:     req.Email,
		Name:      req.Name,
		Password:  req.Password,
		Password2: req.Password2,
		Mobile:    req.Mobile,
		Captcha:   req.Captcha,
	}

	// Call application service
	err := s.authService.Register(ctx, dtoReq)

	// Create response
	resp := &pb.RegisterResponse{
		Success: err == nil,
	}

	if err != nil {
		resp.Message = err.Error()
	}

	return resp, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Convert proto request to DTO
	dtoReq := dto.LoginRequest{
		Account:  req.Account,
		Password: req.Password,
	}

	// Call application service
	result, err := s.authService.Login(ctx, dtoReq)
	if err != nil {
		return nil, err
	}

	// Convert DTO to proto response
	resp := &pb.LoginResponse{}

	// Convert member
	resp.Member = &pb.MemberDTO{
		Id:            result.Member.Id,
		Account:       result.Member.Account,
		Name:          result.Member.Name,
		Mobile:        result.Member.Mobile,
		Status:        int32(result.Member.Status),
		LastLoginTime: result.Member.LastLoginTime,
		Email:         result.Member.Email,
		Avatar:        result.Member.Avatar,
	}

	// Convert token
	resp.TokenList = &pb.TokenDTO{
		AccessToken:    result.TokenList.AccessToken,
		RefreshToken:   result.TokenList.RefreshToken,
		TokenType:      result.TokenList.TokenType,
		AccessTokenExp: result.TokenList.AccessTokenExp,
	}

	// Convert organizations
	resp.OrganizationList = make([]*pb.OrganizationDTO, 0, len(result.OrganizationList))
	for _, org := range result.OrganizationList {
		orgDTO := &pb.OrganizationDTO{}
		_ = copier.Copy(orgDTO, org)
		resp.OrganizationList = append(resp.OrganizationList, orgDTO)
	}

	return resp, nil
}
