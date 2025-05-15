package middleware

import (
	"context"
	pb "github.com/Wafer233/msproject-be/api-gateway/proto/login"
)

type TokenVerifyService struct {
	client pb.LoginServiceClient
}

func (service *TokenVerifyService) VerifyToken(ctx context.Context, token string) (*pb.TokenVerifyResponse, error) {
	rpcMsg := &pb.TokenVerifyMessage{
		Token: token,
	}
	rpcResp, err := service.client.TokenVerify(ctx, rpcMsg)

	if err != nil {
		return nil, err
	}

	return rpcResp, nil
}
