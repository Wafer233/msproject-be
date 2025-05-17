package client

import (
	"github.com/Wafer233/msproject-be/api-gateway/config"
	loginPb "github.com/Wafer233/msproject-be/api-gateway/proto/login"
	projpb "github.com/Wafer233/msproject-be/api-gateway/proto/project"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientManager struct {
	UserConn    *grpc.ClientConn
	ProjectConn *grpc.ClientConn

	LoginClient   loginPb.LoginServiceClient
	ProjectClient projpb.ProjectServiceClient
}

func NewGrpcClientManager(cfg *config.Config) (*GrpcClientManager, error) {
	userConn, err := grpc.Dial(
		cfg.UserService.GrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zap.L().Warn("无法连接 user-service")
		return nil, err
	}

	projectConn, err := grpc.Dial(
		cfg.ProjectService.GrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zap.L().Warn("无法连接 project-service")
		return nil, err
	}

	loginClient := loginPb.NewLoginServiceClient(userConn)
	projectClient := projpb.NewProjectServiceClient(projectConn)

	return &GrpcClientManager{
		UserConn:    userConn,
		ProjectConn: projectConn,

		LoginClient:   loginClient,
		ProjectClient: projectClient,
	}, nil
}

func (m *GrpcClientManager) Close() error {
	var err error
	if m.UserConn != nil {
		err = m.UserConn.Close()
	}
	if m.ProjectConn != nil {
		err2 := m.ProjectConn.Close()
		if err == nil {
			err = err2
		}
	}
	return err
}
