package grpc

import (
	"fmt"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	authpb "github.com/Wafer233/msproject-be/user-service/proto/auth"
	captchapb "github.com/Wafer233/msproject-be/user-service/proto/captcha"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	server     *grpc.Server
	address    string
	authSvc    service.AuthService
	captchaSvc service.CaptchaService
}

func NewGrpcServer(
	address string,
	authSvc service.AuthService,
	captchaSvc service.CaptchaService,
) *GrpcServer {
	return &GrpcServer{
		address:    address,
		authSvc:    authSvc,
		captchaSvc: captchaSvc,
	}
}

func (gs *GrpcServer) Start() error {
	lis, err := net.Listen("tcp", gs.address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// Create gRPC server
	gs.server = grpc.NewServer()

	// Register services
	authpb.RegisterAuthServiceServer(gs.server, NewAuthServiceServer(gs.authSvc))
	captchapb.RegisterCaptchaServiceServer(gs.server, NewCaptchaServiceServer(gs.captchaSvc))

	// Start server
	zap.L().Info("Starting gRPC server", zap.String("address", gs.address))
	return gs.server.Serve(lis)
}

func (gs *GrpcServer) Stop() {
	if gs.server != nil {
		gs.server.GracefulStop()
	}
}
