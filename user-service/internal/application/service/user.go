package service

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/application/dto"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type DefaultUserService struct {
	or repository.OrganizationRepository
}

func NewDefaultUserService(or repository.OrganizationRepository) UserService {
	return &DefaultUserService{
		or: or,
	}
}

func (s *DefaultUserService) GetOrgList(ctx context.Context, memberId int64) ([]dto.OrganizationDTO, error) {
	// 调用仓库获取组织列表
	organizations, err := s.or.FindOrganizationsByMemberId(ctx, memberId)
	if err != nil {
		zap.L().Error("获取组织列表失败", zap.Error(err))
		return nil, err
	}

	// 转换为DTO
	orgDTOs := make([]dto.OrganizationDTO, 0, len(organizations))
	for _, org := range organizations {
		var dto dto.OrganizationDTO
		if err := copier.Copy(&dto, &org); err != nil {
			zap.L().Error("转换组织DTO失败", zap.Error(err))
			continue
		}

		// 转换ID为Code (假设有加密函数)
		//dto.Code = strconv.FormatInt(org.Id, 10) // 实际项目中应使用加密函数

		orgDTOs = append(orgDTOs, dto)
	}

	return orgDTOs, nil
}
