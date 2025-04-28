// api-gateway/internal/interfaces/rest/handler/menu.go

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

type MenuHandler struct {
	ms *service.MenuService
	mc *metrics.MetricsCollector
}

func NewMenuHandler(ms *service.MenuService, mc *metrics.MetricsCollector) *MenuHandler {
	return &MenuHandler{
		ms: ms,
		mc: mc,
	}
}

// Index 处理首页菜单请求
func (h *MenuHandler) Index(c *gin.Context) {
	startTime := time.Now()
	result := &common.Result{}

	// 获取令牌
	token := c.GetHeader("Authorization")
	zap.S().Info("token:" + token)

	// 调用服务
	menuResponse, err := h.ms.GetMenus(c, token)
	if err != nil {
		h.mc.RecordMenuRequest(false)
		h.mc.RecordErrorResponse("/project/index", int(model.SystemError))
		zap.L().Error("获取菜单失败", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(model.SystemError, "获取菜单失败"))
		return
	}

	h.mc.RecordMenuRequest(true)
	h.mc.RecordSuccessResponse("/project/index")
	h.mc.ObserveMenuResponseTime("index", time.Since(startTime).Seconds())

	// 返回结果
	c.JSON(http.StatusOK, result.Success(menuResponse.Menus))
}
