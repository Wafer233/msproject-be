package service

import "context"

type GetCaptchaService interface {
	GetCaptcha(ctx context.Context, mobile string) (string, error)
}
