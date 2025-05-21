package ioc

import (
	"github.com/Wafer233/msproject-be/project-service/config"
	"github.com/Wafer233/msproject-be/project-service/internal/interface/grpc"
	"github.com/Wafer233/msproject-be/project-service/internal/interface/grpc/handler"
)

func ProvideServiceRegister(projectGRPCHandler *handler.ProjectGRPCHandler) *grpc.ServiceRegister {
	return grpc.NewServiceRegister(projectGRPCHandler)
}

func ProvideProjectServer(
	config *config.Config,
	register *grpc.ServiceRegister,
) *grpc.ProjectServer {
	addr := config.GRPC.Addr
	return grpc.NewProjectServer(addr, register)
}
