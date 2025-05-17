package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
	pb "github.com/Wafer233/msproject-be/user-service/proto/login"
)

type DefaultTokenVerifyService struct {
	pb.UnimplementedLoginServiceServer
}

func (d DefaultTokenVerifyService) TokenVerify(ctx context.Context, token string) (*dto.TokenVerifyResponse, error) {
	//TODO implement me
	panic("implement me")
}
