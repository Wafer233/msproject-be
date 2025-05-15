package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
)

func ProvideGetCaptchaHandler(service *service.GatewayGetCaptchaService) *handler.GetCaptchaHandler {
	return handler.NewGetCaptchaHandler(service)
}

func ProvideGetOrgListHandler(service *service.GatewayGetOrgListService) *handler.GetOrgListHandler {
	return handler.NewGetOrgListHandler(service)
}

func ProvideIndexHandler(service *service.GatewayIndexService) *handler.IndexHandler {
	return handler.NewIndexHandler(service)
}

func ProvideLoginHandler(service *service.GatewayLoginService) *handler.LoginHandler {
	return handler.NewLoginHandler(service)
}

func ProvideRegisterHandler(service *service.GatewayRegisterService) *handler.RegisterHandler {
	return handler.NewRegisterHandler(service)
}

func ProvideProjectHandler(service *service.GatewayProjectService) *handler.ProjectHandler {
	return handler.NewProjectHandler(service)
}
