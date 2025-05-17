package service

import (
	"context"
)

type CaptchaService interface {
	GenerateCaptcha(ctx context.Context, mobile string) (string, error)
}
