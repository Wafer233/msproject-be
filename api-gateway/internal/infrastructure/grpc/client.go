package grpc

import (
	"github.com/Wafer233/msproject-be/api-gateway/config"
	authpb "github.com/Wafer233/msproject-be/api-gateway/proto/auth"
	captchapb "github.com/Wafer233/msproject-be/api-gateway/proto/captcha"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientManager struct {
	AuthClient    authpb.AuthServiceClient
	CaptchaClient captchapb.CaptchaServiceClient
	conn          *grpc.ClientConn
}

func NewGrpcClientManager(cfg *config.Config) (*GrpcClientManager, error) {
	// Create connection to user service
	conn, err := grpc.Dial(
		cfg.UserService.GrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	// Create clients
	authClient := authpb.NewAuthServiceClient(conn)
	captchaClient := captchapb.NewCaptchaServiceClient(conn)

	return &GrpcClientManager{
		AuthClient:    authClient,
		CaptchaClient: captchaClient,
		conn:          conn,
	}, nil
}

func (m *GrpcClientManager) Close() error {
	if m.conn != nil {
		return m.conn.Close()
	}
	return nil
}
