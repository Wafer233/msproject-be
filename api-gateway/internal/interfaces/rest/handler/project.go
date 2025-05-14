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

	// Get member ID from the context (set by auth middleware)
	memberId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusOK, result.Fail(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	// Parse pagination parameters
	var req dto.ProjectRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "Invalid parameters"))
		return
	}

	// Set default values if not provided
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	// Call service
	response, err := h.projectService.GetMyProjects(ctx, memberId.(int64), req.Page, req.PageSize)
	if err != nil {
		zap.L().Error("Failed to get my projects", zap.Error(err))
		c.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, "Server error"))
		return
	}

	// Return response
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  response.List,
		"total": response.Total,
	}))
}
