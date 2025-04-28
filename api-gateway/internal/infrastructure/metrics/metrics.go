package metrics

import (
	"github.com/Wafer233/msproject-be/api-gateway/config"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 默认的直方图桶配置
var (
	defaultDurationBuckets = []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}
	defaultSizeBuckets     = []float64{100, 1000, 10000, 100000, 1000000}
)

// MetricsCollector 包含 Prometheus 指标收集器
type MetricsCollector struct {
	config             *config.MetricsConfig
	registry           *prometheus.Registry
	requestsTotal      *prometheus.CounterVec
	requestDuration    *prometheus.HistogramVec
	requestsInProgress *prometheus.GaugeVec
	requestSize        *prometheus.HistogramVec
	responseSize       *prometheus.HistogramVec
}

// NewMetricsCollector 创建新的指标收集器
func NewMetricsCollector(cfg *config.MetricsConfig) *MetricsCollector {
	registry := prometheus.NewRegistry()

	// 注册标准收集器
	registry.MustRegister(prometheus.NewGoCollector())
	registry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))

	// 构建指标名称
	namespace := cfg.Namespace
	subsystem := cfg.Subsystem

	// 创建收集器实例
	mc := &MetricsCollector{
		config:   cfg,
		registry: registry,
		requestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "requests_total",
				Help:      "HTTP请求总数",
			},
			[]string{"method", "path", "status"},
		),
		requestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "request_duration_seconds",
				Help:      "HTTP请求持续时间（秒）",
				Buckets:   defaultDurationBuckets,
			},
			[]string{"method", "path", "status"},
		),
		requestsInProgress: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "requests_in_progress",
				Help:      "当前正在处理的HTTP请求数",
			},
			[]string{"method", "path"},
		),
		requestSize: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "request_size_bytes",
				Help:      "HTTP请求大小（字节）",
				Buckets:   defaultSizeBuckets,
			},
			[]string{"method", "path"},
		),
		responseSize: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "response_size_bytes",
				Help:      "HTTP响应大小（字节）",
				Buckets:   defaultSizeBuckets,
			},
			[]string{"method", "path", "status"},
		),
	}

	// 注册所有指标
	registry.MustRegister(
		mc.requestsTotal,
		mc.requestDuration,
		mc.requestsInProgress,
		mc.requestSize,
		mc.responseSize,
	)

	return mc
}

// Handler 返回指标端点的HTTP处理器
func (mc *MetricsCollector) Handler() gin.HandlerFunc {
	return gin.WrapH(promhttp.HandlerFor(mc.registry, promhttp.HandlerOpts{}))
}

// GetRegistry 返回指标注册表
func (mc *MetricsCollector) GetRegistry() *prometheus.Registry {
	return mc.registry
}
