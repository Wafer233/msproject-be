//go:build wireinject
// +build wireinject

package ioc

import (
	"github.com/google/wire"
)

func InitApp() (*App, error) {
	wire.Build(
		//config
		ProvideViperConfig,
		//engine
		ProvideMiddlewares,
		ProvideGinEngine,
		//client
		ProvideGrpcClientManager,
		//handler
		ProvideLoginHttpHandler,
		ProvideProjectHandler,
		//metrics
		ProvideMetricsCollector,
		//middleware
		ProvideTokenVerifyMiddleware,
		//router
		ProvideUserRouter,
		ProvideProjectRouter,
		ProvideRouters,

		wire.Struct(new(App), "*"),
	)

	return new(App), nil
}
