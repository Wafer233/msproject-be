package ioc

import "github.com/Wafer233/msproject-be/project-service/config"

func ProvideViperConfig() *config.Config {
	return config.NewConfig()
}
