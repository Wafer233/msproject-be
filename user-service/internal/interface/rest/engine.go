package rest

import (
	"github.com/Wafer233/msproject-be/user-service/internal/interface/rest/router"
	"github.com/gin-gonic/gin"
)

func InitWeb(middlewares []gin.HandlerFunc, captchaRouter *router.CaptchaRouter) *gin.Engine {
	engine := gin.Default()
	engine.Use(middlewares...)

	// 路由注册
	captchaRouter.Register(engine)

	return engine
}
