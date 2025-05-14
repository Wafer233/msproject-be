package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Handler returns a Gin handler for Prometheus metrics
func (gm *GORMMetrics) Handler() gin.HandlerFunc {
	return gin.WrapH(promhttp.HandlerFor(gm.registry, promhttp.HandlerOpts{}))
}
