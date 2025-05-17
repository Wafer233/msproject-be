package middleware

import (
	"context"
	pb "github.com/Wafer233/msproject-be/api-gateway/proto/login"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TokenVerifyMiddleware struct {
	client pb.LoginServiceClient
}

func NewTokenVerifyMiddleware(client pb.LoginServiceClient) *TokenVerifyMiddleware {
	return &TokenVerifyMiddleware{
		client: client,
	}
}

func (middleware *TokenVerifyMiddleware) TokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := &common.Result{}

		token := c.GetHeader("Authorization")

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		//ip := GetIp(c)

		grpcReq := &pb.TokenVerifyRequest{
			Token: token,
		}

		grpcResp, err := middleware.client.TokenVerify(ctx, grpcReq)

		if err != nil {
			c.JSON(http.StatusOK, result.Fail(
				common.MiddlewareTokenVerifyServiceFail,
				"调用token验证微服务服务失败"))
			c.Abort()
			return
		}

		c.Set("memberId", grpcResp.Member.Id)
		c.Set("memberName", grpcResp.Member.Name)
		c.Set("organizationCode", grpcResp.Member.OrganizationCode)
		c.Next()
	}
}
