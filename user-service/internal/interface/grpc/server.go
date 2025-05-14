package grpc

import (
	"fmt"
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/metrics"
	authpb "github.com/Wafer233/msproject-be/user-service/proto/auth"
	captchapb "github.com/Wafer233/msproject-be/user-service/proto/captcha"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type GrpcServer struct {
	server      *grpc.Server
	address     string
	authSvc     service.AuthService
	captchaSvc  service.CaptchaService
	metricsConf config.MetricsConfig
	metrics     *metrics.GORMMetrics
}

func NewGrpcServer(
	address string,
	authSvc service.AuthService,
	captchaSvc service.CaptchaService,
	metricsConf config.MetricsConfig,
	metrics *metrics.GORMMetrics,
) *GrpcServer {
	return &GrpcServer{
		address:     address,
		authSvc:     authSvc,
		captchaSvc:  captchaSvc,
		metricsConf: metricsConf,
		metrics:     metrics,
	}
}

func (gs *GrpcServer) Start() error {
	// Start gRPC server
	lis, err := net.Listen("tcp", gs.address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Create gRPC server
	gs.server = grpc.NewServer()

	// Register services
	authpb.RegisterAuthServiceServer(gs.server, NewAuthServiceServer(gs.authSvc))
	captchapb.RegisterCaptchaServiceServer(gs.server, NewCaptchaServiceServer(gs.captchaSvc))

	// Start gRPC server in a goroutine
	go func() {
		zap.L().Info("Starting gRPC server", zap.String("address", gs.address))
		if err := gs.server.Serve(lis); err != nil {
			zap.L().Fatal("Failed to start gRPC server", zap.Error(err))
		}
	}()

	// Start HTTP metrics server if enabled
	if gs.metricsConf.Enabled {
		r := gin.Default()
		r.GET(gs.metricsConf.Endpoint, gs.metrics.Handler())

		metricsAddr := ":8090" // Choose an available port
		zap.L().Info("Starting metrics HTTP server", zap.String("address", metricsAddr))
		return http.ListenAndServe(metricsAddr, r)
	}

	return nil
}

func (gs *GrpcServer) Stop() {
	if gs.server != nil {
		gs.server.GracefulStop()
	}
}
