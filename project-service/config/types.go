package config

import (
	"github.com/Wafer233/msproject-be/common/logs"
	"github.com/spf13/viper"
)

// Config 持有项目服务的配置
type Config struct {
	GRPC  GRPCConfig
	MySQL MySQLConfig
	Zap   logs.LogConfig
	viper *viper.Viper
}

// GRPCConfig 持有gRPC服务器配置
type GRPCConfig struct {
	Name string
	Addr string
}

// MySQLConfig 持有MySQL数据库配置
type MySQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}
