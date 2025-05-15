package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/router"
)

func ProvideIndexRouter(handler *handler.IndexHandler) *router.IndexRouter {
	return router.NewIndexRouter(handler)
}

func ProvideLoginRouter(
	getCaptchaHandler *handler.GetCaptchaHandler,
	loginHandler *handler.LoginHandler,
	registerHandler *handler.RegisterHandler,
) *router.LoginRouter {
	return router.NewLoginRouter(getCaptchaHandler, loginHandler, registerHandler)
}

func ProvideOrganizationRouter(handler *handler.GetOrgListHandler) *router.OrganizationRouter {
	return router.NewOrganizationRouter(handler)
}

func ProvideProjectRouter(handler *handler.ProjectHandler) *router.ProjectRouter {
	return router.NewProjectRouter(handler)
}

// ProvideRouters 提供所有路由
func ProvideRouters(
	indexRouter *router.IndexRouter,
	loginRouter *router.LoginRouter,
	organizationRouter *router.OrganizationRouter,
	projectRouter *router.ProjectRouter,
	// 添加其他路由...
) []router.Router {
	return []router.Router{
		indexRouter,
		loginRouter,
		organizationRouter,
		projectRouter,
	}
}
