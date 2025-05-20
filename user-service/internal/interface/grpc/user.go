package grpc

import (
	"fmt"
	"github.com/Wafer233/msproject-be/user-service/config"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/metrics"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type UserServer struct {
	address  string
	server   *grpc.Server
	register *ServiceRegister

	metricsConf config.MetricsConfig
	metrics     *metrics.GORMMetrics
}

func NewUserServer(
	address string,
	register *ServiceRegister,
	metricsConf config.MetricsConfig,
	metrics *metrics.GORMMetrics,
) *UserServer {
	return &UserServer{
		address:     address,
		register:    register,
		metricsConf: metricsConf,
		metrics:     metrics,
	}
}

func (s *UserServer) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s.server = grpc.NewServer()
	s.register.RegisterAll(s.server)

	go func() {
		zap.L().Info("Starting gRPC server", zap.String("address", s.address))
		if err := s.server.Serve(listener); err != nil {
			zap.L().Fatal("gRPC server stopped unexpectedly", zap.Error(err))
		}
	}()

	if s.metricsConf.Enabled {
		engine := gin.Default()
		engine.GET(s.metricsConf.Endpoint, s.metrics.Handler())
		return http.ListenAndServe(":8090", engine)
	}

	return nil
}
