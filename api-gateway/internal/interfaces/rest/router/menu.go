package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/gin-gonic/gin"
)

// ProjectRouter 项目路由
type MenuRouter struct {
	ph             *handler.MenuHandler
	authMiddleware gin.HandlerFunc
}

func NewMenuRouter(ph *handler.MenuHandler, authMiddleware gin.HandlerFunc) *MenuRouter {
	return &MenuRouter{
		ph:             ph,
		authMiddleware: authMiddleware,
	}
}

// Register routes
func (r *MenuRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// Protected routes with auth middleware
	protected := group.Group("")
	protected.Use(r.authMiddleware)

	// Add protected routes
	protected.POST("/index", r.ph.Index)
}
