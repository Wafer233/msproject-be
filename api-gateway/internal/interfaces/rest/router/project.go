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

func (r *ProjectRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// Protected routes with auth middleware
	protected := group.Group("")
	protected.Use(r.authMiddleware)

	// Add protected routes
	protected.POST("/project/selfList", r.ph.GetMyProjects)
}
