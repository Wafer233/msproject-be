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
		ProvideGetCaptchaHandler,
		ProvideGetOrgListHandler,
		ProvideIndexHandler,
		ProvideLoginHandler,
		ProvideRegisterHandler,
		ProvideProjectHandler,
		//metrics
		ProvideMetricsCollector,
		//middleware
		ProvideTokenVerifyMiddleware,
		//router
		ProvideIndexRouter,
		ProvideLoginRouter,
		ProvideOrganizationRouter,
		ProvideProjectRouter,
		ProvideRouters,
		//service
		ProvideGatewayGetCaptchaService,
		ProvideGatewayGetOrgListService,
		ProvideGatewayIndexService,
		ProvideGatewayLoginService,
		ProvideGatewayProjectService,
		ProvideGatewayRegisterService,

		wire.Struct(new(App), "*"),
	)

	return new(App), nil
}
