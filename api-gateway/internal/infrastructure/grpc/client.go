package grpc

import (
	"github.com/Wafer233/msproject-be/api-gateway/config"
	authpb "github.com/Wafer233/msproject-be/api-gateway/proto/auth"
	captchapb "github.com/Wafer233/msproject-be/api-gateway/proto/captcha"
	menupb "github.com/Wafer233/msproject-be/api-gateway/proto/menu"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientManager struct {
	conn *grpc.ClientConn
	// ----------------- add clients here -----------------
	AuthClient    authpb.AuthServiceClient
	CaptchaClient captchapb.CaptchaServiceClient
	MenuClient    menupb.MenuServiceClient
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
	// ----------------- add clients here -----------------
	authClient := authpb.NewAuthServiceClient(conn)
	captchaClient := captchapb.NewCaptchaServiceClient(conn)
	menuClient := menupb.NewMenuServiceClient(conn)

	return &GrpcClientManager{
		conn: conn,
		// ----------------- add clients here -----------------
		AuthClient:    authClient,
		CaptchaClient: captchaClient,
		MenuClient:    menuClient,
	}, nil
}

func (m *GrpcClientManager) Close() error {
	if m.conn != nil {
		return m.conn.Close()
	}
	return nil
}
