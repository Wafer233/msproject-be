package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/config"
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/grpc"
)

func ProvideGrpcClientManager(cfg *config.Config) (*grpc.GrpcClientManager, error) {
	return grpc.NewGrpcClientManager(cfg)
}
