package grpc

import (
	"github.com/Wafer233/msproject-be/api-gateway/config"
	authpb "github.com/Wafer233/msproject-be/api-gateway/proto/auth"
	captchapb "github.com/Wafer233/msproject-be/api-gateway/proto/captcha"
	menupb "github.com/Wafer233/msproject-be/api-gateway/proto/menu"
	projpb "github.com/Wafer233/msproject-be/api-gateway/proto/project"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientManager struct {
	// ----------------  添加服务需要在这里新增连接 ----------------
	UserConn    *grpc.ClientConn
	ProjectConn *grpc.ClientConn
	// ----------------  添加服务需要在这里新增客户端 ----------------
	AuthClient    authpb.AuthServiceClient
	CaptchaClient captchapb.CaptchaServiceClient
	MenuClient    menupb.MenuServiceClient
	ProjectClient projpb.ProjectServiceClient
}

func NewGrpcClientManager(cfg *config.Config) (*GrpcClientManager, error) {
	userConn, err := grpc.Dial(
		// ----------------  添加服务需要在这里新增配置 ----------------
		cfg.UserService.GrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zap.L().Error("无法连接 user-service", zap.Error(err))
		return nil, err
	}

	// 连接 project-service（包含 MenuService）
	projectConn, err := grpc.Dial(
		// ----------------  添加服务需要在这里新增配置 ----------------
		cfg.ProjectService.GrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		zap.L().Error("无法连接 project-service", zap.Error(err))
		return nil, err
	}

	// ----------------  添加服务需要在这里新增客服端的实现 ----------------
	authClient := authpb.NewAuthServiceClient(userConn)
	captchaClient := captchapb.NewCaptchaServiceClient(userConn)
	menuClient := menupb.NewMenuServiceClient(projectConn)
	projectClient := projpb.NewProjectServiceClient(projectConn)

	return &GrpcClientManager{
		UserConn:      userConn,
		ProjectConn:   projectConn,
		AuthClient:    authClient,
		CaptchaClient: captchaClient,
		MenuClient:    menuClient,
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
