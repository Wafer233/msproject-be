//go:build wireinject

package ioc

import (
	"github.com/google/wire"
)

func InitApp() *App {
	wire.Build(
		//cache
		ProvideRedisCaptchaCache,
		ProvideRedisTokenCache,
		//config
		ProvideViperConfig,
		//db
		ProvideDB,
		//grpc
		ProvideUserServer,
		ProvideServiceRegister,
		//handler
		ProvideLoginGRPCHandler,
		//metrics
		ProvideGORMMetrics,
		//redis
		ProvideRedisClient,
		//repositories
		ProvideGORMMemberRepository,
		ProvideGORMOrganizationRepository,
		//services
		ProvideDefaultCaptchaService,
		ProvideDefaultTokenService,
		ProvideDefaultLoginService,
		ProvideDefaultOrganizationService,
		ProvideDefaultRegisterService,

		wire.Struct(new(App), "*"),
	)

	return new(App)
}
