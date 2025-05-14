package middleware

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/domain/model"
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/grpc"
	authpb "github.com/Wafer233/msproject-be/api-gateway/proto/auth"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func TokenVerifyMiddleware(clientMgr *grpc.GrpcClientManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := &common.Result{}

		// Get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, result.Fail(model.Unauthorized, "Unauthorized"))
			c.Abort()
			return
		}

		// 去除可能的 "bearer " 前缀（不区分大小写）
		token := authHeader
		if len(authHeader) > 7 && strings.ToLower(authHeader[0:7]) == "bearer " {
			token = authHeader[7:]
		}

		// Call token verify service
		ctx := context.Background()
		resp, err := clientMgr.AuthClient.TokenVerify(ctx, &authpb.TokenVerifyRequest{Token: token})

		if err != nil {
			c.JSON(http.StatusUnauthorized, result.Fail(model.Unauthorized, "Invalid token"))
			c.Abort()
			return
		}

		// Set user info to context
		c.Set("userId", resp.Member.Id)
		c.Set("userName", resp.Member.Name)

		c.Next()
	}
}
