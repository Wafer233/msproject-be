package ioc

import (
	"github.com/Wafer233/msproject-be/project-service/config"
	"github.com/Wafer233/msproject-be/project-service/internal/application/service"
	"github.com/Wafer233/msproject-be/project-service/internal/interface/grpc"
)

func ProvideGrpcServer(
	cfg *config.Config,
	menuService service.MenuService,
	projectService service.ProjectService,
) *grpc.GrpcServer {
	return grpc.NewGrpcServer(
		cfg.GRPC.Addr,
		menuService,
		projectService,
	)
}
