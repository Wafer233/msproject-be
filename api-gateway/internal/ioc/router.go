package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/handler"
	"github.com/Wafer233/msproject-be/api-gateway/internal/middleware"
	"github.com/Wafer233/msproject-be/api-gateway/internal/router"
)

func ProvideUserRouter(
	handler *handler.LoginHttpHandler,
	middleware *middleware.TokenVerifyMiddleware,
) *router.UserRouter {
	return router.NewUserRouter(handler, middleware)
}

func ProvideProjectRouter(
	projHandler *handler.ProjectHttpHandler,
	middleware *middleware.TokenVerifyMiddleware,
) *router.ProjectRouter {
	return router.NewProjectRouter(projHandler, middleware)
}

// ProvideRouters 提供所有路由
func ProvideRouters(
	userRouter *router.UserRouter,
	projectRouter *router.ProjectRouter,
) []router.Router {
	return []router.Router{
		userRouter,
		projectRouter,
	}
}
