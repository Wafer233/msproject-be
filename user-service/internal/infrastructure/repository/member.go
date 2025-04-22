package impl

import (
	"context"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/persistence/entity"
	"gorm.io/gorm"
)

type CacheMemberRepository struct {
	db *gorm.DB
}

func NewCachedMemberRepository(db *gorm.DB) repository.MemberRepository {
	return &CacheMemberRepository{db: db}
}

func (cmr CacheMemberRepository) FindMemberByAccount(ctx context.Context, account string) (bool, error) {
	var count int64
	err := cmr.db.WithContext(ctx).Model(&entity.MemberEntity{}).
		Where("account = ?", account).
		Count(&count).Error

	return count > 0, err
}

func (cmr CacheMemberRepository) SaveMember(ctx context.Context, member *model.Member) error {
	var entity entity.MemberEntity
	entity.FromModel(member)

	err := cmr.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		return err
	}

	// 更新ID
	member.Id = entity.ID
	return nil
}
