package grpc

import (
	"github.com/Wafer233/msproject-be/user-service/internal/interface/grpc/handler"
	loginPb "github.com/Wafer233/msproject-be/user-service/proto/login"
	"google.golang.org/grpc"
)

type ServiceRegister struct {
	loginGRPCHandler *handler.LoginGRPCHandler
}

func NewServiceRegister(
	loginGRPCHandler *handler.LoginGRPCHandler,
) *ServiceRegister {
	return &ServiceRegister{
		loginGRPCHandler: loginGRPCHandler,
	}
}

func (service *ServiceRegister) RegisterAll(server *grpc.Server) {
	loginPb.RegisterLoginServiceServer(server, service.loginGRPCHandler)

}
