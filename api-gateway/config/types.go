package config

import (
	"github.com/Wafer233/msproject-be/common/logs"
	"github.com/spf13/viper"
)

type Config struct {
	Server      ServerConfig
	UserService UserServiceConfig
	Zap         logs.LogConfig
	viper       *viper.Viper
}

type ServerConfig struct {
	Name string
	Addr string
}

type UserServiceConfig struct {
	GrpcAddr string
}
