package ioc

import "github.com/Wafer233/msproject-be/user-service/internal/interface/grpc"

type App struct {
	GrpcServer *grpc.UserServer
}
