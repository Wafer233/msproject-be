package grpc

import (
	"fmt"
	"github.com/Wafer233/msproject-be/project-service/internal/application/service"
	menupb "github.com/Wafer233/msproject-be/project-service/proto/menu"
	projpb "github.com/Wafer233/msproject-be/project-service/proto/project"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

// GrpcServer 处理gRPC服务器的设置和生命周期
type GrpcServer struct {
	server  *grpc.Server
	address string

	// add service
	menuService    service.MenuService
	projectService service.ProjectService
}

// NewGrpcServer 创建一个新的gRPC服务器
func NewGrpcServer(
	address string,
	menuService service.MenuService,
	projectService service.ProjectService,
) *GrpcServer {
	return &GrpcServer{
		address:        address,
		menuService:    menuService,
		projectService: projectService,
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
	menupb.RegisterMenuServiceServer(gs.server, NewMenuServiceServer(gs.menuService))
	projpb.RegisterProjectServiceServer(gs.server, NewProjectServiceServer(gs.projectService))

	// Start server
	zap.L().Info("Starting gRPC server", zap.String("address", gs.address))
	return gs.server.Serve(lis)
}

func (gs *GrpcServer) Stop() {
	if gs.server != nil {
		gs.server.GracefulStop()
	}
}
