package handler

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/api-gateway/internal/domain/model"
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/metrics"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type CaptchaHandler struct {
	cs *service.CaptchaService
	mc *metrics.MetricsCollector
}

func NewCaptchaHandler(cs *service.CaptchaService, mc *metrics.MetricsCollector) *CaptchaHandler {
	return &CaptchaHandler{
		cs: cs,
		mc: mc,
	}
}

func (hl *CaptchaHandler) GetCaptcha(ctx *gin.Context) {
	startTime := time.Now()
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")

	if !common.VerifyMobile(mobile) {
		hl.mc.RecordCaptchaRequest(false)
		hl.mc.RecordErrorResponse("/project/login/getCaptcha", int(model.LoginMobileNotLegal))
		ctx.JSON(http.StatusOK, result.Fail(model.LoginMobileNotLegal, "手机号不合法"))
		go func() {
			zap.L().Warn("手机号不合法")
		}()
		return
	}

	// 生成验证码并保存
	code, err := hl.cs.GenerateCaptcha(ctx, mobile)
	if err != nil {
		hl.mc.RecordCaptchaRequest(false)
		hl.mc.RecordErrorResponse("/project/login/getCaptcha", int(model.LoginSendCodeFail))
		ctx.JSON(http.StatusInternalServerError, result.Fail(model.LoginSendCodeFail, "发送失败"))
		go func() {
			zap.L().Warn("发送失败")
		}()
		return
	}

	hl.mc.RecordCaptchaRequest(true)
	hl.mc.RecordSuccessResponse("/project/login/getCaptcha")
	hl.mc.ObserveAuthResponseTime("captcha", time.Since(startTime).Seconds())

	ctx.JSON(http.StatusOK, result.Success(code))
	go func() {
		zap.L().Info("验证码发送成功", zap.String("mobile", mobile), zap.String("code", code))
	}()
}
