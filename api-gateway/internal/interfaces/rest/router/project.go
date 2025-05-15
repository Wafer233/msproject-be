package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/gin-gonic/gin"
)

type ProjectRouter struct {
	projectHandler        *handler.ProjectHandler
	tokenVerifyMiddleware gin.HandlerFunc
}

func NewProjectRouter(
	projectHandler *handler.ProjectHandler,
	tokenVerifyMiddleware gin.HandlerFunc,
) *ProjectRouter {
	return &ProjectRouter{
		projectHandler:        projectHandler,
		tokenVerifyMiddleware: tokenVerifyMiddleware,
	}
}

// api-gateway/internal/interfaces/rest/router/project.go
func (router *ProjectRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// Protected routes with auth middleware
	protected := group.Group("")
	protected.Use(router.tokenVerifyMiddleware)

	group.POST("/project/selfList", router.projectHandler.SelfList)
	group.POST("/project", router.projectHandler.SelfList)
}
