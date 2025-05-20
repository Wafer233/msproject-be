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
	group := engine.Group("/project")
	group.Use(router.middleware.TokenVerify())
	group.POST("/index", router.projHandler.Index)
	group.POST("/project/selfList", router.projHandler.SelfProject)
	group.POST("/project", router.projHandler.SelfProject)
}
