package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/config"
	"github.com/Wafer233/msproject-be/api-gateway/internal/metrics"
	"github.com/Wafer233/msproject-be/api-gateway/internal/router"
	"github.com/gin-gonic/gin"
)

// add middlewares here if needed, engine no need

func ProvideMiddlewares(cfg *config.Config, mc *metrics.MetricsCollector) []gin.HandlerFunc {
	middlewares := []gin.HandlerFunc{
		gin.Logger(),
		gin.Recovery(),
	}

	// 如果启用了指标，添加指标中间件
	if cfg.Metrics.Enabled {
		middlewares = append(middlewares, mc.GinMiddleware())
	}

	return middlewares
}

func ProvideGinEngine(cfg *config.Config,
	middlewares []gin.HandlerFunc,
	routers []router.Router,
	mc *metrics.MetricsCollector,
) *gin.Engine {
	engine := router.InitWeb(middlewares, routers)

	// 如果启用了指标，添加指标端点
	if cfg.Metrics.Enabled {
		engine.GET(cfg.Metrics.Path, mc.Handler())
	}

	return engine
}
