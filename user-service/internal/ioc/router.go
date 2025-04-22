package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest/handler"
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest/router"
)

func ProvideAuthRouter(
	ch *handler.CaptchaHandler,
	lr *handler.LoginHandler,
	rh *handler.RegisterHandler,
) *router.AuthRouter {
	return router.NewAuthRouter(ch, lr, rh)
}
