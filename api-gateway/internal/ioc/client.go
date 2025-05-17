package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/config"
	"github.com/Wafer233/msproject-be/api-gateway/internal/client"
)

// no need to change this file

func ProvideGrpcClientManager(cfg *config.Config) (*client.GrpcClientManager, error) {
	return client.NewGrpcClientManager(cfg)
}
