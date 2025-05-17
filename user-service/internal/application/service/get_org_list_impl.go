package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
	pb "github.com/Wafer233/msproject-be/user-service/proto/organization"
)

type DefaultGetOrgListService struct {
	pb.UnimplementedOrganizationServiceServer
}

func (d DefaultGetOrgListService) GetOrgList(ctx context.Context, memberId int64) ([]*dto.GetOrgListResponse, error) {
	//TODO implement me
	panic("implement me")
}
