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
		ProvideProjectHandler,
		ProvideOrganizationHandler, //add for _getOrgList
		//metrics
		ProvideMetricsCollector,
		//hanlder middlerware
		ProvideAuthMiddleware,
		//router
		ProvideAuthRouter,
		ProvideMenuRouter,
		ProvideProjectRouter,
		ProvideOrganizationRouter, // add for _getOrgList
		ProvideRouters,

		//service
		ProvideAuthService,
		ProvideCaptchaService,
		ProvideMenuService,
		ProvideProjectService,
		ProvideOrganizationService, // add for _getOrgList

		wire.Struct(new(App), "*"),
	)

	return new(App), nil
}
