package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/gin-gonic/gin"
)

// ProjectRouter 项目路由
type MenuRouter struct {
	ph *handler.MenuHandler
}

func NewMenuRouter(ph *handler.MenuHandler) *MenuRouter {
	return &MenuRouter{
		ph: ph,
	}
}

// Register 注册路由
func (r *MenuRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// 添加首页路由
	group.POST("/index", r.ph.Index)
}
