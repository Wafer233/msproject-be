package config

import (
	"github.com/Wafer233/msproject-be/common/logs"
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Zap    logs.LogConfig
	viper  *viper.Viper
	// --- add here ---
	UserService    UserServiceConfig
	ProjectService ProjectServiceConfig
	// --- metrics ---
	Metrics MetricsConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type UserServiceConfig struct {
	GrpcAddr string
}

type ProjectServiceConfig struct {
	GrpcAddr string
}

// MetricsConfig Prometheus 指标配置
type MetricsConfig struct {
	Enabled   bool   // 是否启用指标收集
	Path      string // 指标访问的HTTP路径
	Namespace string // 指标命名空间前缀
	Subsystem string // 指标子系统名称
}
