package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/client"
	"github.com/Wafer233/msproject-be/api-gateway/internal/middleware"
)

func ProvideTokenVerifyMiddleware(clientMgr *client.GrpcClientManager) *middleware.TokenVerifyMiddleware {
	return middleware.NewTokenVerifyMiddleware(clientMgr.LoginClient)
}
