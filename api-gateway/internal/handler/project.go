package handler

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/dto"
	pb "github.com/Wafer233/msproject-be/api-gateway/proto/project"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
)

type ProjectHttpHandler struct {
	client pb.ProjectServiceClient
}

func NewProjectHttpHandler(client pb.ProjectServiceClient) *ProjectHttpHandler {
	return &ProjectHttpHandler{
		client: client,
	}
}

func (handler *ProjectHttpHandler) Index(ctx *gin.Context) {
	result := &common.Result{}

	grpcReq := &pb.IndexRequest{}

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpcResp, err := handler.client.Index(c, grpcReq)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.ProjectIndexServiceFail, "Index服务失败"))
	}

	menus := grpcResp.MenuList
	var dtoResp []*dto.Menu

	er := copier.Copy(&dtoResp, menus)
	if er != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.ProjectIndexCopyFail, "Index服务复制失败"))
	}
	ctx.JSON(http.StatusOK, result.Success(dtoResp))
}

func (handler *ProjectHttpHandler) SelfProject(ctx *gin.Context) {
	result := &common.Result{}

	memberId := ctx.GetInt64("memberId")
	memberName := ctx.GetString("memberName")

	page := &dto.Page{}
	_ = ctx.ShouldBind(&page)
	if page.Page == 0 {
		page.Page = 1
	}
	if page.PageSize == 0 {
		page.PageSize = 10
	}

	selectBy := ctx.PostForm("selectBy")

	grpcReq := &pb.SelfProjectRequest{
		MemberId:   memberId,
		MemberName: memberName,
		SelectBy:   selectBy,
		Page:       page.Page,
		PageSize:   page.PageSize,
	}

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	grpcResp, err := handler.client.SelfProject(c, grpcReq)

	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.ProjectSelfProjectServiceFail, "查到自己的项目服务失败"))
	}

	var projectAndMember []*dto.ProjectAndMember
	err = copier.Copy(&projectAndMember, grpcResp.Projects)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.ProjectSelfProjectCopyFail, "查到自己的项目服务复制失败"))
	}

	if projectAndMember == nil {
		projectAndMember = []*dto.ProjectAndMember{}
	}
	ctx.JSON(http.StatusOK, result.Success(gin.H{
		"list":  projectAndMember, //null nil -> []
		"total": grpcResp.Total,
	}))
}
