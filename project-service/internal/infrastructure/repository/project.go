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

	// Count total projects
	if err := r.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM ms_project a JOIN ms_project_member b ON a.id = b.project_code WHERE b.member_code = ?", memberId).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Query projects with member relation
	var results []*model.ProjectWithMember
	err := r.db.WithContext(ctx).
		Raw(`SELECT a.*, b.member_code, b.join_time, b.is_owner, b.authorize, 
			'子龙' as owner_name, 0 as collected 
			FROM ms_project a JOIN ms_project_member b 
			ON a.id = b.project_code 
			WHERE b.member_code = ? LIMIT ? OFFSET ?`,
			memberId, pageSize, offset).
		Scan(&results).Error

	if err != nil {
		return nil, 0, err
	}

	return results, total, nil
}
