// api-gateway/internal/infrastructure/metrics/metrics.go

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

	// Business-specific metrics
	loginAttempts    *prometheus.CounterVec
	registerAttempts *prometheus.CounterVec
	captchaRequests  *prometheus.CounterVec
	menuRequests     *prometheus.CounterVec
	errorResponses   *prometheus.CounterVec
	successResponses *prometheus.CounterVec
	authResponseTime *prometheus.HistogramVec
	menuResponseTime *prometheus.HistogramVec
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

		// Business specific metrics
		loginAttempts: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "login_attempts_total",
				Help:      "登录尝试总数",
			},
			[]string{"status"}, // success or failure
		),
		registerAttempts: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "register_attempts_total",
				Help:      "注册尝试总数",
			},
			[]string{"status"}, // success or failure
		),
		captchaRequests: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "captcha_requests_total",
				Help:      "验证码请求总数",
			},
			[]string{"status"}, // success or failure
		),
		menuRequests: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "menu_requests_total",
				Help:      "菜单请求总数",
			},
			[]string{"status"}, // success or failure
		),
		errorResponses: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "error_responses_total",
				Help:      "错误响应总数",
			},
			[]string{"path", "error_code"},
		),
		successResponses: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "success_responses_total",
				Help:      "成功响应总数",
			},
			[]string{"path"},
		),
		authResponseTime: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "auth_response_time_seconds",
				Help:      "认证响应时间（秒）",
				Buckets:   defaultDurationBuckets,
			},
			[]string{"operation"}, // login, register, captcha
		),
		menuResponseTime: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "menu_response_time_seconds",
				Help:      "菜单响应时间（秒）",
				Buckets:   defaultDurationBuckets,
			},
			[]string{"operation"}, // index
		),
	}

	// 注册所有指标
	registry.MustRegister(
		mc.requestsTotal,
		mc.requestDuration,
		mc.requestsInProgress,
		mc.requestSize,
		mc.responseSize,
		mc.loginAttempts,
		mc.registerAttempts,
		mc.captchaRequests,
		mc.menuRequests,
		mc.errorResponses,
		mc.successResponses,
		mc.authResponseTime,
		mc.menuResponseTime,
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

// RecordLoginAttempt 记录登录尝试
func (mc *MetricsCollector) RecordLoginAttempt(success bool) {
	status := "success"
	if !success {
		status = "failure"
	}
	mc.loginAttempts.WithLabelValues(status).Inc()
}

// RecordRegisterAttempt 记录注册尝试
func (mc *MetricsCollector) RecordRegisterAttempt(success bool) {
	status := "success"
	if !success {
		status = "failure"
	}
	mc.registerAttempts.WithLabelValues(status).Inc()
}

// RecordCaptchaRequest 记录验证码请求
func (mc *MetricsCollector) RecordCaptchaRequest(success bool) {
	status := "success"
	if !success {
		status = "failure"
	}
	mc.captchaRequests.WithLabelValues(status).Inc()
}

// RecordMenuRequest 记录菜单请求
func (mc *MetricsCollector) RecordMenuRequest(success bool) {
	status := "success"
	if !success {
		status = "failure"
	}
	mc.menuRequests.WithLabelValues(status).Inc()
}

// RecordErrorResponse 记录错误响应
func (mc *MetricsCollector) RecordErrorResponse(path string, errorCode int) {
	mc.errorResponses.WithLabelValues(path, string(errorCode)).Inc()
}

// RecordSuccessResponse 记录成功响应
func (mc *MetricsCollector) RecordSuccessResponse(path string) {
	mc.successResponses.WithLabelValues(path).Inc()
}

// ObserveAuthResponseTime 观察认证响应时间
func (mc *MetricsCollector) ObserveAuthResponseTime(operation string, duration float64) {
	mc.authResponseTime.WithLabelValues(operation).Observe(duration)
}

// ObserveMenuResponseTime 观察菜单响应时间
func (mc *MetricsCollector) ObserveMenuResponseTime(operation string, duration float64) {
	mc.menuResponseTime.WithLabelValues(operation).Observe(duration)
}
