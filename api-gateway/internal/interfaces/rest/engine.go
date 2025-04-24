package rest

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/router"
	"github.com/gin-gonic/gin"
)

func InitWeb(middlewares []gin.HandlerFunc, routers []router.Router) *gin.Engine {
	engine := gin.Default()
	engine.Use(middlewares...)

	// Register all routers
	for _, r := range routers {
		r.Register(engine)
	}

	return engine
}
