package handler

import (
	"context"
	"fmt"
	"github.com/Wafer233/msproject-be/api-gateway/internal/dto"
	pb "github.com/Wafer233/msproject-be/api-gateway/proto/login"
	"github.com/Wafer233/msproject-be/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"time"
)

type LoginHttpHandler struct {
	client pb.LoginServiceClient
}

func NewLoginHttpHandler(client pb.LoginServiceClient) *LoginHttpHandler {
	return &LoginHttpHandler{
		client: client,
	}
}

func (handler *LoginHttpHandler) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	mobile := ctx.PostForm("mobile")

	if mobile == "" {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginGetCaptchaMobileEmpty, "手机号不能为空"))
		return
	}

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	grpcReq := &pb.GetCaptchaRequest{
		Mobile: mobile,
	}

	grpcResp, err := handler.client.GetCaptcha(c, grpcReq)

	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginGetCaptchaServiceFail, "验证码服务失败"))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(grpcResp.Code))
}

func (handler *LoginHttpHandler) Login(ctx *gin.Context) {
	result := &common.Result{}

	dtoReq := &dto.LoginRequest{}

	err := ctx.ShouldBind(&dtoReq)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginLoginRequestBindFail, "登录请求绑定失败"))
		return
	}

	grpcReq := &pb.LoginRequest{}
	err = copier.Copy(grpcReq, dtoReq)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginLoginCopyFail, "登录服务复制有误"))
		return
	}

	//msg.Ip = GetIp(c)
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpcResp, er := handler.client.Login(c, grpcReq)
	if er != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginLoginServiceFail, "登陆服务失败"))
		return
	}

	dtoResp := &dto.LoginResponse{}
	err = copier.Copy(dtoResp, grpcResp)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginLoginCopyFail, "登录服务复制有误"))
		return
	}

	ctx.JSON(http.StatusOK, result.Success(dtoResp))
}

func (handler *LoginHttpHandler) Register(ctx *gin.Context) {
	result := &common.Result{}

	dtoReq := &dto.RegisterRequest{}
	err := ctx.Bind(&dtoReq)

	fmt.Print(dtoReq)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginRegisterRequestBindFail, "注册请求绑定失败"))
		return
	}

	er := common.Verify(dtoReq.Email, dtoReq.Mobile, dtoReq.Password, dtoReq.Password2)
	if er != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginRegisterRequestVerifyFail, "注册请求信息格式有误"))
		return
	}

	grpcReq := &pb.RegisterRequest{}
	err = copier.Copy(grpcReq, dtoReq)

	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginRegisterCopyFail, "注册请求复制有误"))
		return
	}

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = handler.client.Register(c, grpcReq)

	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginRegisterServiceFail, "注册服务失败"))
		return
	}
	//4.返回结果
	ctx.JSON(http.StatusOK, result.Success(""))
}

func (handler *LoginHttpHandler) GetOrgList(ctx *gin.Context) {
	result := &common.Result{}

	memberIdStr, exist := ctx.Get("memberId")
	if !exist {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginGetOrgListMemberNotExist, "会员Id不存在"))
		return
	}

	memberId := memberIdStr.(int64)

	grpcReq := &pb.GetOrgListRequest{
		MemberId: memberId,
	}

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grpcResp, er := handler.client.GetOrgList(c, grpcReq)

	if er != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginGetOrgListServiceFail, "获取组织服务失败"))
		return
	}

	if grpcResp.OrganizationList == nil {
		ctx.JSON(http.StatusOK, result.Success([]*dto.OrganizationList{}))
		return
	}

	organizations := []*dto.OrganizationList{}
	err := copier.Copy(&organizations, grpcResp.OrganizationList)

	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(common.LoginGetOrgListCopyFail, "组织列表复制失败"))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(organizations))
}
