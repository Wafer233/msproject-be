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
