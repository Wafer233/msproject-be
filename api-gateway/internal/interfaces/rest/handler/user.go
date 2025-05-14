package handler

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserHandler struct {
	us *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

func (h *UserHandler) GetOrgList(c *gin.Context) {
	result := &common.Result{}

	// 从上下文中获取成员ID（由auth middleware设置）
	idAny, _ := c.Get("userId")
	id := idAny.(int64)

	// 创建带超时的上下文
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	// 调用服务
	orgs, err := h.us.GetOrgList(ctx, id)
	if err != nil {
		zap.L().Error("获取组织列表失败", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, err.Error()))
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, result.Success(orgs))
}
