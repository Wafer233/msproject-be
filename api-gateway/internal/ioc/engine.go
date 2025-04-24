package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/router"
	"github.com/gin-gonic/gin"
)

// add middlewares here if needed, engine no need

func ProvideMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Logger(),
		gin.Recovery(),
		// add here
	}
}

func ProvideGinEngine(middlewares []gin.HandlerFunc, routers []router.Router) *gin.Engine {
	return rest.InitWeb(middlewares, routers) // 直接传递所有路由器
}
