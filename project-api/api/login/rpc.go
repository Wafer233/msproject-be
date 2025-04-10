package login

import (
	loginService "github.com/Wafer233/msproject-be/project-user/pkg/service/login.service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var UserClient loginService.LoginServiceClient

func InitUserRpc() {
	conn, err := grpc.Dial(":8881", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	UserClient = loginService.NewLoginServiceClient(conn)
}
