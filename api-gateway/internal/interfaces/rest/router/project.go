package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/gin-gonic/gin"
)

type ProjectRouter struct {
	ph             *handler.ProjectHandler
	authMiddleware gin.HandlerFunc
}

func NewProjectRouter(ph *handler.ProjectHandler, authMiddleware gin.HandlerFunc) *ProjectRouter {
	return &ProjectRouter{
		ph:             ph,
		authMiddleware: authMiddleware,
	}
}

// api-gateway/internal/interfaces/rest/router/project.go
func (r *ProjectRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// Protected routes with auth middleware
	protected := group.Group("")
	protected.Use(r.authMiddleware)

	// 修改为明确的路径
	protected.POST("/project/selfList", r.ph.GetMyProjects) // 确保路径匹配
}
