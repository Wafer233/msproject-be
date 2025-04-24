package config

import (
	"github.com/Wafer233/msproject-be/common/logs"
	"github.com/spf13/viper"
	"log"
	"os"
)

// NewConfig 创建并加载配置
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

	cfg.loadGRPC()
	cfg.loadMySQL()
	cfg.loadZap()

	return cfg
}

func (cfg *Config) loadGRPC() {
	cfg.GRPC = GRPCConfig{
		Name: cfg.viper.GetString("grpc.name"),
		Addr: cfg.viper.GetString("grpc.addr"),
	}
}

func (cfg *Config) loadMySQL() {
	cfg.MySQL = MySQLConfig{
		Host:     cfg.viper.GetString("mysql.host"),
		Port:     cfg.viper.GetString("mysql.port"),
		User:     cfg.viper.GetString("mysql.user"),
		Password: cfg.viper.GetString("mysql.password"),
		DBName:   cfg.viper.GetString("mysql.dbname"),
	}
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
