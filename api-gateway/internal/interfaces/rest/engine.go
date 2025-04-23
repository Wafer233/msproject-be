package rest

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/router"
	"github.com/gin-gonic/gin"
)

func InitWeb(middlewares []gin.HandlerFunc, ar *router.AuthRouter) *gin.Engine {
	engine := gin.Default()
	engine.Use(middlewares...)

	// 路由注册
	ar.Register(engine)

	return engine
}
