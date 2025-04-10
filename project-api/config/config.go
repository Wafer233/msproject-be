package config

import (
	"github.com/Wafer233/msproject-be/project-common/logs"
	"github.com/spf13/viper"
	"log"
	"os"
)

var AppConf = InitConfig()

type Config struct {
	viper      *viper.Viper
	AppConfig  *AppConfig
	GrpcConfig *GrpcConfig
}

type AppConfig struct {
	Addr string
	Name string
}
type GrpcConfig struct {
	Services []map[string]any
}

func InitConfig() *Config {
	v := viper.New()
	conf := &Config{viper: v}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("app")
	conf.viper.SetConfigType("yml")
	conf.viper.AddConfigPath(workDir + "/config")

	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	conf.InitZapLog()
	conf.InitAppConfig()
	conf.InitGrpcConfig()
	return conf
}

func (c *Config) InitZapLog() {
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("maxSize"),
		MaxAge:        c.viper.GetInt("maxAge"),
		MaxBackups:    c.viper.GetInt("maxBackups"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}

}

func (c *Config) InitAppConfig() {
	gc := &AppConfig{
		Addr: c.viper.GetString("app.addr"),
		Name: c.viper.GetString("app.name"),
	}
	c.AppConfig = gc
}

func (c *Config) InitGrpcConfig() {
	gc := &GrpcConfig{}
	err := c.viper.UnmarshalKey("grpc", &gc.Services)
	if err != nil {
		log.Fatalln(err)
	}
	c.GrpcConfig = gc
}
