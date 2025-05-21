package ioc

import (
	"github.com/Wafer233/msproject-be/project-service/internal/application/service"
	"github.com/Wafer233/msproject-be/project-service/internal/interface/grpc/handler"
)

func ProvideProjectGRPCHandler(
	indexSvc service.IndexService,
	projectSvc service.ProjectService,
) *handler.ProjectGRPCHandler {
	return handler.NewProjectGRPCHandler(indexSvc, projectSvc)
}
