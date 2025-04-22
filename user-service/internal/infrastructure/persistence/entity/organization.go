package entity

import (
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
)

type OrganizationEntity struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string `gorm:"column:name;type:varchar(255)"`
	Avatar      string `gorm:"column:avatar;type:varchar(511)"`
	Description string `gorm:"column:description;type:varchar(500)"`
	MemberId    int64  `gorm:"column:member_id;type:bigint"`
	CreateTime  int64  `gorm:"column:create_time;type:bigint"`
	Personal    int32  `gorm:"column:personal;type:tinyint;default:0"`
	Address     string `gorm:"column:address;type:varchar(100)"`
	Province    int32  `gorm:"column:province;type:int;default:0"`
	City        int32  `gorm:"column:city;type:int;default:0"`
	Area        int32  `gorm:"column:area;type:int;default:0"`
}

// TableName 设置表名
func (OrganizationEntity) TableName() string {
	return "ms_organization"
}

// ToModel 实体转领域模型
func (e *OrganizationEntity) ToModel() *model.Organization {
	return &model.Organization{
		Id:          e.ID,
		Name:        e.Name,
		Avatar:      e.Avatar,
		Description: e.Description,
		MemberId:    e.MemberId,
		CreateTime:  e.CreateTime,
		Personal:    e.Personal,
		Address:     e.Address,
		Province:    e.Province,
		City:        e.City,
		Area:        e.Area,
	}
}

// FromModel 领域模型转实体
func (e *OrganizationEntity) FromModel(m *model.Organization) {
	e.ID = m.Id
	e.Name = m.Name
	e.Avatar = m.Avatar
	e.Description = m.Description
	e.MemberId = m.MemberId
	e.CreateTime = m.CreateTime
	e.Personal = m.Personal
	e.Address = m.Address
	e.Province = m.Province
	e.City = m.City
	e.Area = m.Area
}
