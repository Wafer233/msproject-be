//go:build wireinject

package ioc

import (
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(

		ProvideViperConfig,
		ProvideRedisClient,
		ProvideRedisCache,
		ProvideCachedCaptchaRepo,
		ProvideCachedCaptchaService,
		ProvideEngine,
		ProvideMiddlewares,
		ProvideCaptchaHandler,
		ProvideCaptchaRouter,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
