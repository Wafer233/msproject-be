package handler

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GetCaptchaHandler struct {
	service *service.GatewayGetCaptchaService
}

func NewGetCaptchaHandler(service *service.GatewayGetCaptchaService) *GetCaptchaHandler {
	return &GetCaptchaHandler{
		service: service,
	}
}

func (handler *GetCaptchaHandler) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")

	c, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	captcha, err := handler.service.GetCaptcha(c, mobile)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginGetCaptchaFail, "生成验证码失败"))
		return
	}

	ctx.JSON(http.StatusOK, result.Success(captcha))

}
