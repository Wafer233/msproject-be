package ioc

import "github.com/Wafer233/msproject-be/user-service/config"

func ProvideViperConfig() *config.Config {
	return config.NewConfig()
}
