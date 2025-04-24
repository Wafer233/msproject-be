package config

import (
	"github.com/Wafer233/msproject-be/common/logs"
	"github.com/spf13/viper"
	"log"
	"os"
)

func NewConfig() *Config {
	v := viper.New()
	cfg := &Config{viper: v}

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalln("failed to get working dir:", err)
	}

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(workDir + "/config")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalln("failed to read config file:", err)
	}

	cfg.loadZap()
	cfg.loadUserService()
	cfg.loadProjectService()

	return cfg
}

func (cfg *Config) loadZap() {
	cfg.Zap = logs.LogConfig{
		DebugFileName: cfg.viper.GetString("zap.debugFileName"),
		InfoFileName:  cfg.viper.GetString("zap.infoFileName"),
		WarnFileName:  cfg.viper.GetString("zap.warnFileName"),
		MaxSize:       cfg.viper.GetInt("zap.maxSize"),
		MaxAge:        cfg.viper.GetInt("zap.maxAge"),
		MaxBackups:    cfg.viper.GetInt("zap.maxBackups"),
	}
	if err := logs.InitLogger(&cfg.Zap); err != nil {
		log.Fatalln("failed to init zap logger:", err)
	}
}

func (cfg *Config) loadUserService() {
	cfg.UserService = UserServiceConfig{
		GrpcAddr: cfg.viper.GetString("userService.grpcAddr"),
	}
}

func (cfg *Config) loadProjectService() {
	cfg.ProjectService = ProjectServiceConfig{
		GrpcAddr: cfg.viper.GetString("projectService.grpcAddr"),
	}
}
