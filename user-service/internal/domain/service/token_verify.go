package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
)

type TokenVerifyService interface {
	TokenVerify(ctx context.Context, token string) (*dto.TokenVerifyResponse, error)
}
