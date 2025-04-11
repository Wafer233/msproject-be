//go:build wireinject

package ioc

import (
	"github.com/google/wire"
)

var captchaTest = wire.NewSet( // 第三方依赖

	ProvideRedisCache,
	ProvideCachedCaptchaRepo,
	ProvideCachedCaptchaService,
	ProvideEngine,
	ProvideMiddlewares,
	ProvideCaptchaHandler,
	ProvideCaptchaRouter,
)

func InitApp() *App {
	wire.Build(
		ProvideViperConfig,
		ProvideRedisClient,

		captchaTest,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
