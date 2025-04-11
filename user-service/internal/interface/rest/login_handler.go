package rest

import (
	"github.com/Wafer233/msproject-be/common"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	captchaService *service.CaptchaService
}

func NewLoginHandler(captchaService *service.CaptchaService) *LoginHandler {
	return &LoginHandler{
		captchaService: captchaService,
	}
}

// GetCaptcha 获取手机验证码
func (h *LoginHandler) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}

	// 1. 获取参数
	mobile := ctx.PostForm("mobile")

	// 2. 验证手机合法性
	if !common.VerifyMobile(mobile) {
		ctx.JSON(200, result.Fail(model.LoginMobileNotLegal, "手机号不合法"))
		return
	}

	// 3. 生成并发送验证码
	code, err := h.captchaService.GenerateCaptcha(ctx, mobile)
	if err != nil {
		ctx.JSON(200, result.Fail(model.SystemError, "系统错误"))
		return
	}

	ctx.JSON(200, result.Success(code))
}
