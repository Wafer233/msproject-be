package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/metrics"
	"github.com/Wafer233/msproject-be/user-service/internal/interface/grpc"
)

func ProvideGrpcServer(
	cfg *config.Config,
	authService service.AuthService,
	captchaService service.CaptchaService,
	userService service.UserService, // add for _getOrgList
	gormMetrics *metrics.GORMMetrics,
) *grpc.GrpcServer {
	return grpc.NewGrpcServer(
		cfg.GRPC.Addr,
		authService,
		captchaService,
		userService, // add for _getOrgList
		cfg.Metrics,
		gormMetrics,
	)
}
