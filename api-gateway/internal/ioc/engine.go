package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/router"
	"github.com/gin-gonic/gin"
)

func ProvideMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Logger(),
		gin.Recovery(),
		// 你可以加入 CORS、JWT 等中间件
	}
}

func ProvideGinEngine(middlewares []gin.HandlerFunc, ar *router.AuthRouter) *gin.Engine {
	return rest.InitWeb(middlewares, ar)
}
