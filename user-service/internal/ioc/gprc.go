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
	gormMetrics *metrics.GORMMetrics,
) *grpc.GrpcServer {
	return grpc.NewGrpcServer(
		cfg.GRPC.Addr,
		authService,
		captchaService,
		cfg.Metrics,
		gormMetrics,
	)
}
