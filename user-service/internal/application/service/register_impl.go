package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
	pb "github.com/Wafer233/msproject-be/user-service/proto/login"
)

type DefaultRegisterService struct {
	pb.UnimplementedLoginServiceServer
}

func (d DefaultRegisterService) Register(ctx context.Context, req dto.RegisterRequest) error {
	//TODO implement me
	panic("implement me")
}
