package handler

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/api-gateway/internal/domain/model"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type MenuHandler struct {
	ms *service.MenuService
}

func NewMenuHandler(ms *service.MenuService) *MenuHandler {
	return &MenuHandler{
		ms: ms,
	}
}

// Index 处理首页菜单请求
func (h *MenuHandler) Index(c *gin.Context) {
	result := &common.Result{}

	// 获取令牌
	token := c.GetHeader("Authorization")
	zap.S().Info("token:" + token)

	// 调用服务
	menuResponse, err := h.ms.GetMenus(c, token)
	if err != nil {
		zap.L().Error("获取菜单失败", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(model.SystemError, "获取菜单失败"))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, result.Success(menuResponse.Menus))
}
