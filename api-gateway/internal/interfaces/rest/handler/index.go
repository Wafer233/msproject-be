package handler

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type IndexHandler struct {
	service *service.GatewayIndexService
}

func NewIndexHandler(service *service.GatewayIndexService) *IndexHandler {
	return &IndexHandler{
		service: service,
	}
}

// Index 处理导航菜单请求
func (handler *IndexHandler) Index(ctx *gin.Context) {
	result := &common.Result{}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 调用服务层获取菜单
	dtoResp, err := handler.service.GetMenus(c)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.IndexIndexServiceFail, "调用Index服务失败"))
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, result.Success(dtoResp))
}
