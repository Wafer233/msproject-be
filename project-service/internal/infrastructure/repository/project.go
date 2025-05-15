package impl

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
	"gorm.io/gorm"
)

type GORMProjectRepository struct {
	db *gorm.DB
}

func NewGORMProjectRepository(db *gorm.DB) repository.ProjectRepository {
	return &GORMProjectRepository{db: db}
}

func (r *GORMProjectRepository) FindProjectsByMemberId(ctx context.Context, memberId int64, page, pageSize int64) ([]*model.ProjectWithMember, int64, error) {
	var total int64
	offset := (page - 1) * pageSize

	// 计算总项目数
	if err := r.db.WithContext(ctx).
		Table("ms_project_member").
		Where("member_code = ?", memberId).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询项目与成员关联数据
	var results []*model.ProjectWithMember
	err := r.db.WithContext(ctx).
		Table("ms_project p").
		Joins("JOIN ms_project_member pm ON p.id = pm.project_code").
		Select("p.*, pm.member_code, pm.join_time, pm.is_owner, pm.authorize, '' as owner_name, 0 as collected").
		Where("pm.member_code = ?", memberId).
		Limit(int(pageSize)).
		Offset(int(offset)).
		Scan(&results).Error

	if err != nil {
		return nil, 0, err
	}

	return results, total, nil
}
