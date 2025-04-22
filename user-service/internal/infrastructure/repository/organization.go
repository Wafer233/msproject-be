package impl

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/persistence/entity"
	"gorm.io/gorm"
)

type CacheOrganizationRepository struct {
	db *gorm.DB
}

func NewCachedOrganizationRepository(db *gorm.DB) repository.OrganizationRepository {
	return &CacheOrganizationRepository{db: db}
}

func (cor CacheOrganizationRepository) SaveOrganization(ctx context.Context, org *model.Organization) error {
	var entity entity.OrganizationEntity
	entity.FromModel(org)

	err := cor.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return err
	}

	// 更新ID
	org.Id = entity.ID
	return nil
}
