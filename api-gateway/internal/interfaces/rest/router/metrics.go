package router

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/metrics"
	"github.com/gin-gonic/gin"
)

type MetricsRouter struct {
	metricsCollector *metrics.MetricsCollector
}

func NewMetricsRouter(mc *metrics.MetricsCollector) *MetricsRouter {
	return &MetricsRouter{metricsCollector: mc}
}

func (r *MetricsRouter) Register(engine *gin.Engine) {
	engine.GET("/metrics", r.metricsCollector.Handler())
}
