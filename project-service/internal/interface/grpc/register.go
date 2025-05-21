package grpc

import (
	"github.com/Wafer233/msproject-be/project-service/internal/interface/grpc/handler"
	projPb "github.com/Wafer233/msproject-be/project-service/proto/project"
	"google.golang.org/grpc"
)

type ServiceRegister struct {
	projectGRPCHandler *handler.ProjectGRPCHandler
}

func NewServiceRegister(
	projectGRPCHandler *handler.ProjectGRPCHandler,
) *ServiceRegister {
	return &ServiceRegister{
		projectGRPCHandler: projectGRPCHandler,
	}
}

func (service *ServiceRegister) RegisterAll(server *grpc.Server) {
	projPb.RegisterProjectServiceServer(server, service.projectGRPCHandler)

}
