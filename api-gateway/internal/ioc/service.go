package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/grpc"
)

func ProvideAuthService(clientMgr *grpc.GrpcClientManager) *service.AuthService {
	return service.NewAuthService(clientMgr.AuthClient)
}

func ProvideCaptchaService(clientMgr *grpc.GrpcClientManager) *service.CaptchaService {
	return service.NewCaptchaService(clientMgr.CaptchaClient)
}
