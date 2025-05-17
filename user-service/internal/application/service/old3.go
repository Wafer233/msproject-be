package service

import (
	"context"
	"errors"
	"github.com/Wafer233/msproject-be/common"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"time"
)

//type DefaultptchaService struct {
//	cr repository.CaptchaRepository
//	// 可能还需要短信服务等依赖
//}
//
//func NewDefaultCachaService(cr repository.CaptchaRepository) CaptchaService {
//	return &DefaulCaptchaService{
//		cr: cr,
//	}
//}

func (dcs DefaultCaptchaService) GenerateCaptcha(ctx context.Context, mobile string) (string, error) {
	if !common.VerifyMobile(mobile) {
		return "", errors.New("invalid mobile number")
	}

	code := "123456" // 应生成随机码
	err := dcs.cr.SaveCaptcha(ctx, "REGISTER_"+mobile, code, 15*time.Minute)
	if err != nil {
		return "", err
	}

	return code, nil
}
