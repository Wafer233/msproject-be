package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest/handler"
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest/router"
)

func ProvideCaptchaRouter(hdl *handler.CaptchaHandler) *router.CaptchaRouter {
	return router.NewCaptchaRouter(hdl)
}
