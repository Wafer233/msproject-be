//go:build wireinject
// +build wireinject

package ioc

import "github.com/google/wire"

func InitApp() (*App, error) {
	wire.Build(
		//config
		ProvideViperConfig,
		//engine
		ProvideMiddlewares,
		ProvideGinEngine,
		//grpc
		ProvideGrpcClientManager,
		//handler
		ProvideCaptchaHandler,
		ProvideRegisterHandler,
		ProvideLoginHandler,
		ProvideMenuHandler,
		//metrics
		ProvideMetricsCollector,
		//hanlder middlerware
		ProvideAuthMiddleware,
		//router
		ProvideAuthRouter,
		ProvideMenuRouter,
		ProvideRouters,
		//service
		ProvideAuthService,
		ProvideCaptchaService,
		ProvideMenuService,

		wire.Struct(new(App), "*"),
	)

	return new(App), nil
}
