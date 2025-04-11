package repository

import (
	"context"
	"time"
)

type CaptchaRepository interface {
	SaveCaptcha(ctx context.Context, mobile string, code string, expiration time.Duration) error
	GetCaptcha(ctx context.Context, mobile string) (string, error)
}
