package middleware

import (
	"context"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//func TokokenVerifyMiddleware(clientMgr *grpc.GrpcClientManager) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		result := &common.Result{}
//
//		// Get token from header
//		token := ctx.GetHeader("Authorization")
//
//		// 去除可能的 "bearer " 前缀（不区分大小写）
//		token := authHeader
//		if len(authHeader) > 7 && strings.ToLower(authHeader[0:7]) == "bearer " {
//			token = authHeader[7:]
//		}
//
//		// Call token verify service
//		ctx := context.Background()
//		resp, err := clientMgr.AuthClient.TokenVerify(ctx, &authpb.TokenVerifyRequest{Token: token})
//
//		if err != nil {
//			ctx.JSON(http.StatusUnauthorized, result.Fail(1000, "Invalid token"))
//			ctx.Abort()
//			return
//		}
//
//		//3.处理结果，认证通过 将信息放入gin的上下文 失败返回未登录
//		ctx.Set("memberId", response.Member.Id)
//		ctx.Set("memberName", response.Member.Name)
//		ctx.Set("organizationCode", response.Member.OrganizationCode)
//		ctx.Next()
//	}
//}

func TokenVerifyMiddleware() func(*gin.Context) {
	return func(ctx *gin.Context) {
		result := &common.Result{}
		//1.从header中获取token
		token := ctx.GetHeader("Authorization")
		//2.调佣user服务进行token认证
		c, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancelFunc()

		var service *TokenVerifyService

		grpcResp, err := service.VerifyToken(c, token)
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
