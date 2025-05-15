package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/grpc"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/middleware"
	"github.com/gin-gonic/gin"
)

func ProvideTokenVerifyMiddleware(clientMgr *grpc.GrpcClientManager) gin.HandlerFunc {
	return middleware.NewTokenVerifyMiddleware(clientMgr.LoginClient)
}
