package handler

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// LoginHandler 处理登录相关请求
type LoginHandler struct {
	service *service.GatewayLoginService
}

func NewLoginHandler(service *service.GatewayLoginService) *LoginHandler {
	return &LoginHandler{
		service: service,
	}
}

func (handler *LoginHandler) Login(ctx *gin.Context) {
	result := &common.Result{}

	// 解析请求
	var dtoReq *dto.LoginRequest
	if err := ctx.ShouldBind(&dtoReq); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginLoginReqBindFail, "登录请求参数绑定失败"))
		return
	}

	// 调用服务层
	c, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()
	dtoResp, err := handler.service.Login(c, dtoReq)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, err.Error()))
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, result.Success(dtoResp))
}
