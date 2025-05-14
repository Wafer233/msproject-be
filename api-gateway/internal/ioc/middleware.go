package ioc

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/grpc"
	"github.com/Wafer233/msproject-be/api-gateway/internal/interfaces/rest/middleware"
	"github.com/gin-gonic/gin"
)

func ProvideAuthMiddleware(clientMgr *grpc.GrpcClientManager) gin.HandlerFunc {
	return middleware.TokenVerifyMiddleware(clientMgr)
}
