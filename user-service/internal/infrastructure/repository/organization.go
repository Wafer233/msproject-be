package impl

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/persistence/entity"
	"gorm.io/gorm"
)

type GORMOrganizationRepository struct {
	db *gorm.DB
}

func NewGORMOrganizationRepository(db *gorm.DB) repository.OrganizationRepository {
	return &GORMOrganizationRepository{db: db}
}

func (gor GORMOrganizationRepository) SaveOrganization(ctx context.Context, org *model.Organization) error {
	var entity entity.OrganizationEntity
	entity.FromModel(org)

	err := gor.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return err
	}

	// 更新ID
	org.Id = entity.ID
	return nil
}

func (gor GORMOrganizationRepository) FindOrganizationsByMemberId(ctx context.Context, memberId int64) ([]model.Organization, error) {
	var entities []entity.OrganizationEntity
	err := gor.db.WithContext(ctx).
		Where("member_id = ?", memberId).
		Find(&entities).Error

	if err != nil {
		return nil, err
	}

	// 转换为领域模型
	organizations := make([]model.Organization, len(entities))
	for i, e := range entities {
		org := e.ToModel()
		organizations[i] = *org
	}

	return organizations, nil
}
