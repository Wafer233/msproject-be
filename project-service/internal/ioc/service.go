package ioc

import (
	"github.com/Wafer233/msproject-be/project-service/internal/application/service"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
)

func ProvideDefaultMenuService(mr repository.MenuRepository) service.MenuService {
	return service.NewDefaultMenuService(mr)
}
