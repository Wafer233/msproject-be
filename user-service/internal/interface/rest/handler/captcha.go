package handler

import (
	"github.com/Wafer233/msproject-be/common"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CaptchaHandler struct {
	cs service.CaptchaService
}

func NewCaptchaHandler(cs service.CaptchaService) *CaptchaHandler {
	return &CaptchaHandler{cs: cs}
}

func (hl *CaptchaHandler) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")

	if !common.VerifyMobile(mobile) {
		ctx.JSON(http.StatusOK, result.Fail(model.LoginMobileNotLegal, "手机号不合法"))
		return
	}

	// 生成验证码并保存
	code, err := hl.cs.GenerateCaptcha(ctx, mobile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, result.Fail(model.LoginSendCodeFail, "发送失败"))
		return
	}

	ctx.JSON(http.StatusOK, result.Success(code))
}
