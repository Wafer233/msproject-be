package login_service

import (
	"context"
	"errors"
	common "github.com/Wafer233/msproject-be/project-common"
	"github.com/Wafer233/msproject-be/project-user/pkg/dao"
	"github.com/Wafer233/msproject-be/project-user/pkg/repo"
	"go.uber.org/zap"
	"log"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {

	//1. 获取参数
	mobile := msg.Mobile
	//2. 验证手机合法性
	if !common.VerifyMobile(mobile) {
		return nil, errors.New("手机号不合法")
	}
	//3.生成验证码
	code := "123456"
	//4. 发送验证码
	go func() {
		time.Sleep(2 * time.Second)
		zap.L().Info("调用短信平台发送短信")
		//发送成功 存入redis
		err := ls.cache.Put("REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Println("验证码存入redis发生错误，cause by :", err)
		}
		zap.L().Info("发送短信成功")
	}()
	return &CaptchaResponse{}, nil

}
