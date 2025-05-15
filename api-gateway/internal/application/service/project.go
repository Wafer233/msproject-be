package service

import (
	"context"
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	pb "github.com/Wafer233/msproject-be/api-gateway/proto/project"
	"github.com/jinzhu/copier"
)

type GatewayProjectService struct {
	client pb.ProjectServiceClient
}

func NewGatewayProjectService(client pb.ProjectServiceClient) *GatewayProjectService {
	return &GatewayProjectService{
		client: client,
	}
}

func (service *GatewayProjectService) GetMyProjects(
	ctx context.Context,
	page *dto.DTOPage,
	selectBy string,
	memberId int64,
	memberName string,
) ([]*dto.DTOProject, int64, error) {

	grpcMsg := &pb.SelfListMessage{
		MemberId:   memberId,
		MemberName: memberName,
		SelectBy:   selectBy,
		Page:       page.Page,
		PageSize:   page.PageSize,
	}

	grpcResp, err := service.client.SelfList(ctx, grpcMsg)
	if err != nil {
		return nil, 0, err
	}

	total := grpcResp.Total

	var list []*dto.DTOProject
	err = copier.Copy(list, grpcResp.List)
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
