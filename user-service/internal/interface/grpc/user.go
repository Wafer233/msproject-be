package grpc

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/service"
	pb "github.com/Wafer233/msproject-be/user-service/proto/user"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type UserServiceServer struct {
	pb.UnimplementedOrganizationServiceServer
	userService service.UserService
}

// NewUserServiceServer 创建用户服务gRPC处理器
func NewUserServiceServer(userService service.UserService) *UserServiceServer {
	return &UserServiceServer{
		userService: userService,
	}
}

// GetOrgList 获取用户组织列表
func (s *UserServiceServer) GetOrgList(ctx context.Context, req *pb.OrgListRequest) (*pb.OrgListResponse, error) {
	// 调用应用服务
	orgs, err := s.userService.GetOrgList(ctx, req.MemberId)
	if err != nil {
		zap.L().Error("获取组织列表失败", zap.Error(err))
		return nil, err
	}

	// 转换为gRPC响应
	response := &pb.OrgListResponse{
		OrganizationList: make([]*pb.OrganizationDTO, 0, len(orgs)),
	}

	for _, org := range orgs {
		pbOrg := &pb.OrganizationDTO{}
		if err := copier.Copy(pbOrg, &org); err != nil {
			zap.L().Error("组织DTO转换失败", zap.Error(err))
			continue
		}
		response.OrganizationList = append(response.OrganizationList, pbOrg)
	}

	return response, nil
}
