package handler

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginHandler 处理登录相关请求
type LoginHandler struct {
	as service.AuthService
}

func NewLoginHandler(as service.AuthService) *LoginHandler {
	return &LoginHandler{
		as: as,
	}
}

func (lh *LoginHandler) Login(c *gin.Context) {
	result := &common.Result{}

	// 解析请求
	var req dto.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数传递有误"))
		return
	}

	// 调用服务层
	resp, err := lh.as.Login(c, req)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(http.StatusInternalServerError, err.Error()))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, result.Success(resp))
}
