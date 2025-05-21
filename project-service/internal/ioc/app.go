package ioc

import "github.com/Wafer233/msproject-be/project-service/internal/interface/grpc"

type App struct {
	GrpcServer *grpc.ProjectServer
}
