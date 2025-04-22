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
		ProvideRedisCaptchaRepository,
		ProvideGORMMemberRepository,
		ProvideGORMOrganizationRepository,
		//routers
		ProvideAuthRouter,
		//services
		ProvideDefaultCaptchaService,
		ProvideDefaultAuthService,
		ProvidePasswordService,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
