package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/config"
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/grpc"
)

// no need to change this file

func ProvideGrpcClientManager(cfg *config.Config) (*grpc.GrpcClientManager, error) {
	return grpc.NewGrpcClientManager(cfg)
}
