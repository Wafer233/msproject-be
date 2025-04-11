package service

import (
	"context"
	"errors"
	"github.com/Wafer233/msproject-be/common"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"time"
)

type CaptchaService struct {
	captchaRepo repository.CaptchaRepository
	// 可能还需要短信服务等依赖
}

func NewCaptchaService(captchaRepo repository.CaptchaRepository) *CaptchaService {
	return &CaptchaService{
		captchaRepo: captchaRepo,
	}
}

// GenerateCaptcha 生成并发送验证码
func (s *CaptchaService) GenerateCaptcha(ctx context.Context, mobile string) (string, error) {
	// 1. 验证手机合法性
	if !common.VerifyMobile(mobile) {
		return "", errors.New("invalid mobile number")
	}

	// 2. 生成验证码
	code := "123456" // 实际中应随机生成

	// 3. 发送验证码
	// 实际项目中应该调用短信服务发送验证码

	// 4. 保存验证码到存储
	err := s.captchaRepo.SaveCaptcha(ctx, mobile, code, 15*time.Minute)
	if err != nil {
		return "", err
	}

	return code, nil
}
