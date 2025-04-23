package handler

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterHandler 处理注册相关请求
type RegisterHandler struct {
	as *service.AuthService
}

// NewRegisterHandler 创建注册处理器
func NewRegisterHandler(as *service.AuthService) *RegisterHandler {
	return &RegisterHandler{
		as: as,
	}
}

func (h *RegisterHandler) Register(c *gin.Context) {
	result := &common.Result{}

	// 解析请求
	var req dto.RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}

	// 验证参数
	if err := req.Verify(); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
	}

	// 调用服务层
	err := h.as.Register(c, req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, err.Error()))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, result.Success(nil))
}
