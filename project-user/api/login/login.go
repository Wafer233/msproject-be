package login

import (
	"fmt"
	common "github.com/Wafer233/msproject-be/project-common"
	"github.com/Wafer233/project-user/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type RouterLogin struct {
}

func (*RouterLogin) Register(r *gin.Engine) {
	g := r.Group("/project/login")
	h := HandlerLogin{}
	g.POST("/getCaptcha", h.GetCaptcha)
}

type HandlerLogin struct {
}

// GetCaptcha 获取手机验证码
func (HandlerLogin) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	//1. 获取参数
	mobile := ctx.PostForm("mobile")
	//2. 验证手机合法性
	if !common.VerifyMobile(mobile) {
		ctx.JSON(200, result.Fail(model.LoginMobileNotLegal, "不合法"))
		return
	}
	//3.生成验证码
	code := "123456"
	//4. 发送验证码
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("调用短信平台发送短信")
		//发送成功 存入redis
		fmt.Println(mobile, code)
	}()
	ctx.JSON(200, result.Success("123456"))
}
