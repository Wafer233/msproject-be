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
	// Create gRPC request
	req := &captchapb.GetCaptchaRequest{
		Mobile: mobile,
	}

	// Call gRPC service
	resp, err := s.client.GetCaptcha(ctx, req)
	if err != nil {
		return "", err
	}

	// Check response
	if !resp.Success {
		return "", errors.New(resp.Message)
	}

	return resp.Code, nil
}
