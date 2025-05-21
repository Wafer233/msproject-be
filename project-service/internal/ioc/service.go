package ioc

import (
	"github.com/Wafer233/msproject-be/project-service/internal/application/service"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
)

func ProvideDefaultIndexService(mr repository.MenuRepo) service.IndexService {
	return service.NewDefaultIndexService(mr)
}

func ProvideDefaultProjectService(pr repository.ProjectRepo) service.ProjectService {
	return service.NewDefaultProjectService(pr)
}
