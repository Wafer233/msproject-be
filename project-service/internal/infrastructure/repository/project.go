package dao

import (
	"context"
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
	repo "github.com/Wafer233/msproject-be/project-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/project-service/internal/infrastructure/persistence/entity"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GORMProjectDAO struct {
	db *gorm.DB
}

func (dao *GORMProjectDAO) GetByMemberId(ctx context.Context, memberId int64, condition string, page int64, size int64) ([]*model.ProjectAndMember, int64, error) {
	var projAndMems []*model.ProjectAndMember
	var total int64

	offset := (page - 1) * size

	// 主查询：分页项目数据
	query := dao.db.WithContext(ctx).
		Table("ms_project AS a").
		Select("a.*, b.member_code").
		Joins("JOIN ms_project_member b ON a.id = b.project_code").
		Where("b.member_code = ?", memberId)

	if condition != "" {
		query = query.Where(condition)
	}

	// 1️⃣ 查询分页数据
	err := query.
		Order("a.sort").
		Offset(int(offset)).
		Limit(int(size)).
		Scan(&projAndMems).Error
	if err != nil {
		zap.L().Warn("查询project失败")
		return nil, 0, err
	}

	// 2️⃣ 单独查询总数（不要复用 query）
	countQuery := dao.db.WithContext(ctx).
		Table("ms_project AS a").
		Joins("JOIN ms_project_member b ON a.id = b.project_code").
		Where("b.member_code = ?", memberId)

	if condition != "" {
		countQuery = countQuery.Where(condition)
	}

	err = countQuery.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return projAndMems, total, nil
}

func (dao *GORMProjectDAO) GetCollectByMemId(ctx context.Context, memberId int64, page int64, size int64) ([]*model.ProjectAndMember, int64, error) {
	var projAndMems []*model.ProjectAndMember
	var total int64

	offset := (page - 1) * size

	// 查询分页数据
	err := dao.db.Table("ms_project AS p").
		Joins("JOIN ms_project_collection pc ON pc.project_code = p.id").
		Where("pc.member_code = ?", memberId).
		Order("p.sort").
		Offset(int(offset)).
		Limit(int(size)).
		Scan(&projAndMems).Error
	if err != nil {
		zap.L().Warn("查询分页错误")
		return nil, 0, err
	}

	// 查询总条数
	err = dao.db.Model(&entity.ProjectCollection{}).
		Where("member_code = ?", memberId).
		Count(&total).Error
	if err != nil {
		zap.L().Warn("查询total失败")
		return nil, 0, err
	}

	return projAndMems, total, nil

}

func NewGORMProjectRepo(db *gorm.DB) repo.ProjectRepo {
	return &GORMProjectDAO{
		db: db,
	}
}
