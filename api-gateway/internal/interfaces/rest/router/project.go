package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectRouter struct {
	projectHandler *handler.ProjectHandler
}

func NewProjectRouter(projectHandler *handler.ProjectHandler) *ProjectRouter {
	return &ProjectRouter{
		projectHandler: projectHandler,
	}
}

// api-gateway/internal/interfaces/rest/router/project.go
func (router *ProjectRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// Protected routes with auth middleware
	protected := group.Group("")
	protected.Use(middleware.TokenVerifyMiddleware())

	group.POST("/project/selfList", router.projectHandler.SelfList)
	group.POST("/project", router.projectHandler.SelfList)
}
