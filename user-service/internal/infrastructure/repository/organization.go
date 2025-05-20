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

type GORMOrganizationDAO struct {
	db *gorm.DB
}

func (dao *GORMOrganizationDAO) Save(ctx context.Context, org *model.Organization) error {

	var entityOrganization entity.Organization

	err := copier.Copy(&entityOrganization, org)
	if err != nil {
		return errors.New("组织模型与实体copy失败")
	}

	err = dao.db.Model(&entity.Organization{}).
		Create(&entityOrganization).Error

	if err != nil {
		return errors.New("创建组织失败")
	}

	return nil
}

func (dao *GORMOrganizationDAO) FindByMemberId(ctx context.Context, memberId int64) ([]*model.Organization, error) {

	var domainOrganizations []*model.Organization

	var entityOrganizations []*entity.Organization

	err := dao.db.Model(&entity.Organization{}).
		Where("member_id = ?", memberId).
		Find(&entityOrganizations).Error

	if err != nil {
		return nil, errors.New("查询组织失败")
	}

	er := copier.Copy(&domainOrganizations, &entityOrganizations)
	if er != nil {
		return nil, errors.New("组织模型与实体copy失败")
	}

	return domainOrganizations, nil

}

func NewGORMOrganizationRepository(db *gorm.DB) repository.OrganizationRepo {
	return &GORMOrganizationDAO{db: db}
}
