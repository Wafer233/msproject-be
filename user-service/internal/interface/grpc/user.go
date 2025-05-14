package grpc

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	pb "github.com/Wafer233/msproject-be/user-service/proto/user"
	"go.uber.org/zap"
	"strconv"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	userService service.UserService
}

func NewUserServiceServer(userService service.UserService) *UserServiceServer {
	return &UserServiceServer{
		userService: userService,
	}
}

func (s *UserServiceServer) GetOrganizationsByMemberId(ctx context.Context, req *pb.GetOrgListRequest) (*pb.GetOrgListResponse, error) {
	// 获取成员的组织
	organizations, err := s.userService.GetOrganizationsByMemberId(ctx, req.MemberId)
	if err != nil {
		zap.L().Error("获取组织失败", zap.Error(err))
		return nil, err
	}

	// 转换为proto响应
	resp := &pb.GetOrgListResponse{}
	resp.OrganizationList = make([]*pb.OrganizationDTO, 0, len(organizations))

	for _, org := range organizations {
		orgDTO := &pb.OrganizationDTO{
			Id:          org.Id,
			Name:        org.Name,
			Avatar:      org.Avatar,
			Description: org.Description,
			OwnerCode:   org.MemberId,
			CreateTime:  "",
			Personal:    org.Personal,
			Code:        "", // 需要处理
			Address:     org.Address,
			Province:    org.Province,
			City:        org.City,
			Area:        org.Area,
		}

		// 处理创建时间，如果它是整数类型的时间戳
		if org.CreateTime > 0 {
			orgDTO.CreateTime = strconv.FormatInt(org.CreateTime, 10)
		}

		resp.OrganizationList = append(resp.OrganizationList, orgDTO)
	}

	return resp, nil
}
