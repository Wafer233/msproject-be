package ioc

import (
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	"github.com/Wafer233/msproject-be/user-service/internal/interface/grpc"
)

func ProvideGrpcServer(
	cfg *config.Config,
	authService service.AuthService,
	captchaService service.CaptchaService,
) *grpc.GrpcServer {
	return grpc.NewGrpcServer(
		cfg.GRPC.Addr,
		authService,
		captchaService,
	)
}
