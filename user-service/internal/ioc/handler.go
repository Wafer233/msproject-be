package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest/handler"
)

func ProvideCaptchaHandler(svc service.CaptchaService) *handler.CaptchaHandler {
	return handler.NewCaptchaHandler(svc)
}
