package middleware

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/domain/model"
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/grpc"
	authpb "github.com/Wafer233/msproject-be/api-gateway/proto/auth"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenVerifyMiddleware(clientMgr *grpc.GrpcClientManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := &common.Result{}

		// Get token from header
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, result.Fail(model.Unauthorized, "Unauthorized"))
			c.Abort()
			return
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
