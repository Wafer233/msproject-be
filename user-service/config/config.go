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

	cfg.loadServer()
	cfg.loadGRPC()
	cfg.loadMySQL()
	cfg.loadRedis()
	cfg.loadZap()

	return cfg
}

func (cfg *Config) loadServer() {
	cfg.Server = ServerConfig{
		Name: cfg.viper.GetString("server.name"),
		Addr: cfg.viper.GetString("server.addr"),
	}
}

func (cfg *Config) loadGRPC() {
	cfg.GRPC = GRPCConfig{
		Name:    cfg.viper.GetString("grpc.name"),
		Addr:    cfg.viper.GetString("grpc.addr"),
		Version: cfg.viper.GetString("grpc.version"),
		Weight:  cfg.viper.GetInt64("grpc.weight"),
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

func (cfg *Config) loadRedis() {
	cfg.Redis = RedisConfig{
		Host:     cfg.viper.GetString("redis.host"),
		Port:     cfg.viper.GetString("redis.port"),
		Password: cfg.viper.GetString("redis.password"),
		DB:       cfg.viper.GetInt("redis.db"),
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
