package metrics

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// GinMiddleware 创建Gin中间件进行指标收集
func (mc *MetricsCollector) GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过指标端点以避免循环测量
		if c.Request.URL.Path == mc.config.Path {
			c.Next()
			return
		}

		// 启动计时器并记录正在进行的请求
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		requestSize := float64(c.Request.ContentLength)

		// 跟踪正在进行的请求
		mc.requestsInProgress.WithLabelValues(method, path).Inc()
		// 记录请求大小
		mc.requestSize.WithLabelValues(method, path).Observe(requestSize)

		// 处理请求
		c.Next()

		// 停止计时器并收集指标
		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())
		responseSize := float64(c.Writer.Size())

		// 减少正在进行的计数器
		mc.requestsInProgress.WithLabelValues(method, path).Dec()
		// 记录请求数
		mc.requestsTotal.WithLabelValues(method, path, status).Inc()
		// 记录请求持续时间
		mc.requestDuration.WithLabelValues(method, path, status).Observe(duration)
		// 记录响应大小
		mc.responseSize.WithLabelValues(method, path, status).Observe(responseSize)
	}
}
