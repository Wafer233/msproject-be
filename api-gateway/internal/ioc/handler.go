package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/client"
	"github.com/Wafer233/msproject-be/api-gateway/internal/handler"
)

func ProvideLoginHttpHandler(clientMgr *client.GrpcClientManager) *handler.LoginHttpHandler {
	return handler.NewLoginHttpHandler(clientMgr.LoginClient)
}

func ProvideProjectHandler(clientMgr *client.GrpcClientManager) *handler.ProjectHttpHandler {
	return handler.NewProjectHttpHandler(clientMgr.ProjectClient)
}
