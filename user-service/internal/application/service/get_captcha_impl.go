package service

import (
	"context"
	pb "github.com/Wafer233/msproject-be/user-service/proto/login"
)

type DefaultGetCaptchaService struct {
	pb.UnimplementedLoginServiceServer
}

func (d DefaultGetCaptchaService) GetCaptcha(ctx context.Context, mobile string) (string, error) {
	//TODO implement me
	panic("implement me")
}
