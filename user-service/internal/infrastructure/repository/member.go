package impl

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/persistence/entity"
	"gorm.io/gorm"
)

type GORMMemberRepository struct {
	db *gorm.DB
}

func NewGORMMemberRepository(db *gorm.DB) repository.MemberRepository {
	return &GORMMemberRepository{db: db}
}

func (gmr GORMMemberRepository) FindMemberByAccount(ctx context.Context, account string) (bool, error) {
	var count int64
	err := gmr.db.WithContext(ctx).Model(&entity.MemberEntity{}).
		Where("account = ?", account).
		Count(&count).Error

	return count > 0, err
}

func (gmr GORMMemberRepository) SaveMember(ctx context.Context, member *model.Member) error {
	var entity entity.MemberEntity
	entity.FromModel(member)

	err := gmr.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return err
	}

	// 更新ID
	member.Id = entity.ID
	return nil
}
