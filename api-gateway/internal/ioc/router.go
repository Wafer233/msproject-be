package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/router"
	"github.com/gin-gonic/gin"
)

func ProvideAuthRouter(
	ch *handler.CaptchaHandler,
	lr *handler.LoginHandler,
	rh *handler.RegisterHandler,
) *router.AuthRouter {
	return router.NewAuthRouter(ch, lr, rh)
}

// ------------------- adding router -------------------
func ProvideMenuRouter(
	mh *handler.MenuHandler,
	authMiddleware gin.HandlerFunc,
) *router.MenuRouter {
	return router.NewMenuRouter(mh, authMiddleware)
}

func ProvideProjectRouter(
	ph *handler.ProjectHandler,
	authMiddleware gin.HandlerFunc,
) *router.ProjectRouter {
	return router.NewProjectRouter(ph, authMiddleware)
}

func ProvideRouters(
	authRouter *router.AuthRouter,
	menuRouter *router.MenuRouter,
	projectRouter *router.ProjectRouter,
	// add here
) []router.Router {
	return []router.Router{
		authRouter,
		menuRouter,
		projectRouter,
	}
}
