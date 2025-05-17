package grpc

import (
	appSvc "github.com/Wafer233/msproject-be/user-service/internal/application/service"
	domainSvc "github.com/Wafer233/msproject-be/user-service/internal/domain/service"
	pb "github.com/Wafer233/msproject-be/user-service/proto/login"
)

type LoginService struct {
	pb.UnimplementedLoginServiceServer
	getCaptchaService appSvc.GetCaptchaService
	service           appSvc.LoginService
	appSvc.RegisterService
	domainSvc.TokenVerifyService
}
