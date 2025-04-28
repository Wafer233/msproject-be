package handler

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/api-gateway/internal/domain/model"
	"github.com/Wafer233/msproject-be/api-gateway/internal/infrastructure/metrics"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// RegisterHandler 处理注册相关请求
type RegisterHandler struct {
	as *service.AuthService
	mc *metrics.MetricsCollector
}

// NewRegisterHandler 创建注册处理器
func NewRegisterHandler(as *service.AuthService, mc *metrics.MetricsCollector) *RegisterHandler {
	return &RegisterHandler{
		as: as,
		mc: mc,
	}
}

func (h *RegisterHandler) Register(c *gin.Context) {
	startTime := time.Now()
	result := &common.Result{}

	// 解析请求
	var req dto.RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		h.mc.RecordRegisterAttempt(false)
		h.mc.RecordErrorResponse("/project/login/register", http.StatusBadRequest)
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}

	// 验证参数
	if err := req.Verify(); err != nil {
		h.mc.RecordRegisterAttempt(false)
		h.mc.RecordErrorResponse("/project/login/register", int(model.RegisterInvalidData))
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}

	// 调用服务层
	err := h.as.Register(c, req)
	if err != nil {
		h.mc.RecordRegisterAttempt(false)
		h.mc.RecordErrorResponse("/project/login/register", int(model.RegisterFailed))
		c.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, err.Error()))
		return
	}

	h.mc.RecordRegisterAttempt(true)
	h.mc.RecordSuccessResponse("/project/login/register")
	h.mc.ObserveAuthResponseTime("register", time.Since(startTime).Seconds())

	// 返回结果
	c.JSON(http.StatusOK, result.Success(nil))
}
