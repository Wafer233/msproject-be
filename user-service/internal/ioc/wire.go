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
		//grpc
		ProvideGrpcServer,
		//redis
		ProvideRedisClient,
		//repositories
		ProvideRedisCaptchaRepository,
		ProvideGORMMemberRepository,
		ProvideGORMOrganizationRepository,
		//services
		ProvideDefaultCaptchaService,
		ProvideDefaultAuthService,
		ProvidePasswordService,
		ProvideJWTTokenService,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
