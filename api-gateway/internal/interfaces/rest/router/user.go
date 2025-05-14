package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	oh             *handler.UserHandler
	authMiddleware gin.HandlerFunc
}

func NewUserRouter(oh *handler.UserHandler, authMiddleware gin.HandlerFunc) *UserRouter {
	return &UserRouter{
		oh:             oh,
		authMiddleware: authMiddleware,
	}
}

func (r *UserRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// 受保护的路由与auth中间件
	protected := group.Group("")
	protected.Use(r.authMiddleware)

	// 添加受保护的路由
	protected.POST("/organization/_getOrgList", r.oh.GetOrgList)
}
