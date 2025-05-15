package grpc

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	pb "github.com/Wafer233/msproject-be/user-service/proto/auth"
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

func (s *AuthServiceServer) Register(ctx context.Context, req *pb.RegisterMessage) (*pb.RegisterResponse, error) {
	// Convert proto request to DTO
	dtoReq := dto.RegisterRequest{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		Mobile:   req.Mobile,
		Captcha:  req.Captcha,
	}

	// Call application service
	err := s.authService.Register(ctx, dtoReq)

	if err != nil {
		return nil, err
	}
	// Create response
	resp := &pb.RegisterResponse{}

	return resp, nil
}

func (s *AuthServiceServer) Login(ctx context.Context, req *pb.LoginMessage) (*pb.LoginResponse, error) {
	// Convert proto request to DTO
	dtoReq := &dto.LoginReq{
		Account:  req.Account,
		Password: req.Password,
	}

	// Call application service
	dtoRsp, err := s.authService.Login(ctx, *dtoReq)
	if err != nil {
		return nil, err
	}

	// Convert DTO to proto response
	resp := &pb.LoginResponse{}

	// Convert member
	resp.Member = &pb.MemberMessage{
		Id:     dtoRsp.Member.Id,
		Name:   dtoRsp.Member.Name,
		Mobile: dtoRsp.Member.Mobile,
		Status: int32(dtoRsp.Member.Status),
		// 以下是可选字段，dto 中没有就不填
		// Realname: "",
		// Account:  "",
		// LastLoginTime: 0,
		// Address: "",
		// Province: 0,
		// City: 0,
		// Area: 0,
		// Email: "",
	}

	resp.TokenList = &pb.TokenMessage{
		AccessToken:    dtoRsp.TokenList.AccessToken,
		RefreshToken:   dtoRsp.TokenList.RefreshToken,
		TokenType:      dtoRsp.TokenList.TokenType,
		AccessTokenExp: dtoRsp.TokenList.AccessTokenExp,
	}

	// Convert organizations
	resp.OrganizationList = make([]*pb.OrganizationMessage, 0, len(dtoRsp.OrganizationList))
	for _, org := range dtoRsp.OrganizationList {
		orgDTO := &pb.OrganizationMessage{
			Id:          org.Id,
			Name:        org.Name,
			Avatar:      org.Avatar,
			Description: org.Description,
			MemberId:    org.MemberId,
			CreateTime:  org.CreateTime,
			Personal:    org.Personal,
			Address:     org.Address,
			Province:    org.Province,
			City:        org.City,
			Area:        org.Area,
		}
		resp.OrganizationList = append(resp.OrganizationList, orgDTO)
	}

	return resp, nil
}

func (s *AuthServiceServer) TokenVerify(ctx context.Context, req *pb.TokenVerifyRequest) (*pb.TokenVerifyResponse, error) {
	// Call application service
	tokenRsp, err := s.authService.TokenVerify(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	// Convert DTO to proto response
	resp := &pb.TokenVerifyResponse{
		Member: &pb.MemberMessage{
			Id:     tokenRsp.Id,
			Name:   tokenRsp.Name,
			Mobile: tokenRsp.Mobile,
			Status: int32(tokenRsp.Status),
		},
	}

	return resp, nil
}
