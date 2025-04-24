package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/router"
)

func ProvideAuthRouter(
	ch *handler.CaptchaHandler,
	lr *handler.LoginHandler,
	rh *handler.RegisterHandler,
) *router.AuthRouter {
	return router.NewAuthRouter(ch, lr, rh)
}

func ProvideMenuRouter(
	mh *handler.MenuHandler,
) *router.MenuRouter {
	return router.NewMenuRouter(mh)
}

func ProvideRouters(
	authRouter *router.AuthRouter,
	menuRouter *router.MenuRouter,
	// add here
) []router.Router {
	return []router.Router{
		authRouter,
		menuRouter,
	}
}
