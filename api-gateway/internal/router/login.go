package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/handler"
	"github.com/Wafer233/msproject-be/api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	handler    *handler.LoginHttpHandler
	middleware *middleware.TokenVerifyMiddleware
}

func NewUserRouter(
	handler *handler.LoginHttpHandler,
	middleware *middleware.TokenVerifyMiddleware,
) *UserRouter {
	return &UserRouter{
		handler:    handler,
		middleware: middleware,
	}
}

func (router *UserRouter) Register(engine *gin.Engine) {
	engine.POST("/project/login/getCaptcha", router.handler.GetCaptcha)
	engine.POST("/project/login/register", router.handler.Register)
	engine.POST("/project/login", router.handler.Login)

	organization := engine.Group("/project/organization")
	organization.Use(router.middleware.TokenVerify())
	organization.POST("_getOrgList", router.handler.GetOrgList)
}
