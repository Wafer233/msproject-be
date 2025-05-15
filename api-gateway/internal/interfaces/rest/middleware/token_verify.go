package middleware

import (
	"context"
	pb "github.com/Wafer233/msproject-be/api-gateway/proto/login"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NewTokenVerifyMiddleware(client pb.LoginServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := &common.Result{}
		//1.从header中获取token
		token := ctx.GetHeader("Authorization")
		//2.调佣user服务进行token认证
		c, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancelFunc()

		grpcMsg := &pb.TokenVerifyMessage{
			Token: token,
		}

		grpcResp, err := client.TokenVerify(c, grpcMsg)
		if err != nil {
			ctx.JSON(http.StatusOK, result.Fail(common.LoginTokenVerifyMiddlewareServiceFail, "token服务错误"))
			ctx.Abort()
			return
		}
		//3.处理结果，认证通过 将信息放入gin的上下文 失败返回未登录
		ctx.Set("memberId", grpcResp.Id)
		ctx.Set("memberName", grpcResp.Name)
		ctx.Set("organizationCode", grpcResp.OrganizationCode)
		ctx.Next()
	}
}
