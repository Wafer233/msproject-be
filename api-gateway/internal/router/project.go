package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/handler"
	"github.com/Wafer233/msproject-be/api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectRouter struct {
	projHandler *handler.ProjectHttpHandler
	middleware  *middleware.TokenVerifyMiddleware
}

func NewProjectRouter(
	projHandler *handler.ProjectHttpHandler,
	middleware *middleware.TokenVerifyMiddleware,
) *ProjectRouter {
	return &ProjectRouter{
		projHandler: projHandler,
		middleware:  middleware,
	}
}

func (router *ProjectRouter) Register(engine *gin.Engine) {
	engine.Use(router.middleware.TokenVerify())

	engine.POST("/index", router.projHandler.Index)
	engine.POST("/project/selfList", router.projHandler.SelfProject)
	engine.POST("/project", router.projHandler.SelfProject)
}
