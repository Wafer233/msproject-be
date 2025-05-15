package handler

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetOrgListHandler struct {
	service *service.GatewayGetOrgListService
}

func NewGetOrgListHandler(service *service.GatewayGetOrgListService) *GetOrgListHandler {
	return &GetOrgListHandler{
		service: service,
	}
}

func (handler *GetOrgListHandler) GetOrgList(ctx *gin.Context) {
	result := &common.Result{}

	// 从上下文获取用户ID（由认证中间件设置）- 注意键名一致性
	memberIdStr, _ := ctx.Get("memberId")
	memberId := memberIdStr.(int64)

	// --------------------- 调用服务 --------------------------------------------
	dtoResp, err := handler.service.GetOrgList(ctx, memberId)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.OganizationServiceFail, "调用GetOrgList服务失败"))
	}

	// 直接返回数组，与文档格式一致
	ctx.JSON(http.StatusOK, result.Success(dtoResp))
}
