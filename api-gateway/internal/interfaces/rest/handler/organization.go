package handler

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type OrganizationHandler struct {
	organizationService service.OrganizationService
}

func NewOrganizationHandler(organizationService service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{
		organizationService: organizationService,
	}
}

func (h *OrganizationHandler) GetOrgList(c *gin.Context) {
	result := &common.Result{}

	// 从上下文获取用户ID（由认证中间件设置）
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusUnauthorized, "未授权"))
		return
	}

	// 创建带超时的上下文
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	// 调用服务
	orgs, err := h.organizationService.GetOrgList(ctx, userId.(int64))
	if err != nil {
		zap.L().Error("获取组织列表失败", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, "服务器错误"))
		return
	}

	// 检查是否为空列表，如果是则返回空数组而不是nil
	if orgs == nil {
		c.JSON(http.StatusOK, result.Success([]dto.OrganizationDTO{}))
		return
	}

	// 直接返回组织数组，不需要额外包装
	c.JSON(http.StatusOK, result.Success(orgs))
}
