//go:build wireinject

package ioc

import (
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		//config
		ProvideViperConfig,
		//db
		ProvideDB,
		//grpc
		ProvideGrpcServer,
		//repositories
		ProvideGORMMenuRepository,
		ProvideGORMProjectRepository,
		//services
		ProvideDefaultMenuService,
		ProvideDefaultProjectService,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
