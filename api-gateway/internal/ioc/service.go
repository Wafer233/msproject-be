package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/grpc"
)

// add all the services (domain/application) here

func ProvideAuthService(clientMgr *grpc.GrpcClientManager) *service.AuthService {
	return service.NewAuthService(clientMgr.AuthClient)
}

func ProvideCaptchaService(clientMgr *grpc.GrpcClientManager) *service.CaptchaService {
	return service.NewCaptchaService(clientMgr.CaptchaClient)
}

func ProvideMenuService(clientMgr *grpc.GrpcClientManager) *service.MenuService {
	return service.NewMenuService(clientMgr.MenuClient)
}

func ProvideProjectService(clientMgr *grpc.GrpcClientManager) *service.ProjectService {
	return service.NewProjectService(clientMgr.ProjectClient)
}

func ProvideUserService(clientMgr *grpc.GrpcClientManager) *service.UserService {
	return service.NewUserService(clientMgr.UserClient)
}
