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
		//router
		ProvideAuthRouter,
		//service
		ProvideAuthService,
		ProvideCaptchaService,

		wire.Struct(new(App), "*"),
	)

	return new(App), nil
}
