package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequest) error
	Login(ctx context.Context, req dto.LoginReq) (*dto.LoginRsp, error)
	TokenVerify(ctx context.Context, token string) (*dto.Member, error)
}
