package handler

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
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

// Index 处理导航菜单请求
func (h *MenuHandler) Index(c *gin.Context) {
	result := &common.Result{}

	// 获取授权令牌
	token := c.GetHeader("Authorization")
	zap.L().Debug("处理导航菜单请求", zap.String("token", token))

	// 调用服务层获取菜单
	menuResponse, err := h.ms.GetMenus(c, token)
	if err != nil {
		zap.L().Error("获取菜单失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, result.Success(menuResponse.Menus))
}
