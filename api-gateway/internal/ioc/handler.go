package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
)

// add all the handlers here to support the router

func ProvideCaptchaHandler(svc *service.CaptchaService) *handler.CaptchaHandler {
	return handler.NewCaptchaHandler(svc)
}

func ProvideRegisterHandler(as *service.AuthService) *handler.RegisterHandler {
	return handler.NewRegisterHandler(as)
}

func ProvideLoginHandler(as *service.AuthService) *handler.LoginHandler {
	return handler.NewLoginHandler(as)
}

func ProvideMenuHandler(ms *service.MenuService) *handler.MenuHandler {
	return handler.NewMenuHandler(ms)
}
