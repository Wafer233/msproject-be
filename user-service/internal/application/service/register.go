package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
)

type RegisterService interface {
	Register(ctx context.Context, req dto.RegisterRequest) error
}
