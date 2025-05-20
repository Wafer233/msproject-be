package config

import (
	"github.com/Wafer233/msproject-be/common/logs"
	"github.com/spf13/viper"
)

type Config struct {
	Server  ServerConfig
	GRPC    GRPCConfig
	MySQL   MySQLConfig
	Redis   RedisConfig
	Zap     logs.LogConfig
	viper   *viper.Viper
	JWT     JWTConfig
	Metrics MetricsConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type GRPCConfig struct {
	Name    string
	Addr    string
	Version string
	Weight  int64
}

type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type JWTConfig struct {
	AccessExp     int64
	RefreshExp    int64
	AccessSecret  string
	RefreshSecret string
}

type MetricsConfig struct {
	Enabled   bool   // Whether metrics are enabled
	Endpoint  string // HTTP endpoint for metrics
	Namespace string // Metrics namespace
}
