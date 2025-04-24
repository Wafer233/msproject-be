package ioc

import "github.com/Wafer233/msproject-be/api-gateway/config"

// no need to add any new imports here

func ProvideViperConfig() *config.Config {
	return config.NewConfig()
}
