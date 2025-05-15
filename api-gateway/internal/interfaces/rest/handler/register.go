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

// RegisterHandler 处理注册相关请求s
type RegisterHandler struct {
	service *service.GatewayRegisterService
}

// NewRegisterHandler 创建注册处理器
func NewRegisterHandler(service *service.GatewayRegisterService) *RegisterHandler {
	return &RegisterHandler{
		service: service,
	}
}

func (handler *RegisterHandler) Register(ctx *gin.Context) {
	result := &common.Result{}

	c, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	var dtoReq dto.RegisterRequest
	if err := ctx.ShouldBind(&dtoReq); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginRegisterBindFail, "注册请求Bind失败"))
		return
	}

	// 验证参数
	if err := dtoReq.Verify(); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginRegisterReqFormatError, "注册请求参数格式错误"))
		return
	}

	// 调用服务层
	err := handler.service.Register(c, dtoReq)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginRegisterServiceError, "注册服务调用失败"))
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, result.Success(""))
}
