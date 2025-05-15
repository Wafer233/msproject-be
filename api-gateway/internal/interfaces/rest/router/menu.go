package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/gin-gonic/gin"
)

// MenuRouter 菜单路由
type MenuRouter struct {
	mh             *handler.MenuHandler
	authMiddleware gin.HandlerFunc
}

func NewMenuRouter(ph *handler.MenuHandler, authMiddleware gin.HandlerFunc) *MenuRouter {
	return &MenuRouter{
		mh:             ph,
		authMiddleware: authMiddleware,
	}
}

// Register 注册路由
func (r *MenuRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// 使用认证中间件保护路由
	protected := group.Group("")
	protected.Use(r.authMiddleware)

	// 添加导航菜单获取路由
	protected.POST("/index", r.mh.Index)
}
