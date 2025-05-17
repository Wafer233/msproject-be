package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
	pb "github.com/Wafer233/msproject-be/user-service/proto/login"
)

type DefaultLoginService struct {
	pb.UnimplementedLoginServiceServer
}

func (d DefaultLoginService) Login(ctx context.Context, dtoReq *dto.LoginRequest) (*dto.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}
