package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	pb "github.com/Wafer233/msproject-be/api-gateway/proto/login"
)

type GatewayRegisterService struct {
	client pb.LoginServiceClient
}

func NewGatewayRegisterService(client pb.LoginServiceClient) *GatewayRegisterService {
	return &GatewayRegisterService{
		client: client,
	}
}

func (service *GatewayRegisterService) Register(ctx context.Context, req dto.RegisterRequest) error {
	// 转换到gRPC请求
	grpcMsg := &pb.RegisterMessage{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
		Mobile:   req.Mobile,
		Captcha:  req.Captcha,
	}

	// 调用gRPC服务
	_, err := service.client.Register(ctx, grpcMsg)
	if err != nil {
		return err
	}

	return nil
}
