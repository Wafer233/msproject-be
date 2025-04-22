//go:build wireinject

package ioc

import (
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		//cache
		ProvideRedisCache,
		//config
		ProvideViperConfig,
		//db
		ProvideDB,
		//engine
		ProvideMiddlewares,
		ProvideEngine,
		//handlers
		ProvideCaptchaHandler,
		ProvideRegisterHandler,
		ProvideLoginHandler,
		//redis
		ProvideRedisClient,
		//repositories
		ProvideCachedCaptchaRepo,
		ProvideCachedMemberRepository,
		ProvideCachedOrganizationRepository,
		//routers
		ProvideAuthRouter,
		//services
		ProvideCachedCaptchaService,
		ProvideCachedAuthService,
		ProvidePasswordService,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
