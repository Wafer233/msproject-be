// api-gateway/internal/interfaces/rest/handler/login.go

package handler

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/metrics"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// LoginHandler 处理登录相关请求
type LoginHandler struct {
	as *service.AuthService
	mc *metrics.MetricsCollector
}

func NewLoginHandler(as *service.AuthService, mc *metrics.MetricsCollector) *LoginHandler {
	return &LoginHandler{
		as: as,
		mc: mc,
	}
}

func (lh *LoginHandler) Login(c *gin.Context) {
	startTime := time.Now()
	result := &common.Result{}

	// 解析请求
	var req dto.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		lh.mc.RecordLoginAttempt(false)
		lh.mc.RecordErrorResponse("/project/login", http.StatusBadRequest)
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}

	// 调用服务层
	resp, err := lh.as.Login(c, req)
	if err != nil {
		lh.mc.RecordLoginAttempt(false)
		lh.mc.RecordErrorResponse("/project/login", http.StatusInternalServerError)
		c.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, err.Error()))
		return
	}

	lh.mc.RecordLoginAttempt(true)
	lh.mc.RecordSuccessResponse("/project/login")
	lh.mc.ObserveAuthResponseTime("login", time.Since(startTime).Seconds())

	// 返回结果
	c.JSON(http.StatusOK, result.Success(resp))
}
