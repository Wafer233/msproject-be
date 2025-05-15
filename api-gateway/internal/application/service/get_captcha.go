package service

import (
	"context"
	pb "github.com/Wafer233/msproject-be/api-gateway/proto/login"
)

type GatewayGetCaptchaService struct {
	client pb.LoginServiceClient
}

func NewGatewayGetCaptchaService(client pb.LoginServiceClient) *GatewayGetCaptchaService {
	return &GatewayGetCaptchaService{
		client: client,
	}
}

func (service *GatewayGetCaptchaService) GetCaptcha(ctx context.Context, mobile string) (string, error) {
	// 创建gRPC请求
	grpcMsg := &pb.GetCaptchaMessage{
		Mobile: mobile,
	}

	// 调用gRPC服务
	grpcResp, err := service.client.GetCaptcha(ctx, grpcMsg)
	if err != nil {
		return "", err
	}

	return grpcResp.Captcha, nil
}
