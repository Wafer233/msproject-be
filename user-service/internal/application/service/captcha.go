package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	repo "github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"time"
)

type CaptchaService interface {
	GenerateCaptcha(ctx context.Context, mobile string) (string, error)
	ValidateCaptcha(ctx context.Context, mobile, code string) (bool, error)
}

type DefaultCaptchaService struct {
	captchaRepo repo.CaptchaRepo
}

func (service *DefaultCaptchaService) ValidateCaptcha(ctx context.Context, mobile, code string) (bool, error) {
	key := model.KeyRegister + mobile
	redisCode, err := service.captchaRepo.Get(ctx, key)
	if err != nil {
		return false, err
	}

	if redisCode != code {
		return false, nil
	}

	return true, nil
}

func (service *DefaultCaptchaService) GenerateCaptcha(ctx context.Context, mobile string) (string, error) {
	code := "123456"
	key := model.KeyRegister + mobile
	err := service.captchaRepo.Put(ctx, key, code, 15*time.Minute)
	if err != nil {
		return "", err
	}
	return code, nil
}

func NewDefaultCaptchaService(captchaRepo repo.CaptchaRepo) CaptchaService {
	return &DefaultCaptchaService{
		captchaRepo: captchaRepo,
	}
}
