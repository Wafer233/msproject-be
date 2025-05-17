package router

import (
	"github.com/gin-gonic/gin"
)

func InitWeb(middlewares []gin.HandlerFunc, routers []Router) *gin.Engine {
	engine := gin.Default()
	engine.Use(middlewares...)

	// Register all routers
	for _, r := range routers {
		r.Register(engine)
	}

	return engine
}
