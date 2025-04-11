//go:build wireinject

package ioc

import (
	"github.com/google/wire"
)

var captchaTest = wire.NewSet( // 这里是第一次构建项目的验证码
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
