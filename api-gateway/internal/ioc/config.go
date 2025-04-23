package ioc

import "github.com/Wafer233/msproject-be/api-gateway/config"

func ProvideViperConfig() *config.Config {
	return config.NewConfig()
}
