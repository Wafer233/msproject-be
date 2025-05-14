package impl

import (
	"context"
	"errors"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/persistence/entity"
	"gorm.io/gorm"
)

type GORMMemberRepository struct {
	db *gorm.DB
}

func (gmr GORMMemberRepository) FindMemberById(ctx context.Context, id int64) (*model.Member, error) {
	var entity entity.MemberEntity

	err := gmr.db.WithContext(ctx).
		Where("id = ?", id).
		First(&entity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return entity.ToModel(), nil
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

func (gmr GORMMemberRepository) FindMember(ctx context.Context, account, password string) (*model.Member, error) {
	var entity entity.MemberEntity

	err := gmr.db.WithContext(ctx).
		Where("account = ?", account).
		First(&entity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	err = gmr.db.WithContext(ctx).
		Where("account = ? AND password = ?", account, password).
		First(&entity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在或密码错误")
		}
		return nil, err
	}

	return entity.ToModel(), nil
}
