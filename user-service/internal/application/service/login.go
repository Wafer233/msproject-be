package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
)

type LoginService interface {
	Login(ctx context.Context, dtoReq *dto.LoginRequest) (*dto.LoginResponse, error)
}
