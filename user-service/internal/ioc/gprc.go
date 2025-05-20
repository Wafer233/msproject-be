package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/metrics"
	"github.com/Wafer233/msproject-be/user-service/internal/interface/grpc"
	"github.com/Wafer233/msproject-be/user-service/internal/interface/grpc/handler"
)

func ProvideUserServer(
	config *config.Config,
	register *grpc.ServiceRegister,
	metrics *metrics.GORMMetrics,
) *grpc.UserServer {
	address := config.GRPC.Addr
	metricsConf := config.Metrics
	return grpc.NewUserServer(address, register, metricsConf, metrics)
}

func ProvideServiceRegister(
	loginGRPCHandler *handler.LoginGRPCHandler,
) *grpc.ServiceRegister {
	return grpc.NewServiceRegister(
		loginGRPCHandler,
	)
}
