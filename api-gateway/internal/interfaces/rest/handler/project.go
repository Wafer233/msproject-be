package handler

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/service"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ProjectHandler struct {
	service *service.GatewayProjectService
}

func NewProjectHandler(projectService *service.GatewayProjectService) *ProjectHandler {
	return &ProjectHandler{
		service: projectService,
	}
}

func (handler *ProjectHandler) SelfList(ctx *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	memberId := ctx.GetInt64("memberId")
	memberName := ctx.GetString("memberName")

	page := &dto.DTOPage{}
	page.Bind(ctx)
	selectBy := ctx.PostForm("selectBy")

	list, total, err := handler.service.GetMyProjects(c, page, selectBy, memberId, memberName)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.ProjectSelfListFail, "调用selfList服务失败"))
		return
	}

	ctx.JSON(http.StatusOK, result.Success(gin.H{
		"list":  list, //null nil -> []
		"total": total,
	}))
}
