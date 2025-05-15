package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/grpc"
)

// add all the services (domain/application) here

func ProvideGatewayGetCaptchaService(clientMgr *grpc.GrpcClientManager) *service.GatewayGetCaptchaService {
	return service.NewGatewayGetCaptchaService(clientMgr.LoginClient)
}

func ProvideGatewayGetOrgListService(clientMgr *grpc.GrpcClientManager) *service.GatewayGetOrgListService {
	return service.NewGatewayGetOrgListService(clientMgr.OrganizationClient)
}

func ProvideGatewayIndexService(clientMgr *grpc.GrpcClientManager) *service.GatewayIndexService {
	return service.NewGatewayIndexService(clientMgr.IndexClient)
}

func ProvideGatewayLoginService(clientMgr *grpc.GrpcClientManager) *service.GatewayLoginService {
	return service.NewGatewayLoginService(clientMgr.LoginClient)
}

func ProvideGatewayProjectService(clientMgr *grpc.GrpcClientManager) *service.GatewayProjectService {
	return service.NewGatewayProjectService(clientMgr.ProjectClient)
}

func ProvideGatewayRegisterService(clientMgr *grpc.GrpcClientManager) *service.GatewayRegisterService {
	return service.NewGatewayRegisterService(clientMgr.LoginClient)
}
