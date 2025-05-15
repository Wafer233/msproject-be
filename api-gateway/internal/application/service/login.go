package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	pb "github.com/Wafer233/msproject-be/api-gateway/proto/login"
	"github.com/jinzhu/copier"
)

type GatewayLoginService struct {
	client pb.LoginServiceClient
}

func NewGatewayLoginService(client pb.LoginServiceClient) *GatewayLoginService {
	return &GatewayLoginService{
		client: client,
	}
}

func (service *GatewayLoginService) Login(ctx context.Context, dtoReq *dto.LoginRequest) (*dto.LoginResponse, error) {

	// 转换到gRPC请求
	grpcMsg := &pb.LoginMessage{}
	err := copier.Copy(grpcMsg, dtoReq)
	if err != nil {
		return &dto.LoginResponse{}, err
	}

	// 调用gRPC服务
	grpcResp, err := service.client.Login(ctx, grpcMsg)
	if err != nil {
		return &dto.LoginResponse{}, err
	}

	// 转换到DTO
	dtoResp := &dto.LoginResponse{}

	err = copier.Copy(dtoResp, grpcResp)
	if err != nil {
		return &dto.LoginResponse{}, err
	}

	return dtoResp, nil
}
