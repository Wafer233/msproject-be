package grpc

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type ProjectServer struct {
	address  string
	server   *grpc.Server
	register *ServiceRegister
}

func NewProjectServer(
	address string,
	register *ServiceRegister,

) *ProjectServer {
	return &ProjectServer{
		address:  address,
		register: register,
	}
}

func (s *ProjectServer) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", s.address, err)
	}

	s.server = grpc.NewServer()
	s.register.RegisterAll(s.server)

	zap.L().Info("gRPC server started", zap.String("address", s.address))

	// Serve 是阻塞的，正常运行时不会返回
	if err := s.server.Serve(listener); err != nil {
		zap.L().Fatal("gRPC server stopped unexpectedly", zap.Error(err))
		return err
	}

	return nil
}

func (s *ProjectServer) Stop() {
	if s.server != nil {
		zap.L().Info("Stopping gRPC server...")
		s.server.GracefulStop()
	}
}
