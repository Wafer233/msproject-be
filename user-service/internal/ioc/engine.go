package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest"
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest/router"
	"github.com/gin-gonic/gin"
)

func ProvideMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Logger(),
		gin.Recovery(),
		// 你可以加入 CORS、JWT 等中间件
	}
}

func ProvideEngine(middlewares []gin.HandlerFunc, captchaRouter *router.CaptchaRouter) *gin.Engine {
	return rest.InitWeb(middlewares, captchaRouter)
}
