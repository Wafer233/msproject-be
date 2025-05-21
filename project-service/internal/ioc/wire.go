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
		ProvideServiceRegister,
		ProvideProjectServer,
		//handler
		ProvideProjectGRPCHandler,
		//repositories
		ProvideGORMMenuRepo,
		ProvideGORMProjectRepo,
		//services
		ProvideDefaultIndexService,
		ProvideDefaultProjectService,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
