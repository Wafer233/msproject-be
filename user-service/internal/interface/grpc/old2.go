package grpc

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	pb "github.com/Wafer233/msproject-be/user-service/proto/captcha"
)

type CaptchaServiceServer struct {
	pb.UnimplementedCaptchaServiceServer
	captchaService service.CaptchaService
}

func NewCaptchaServiceServer(captchaService service.CaptchaService) *CaptchaServiceServer {
	return &CaptchaServiceServer{
		captchaService: captchaService,
	}
}

func (s *CaptchaServiceServer) GetCaptcha(ctx context.Context, req *pb.GetCaptchaRequest) (*pb.GetCaptchaResponse, error) {
	// Call application service
	code, err := s.captchaService.GenerateCaptcha(ctx, req.Mobile)

	// Create response
	resp := &pb.GetCaptchaResponse{
		Success: err == nil,
	}

	if err != nil {
		resp.Message = err.Error()
	} else {
		resp.Code = code
	}

	return resp, nil
}
