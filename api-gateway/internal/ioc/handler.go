package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/metrics"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
)

// add all the handlers here to support the router

func ProvideCaptchaHandler(svc *service.CaptchaService, mc *metrics.MetricsCollector) *handler.CaptchaHandler {
	return handler.NewCaptchaHandler(svc, mc)
}

func ProvideRegisterHandler(as *service.AuthService, mc *metrics.MetricsCollector) *handler.RegisterHandler {
	return handler.NewRegisterHandler(as, mc)
}

func ProvideLoginHandler(as *service.AuthService, mc *metrics.MetricsCollector) *handler.LoginHandler {
	return handler.NewLoginHandler(as, mc)
}

func ProvideMenuHandler(msvc *service.MenuService, mc *metrics.MetricsCollector) *handler.MenuHandler {
	return handler.NewMenuHandler(msvc, mc)
}
