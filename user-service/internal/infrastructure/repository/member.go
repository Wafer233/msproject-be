package dao

import (
	"context"
	"errors"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
	"github.com/Wafer233/msproject-be/user-service/internal/domain/repository"
	"github.com/Wafer233/msproject-be/user-service/internal/infrastructure/persistence/entity"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type GORMMemberDAO struct {
	db *gorm.DB
}

func (dao *GORMMemberDAO) ExistByEmail(ctx context.Context, email string) (bool, error) {
	var count int64

	err := dao.db.WithContext(ctx).
		Model(&entity.Member{}).
		Where("email = ?", email).
		Count(&count).Error

	return count > 0, err

}

func (dao *GORMMemberDAO) ExistByAccount(ctx context.Context, account string) (bool, error) {
	var count int64

	err := dao.db.WithContext(ctx).
		Model(&entity.Member{}).
		Where("account = ?", account).
		Count(&count).Error

	return count > 0, err
}

func (dao *GORMMemberDAO) ExistByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64

	err := dao.db.WithContext(ctx).
		Model(&entity.Member{}).
		Where("mobile = ?", mobile).
		Count(&count).Error

	return count > 0, err
}

func (dao *GORMMemberDAO) Save(ctx context.Context, member *model.Member) error {
	var entityMember entity.Member
	err := copier.Copy(&entityMember, member)
	if err != nil {
		return errors.New("member实体和领域模型转换错误")
	}

	err = dao.db.WithContext(ctx).
		Model(&entity.Member{}).
		Create(&entityMember).Error

	if err != nil {
		return errors.New("创建member失败")
	}

	member.Id = entityMember.ID

	return nil
}

func (dao *GORMMemberDAO) GetByCredentials(ctx context.Context, account string, password string) (*model.Member, error) {

	var entityMember entity.Member

	err := dao.db.WithContext(ctx).
		Model(&entity.Member{}).
		Where("account = ? AND password = ?", account, password).
		First(&entityMember).Error

	if err != nil {
		return nil, err
	}

	var member model.Member
	err = copier.Copy(&member, &entityMember)
	if err != nil {
		return nil, errors.New("member实体和领域模型转换错误")
	}

	return &member, nil
}

func NewGORMMemberRepository(db *gorm.DB) repository.MemberRepo {
	return &GORMMemberDAO{db: db}
}
