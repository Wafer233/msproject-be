package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/router"
	"github.com/gin-gonic/gin"
)

// add middlewares here if needed, engine no need

func ProvideMiddlewares() []gin.HandlerFunc {
	return nil
}

func ProvideGinEngine(
	middlewares []gin.HandlerFunc,
	routers []router.Router,
) *gin.Engine {

	engine := rest.InitWeb(middlewares, routers)

	return engine
}
