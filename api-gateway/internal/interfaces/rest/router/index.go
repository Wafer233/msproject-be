package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/handler"
	"github.com/gin-gonic/gin"
)

// IndexRouter 菜单路由
type IndexRouter struct {
	indexHandler          *handler.IndexHandler
	tokenVerifyMiddleware gin.HandlerFunc
}

func NewIndexRouter(
	indexHandler *handler.IndexHandler,
	tokenVerifyMiddleware gin.HandlerFunc,
) *IndexRouter {
	return &IndexRouter{
		indexHandler:          indexHandler,
		tokenVerifyMiddleware: tokenVerifyMiddleware,
	}
}

// Register 注册路由
func (router *IndexRouter) Register(engine *gin.Engine) {
	group := engine.Group("/project")

	// 使用认证中间件保护路由
	protected := group.Group("")
	protected.Use(router.tokenVerifyMiddleware)

	// 添加导航菜单获取路由
	protected.POST("/index", router.indexHandler.Index)
}
