package login

import (
	common "github.com/Wafer233/msproject-be/project-common"
	"github.com/Wafer233/msproject-be/project-user/pkg/dao"
	"github.com/Wafer233/msproject-be/project-user/pkg/model"
	"github.com/Wafer233/msproject-be/project-user/pkg/repo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"time"
)

type HandlerLogin struct {
	cache repo.Cache
}

func New() *HandlerLogin {
	return &HandlerLogin{
		cache: dao.Rc,
	}
}

// GetCaptcha 获取手机验证码
func (hl *HandlerLogin) GetCaptcha(ctx *gin.Context) {
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
		zap.L().Info("调用短信平台发送短信")
		//发送成功 存入redis
		err := hl.cache.Put("REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Println("验证码存入redis发生错误，cause by :", err)
		}
		zap.L().Info("发送短信成功")
	}()
	ctx.JSON(200, result.Success("123456"))
}
