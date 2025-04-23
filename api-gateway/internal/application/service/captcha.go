package service

import (
	"context"
	"errors"
	captchapb "github.com/Wafer233/msproject-be/api-gateway/proto/captcha"
)

type CaptchaService struct {
	client captchapb.CaptchaServiceClient
}

func NewCaptchaService(client captchapb.CaptchaServiceClient) *CaptchaService {
	return &CaptchaService{
		client: client,
	}
}

func (s *CaptchaService) GenerateCaptcha(ctx context.Context, mobile string) (string, error) {
	// 创建gRPC请求
	req := &captchapb.GetCaptchaRequest{
		Mobile: mobile,
	}

	// 调用gRPC服务
	resp, err := s.client.GetCaptcha(ctx, req)
	if err != nil {
		return "", err
	}

	// 检查响应
	if !resp.Success {
		return "", errors.New(resp.Message)
	}

	return resp.Code, nil
}
