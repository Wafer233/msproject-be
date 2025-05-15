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

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(projectService *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

func (h *ProjectHandler) GetMyProjects(c *gin.Context) {
	result := &common.Result{}

	// 从上下文获取用户ID - 注意键名要与中间件设置一致
	userId, exists := c.Get("userId") // 确保与中间件中设置的键名一致
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	// 解析分页参数
	var req dto.ProjectRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "Invalid parameters"))
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 创建超时上下文
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	// 调用服务
	response, err := h.projectService.GetMyProjects(ctx, userId.(int64), req.Page, req.PageSize)
	if err != nil {
		zap.L().Error("Failed to get my projects", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, "Server error"))
		return
	}

	// 返回结果 - 格式为 { "list": [...], "total": ... }
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  response.List,
		"total": response.Total,
	}))
}
