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
		//services
		ProvideDefaultMenuService,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
