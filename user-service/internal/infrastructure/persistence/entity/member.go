package entity

import (
	"github.com/Wafer233/msproject-be/user-service/internal/domain/model"
)

type MemberEntity struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement"`
	Account       string `gorm:"column:account;type:varchar(100);not null;index"`
	Password      string `gorm:"column:password;type:varchar(100);not null"`
	Name          string `gorm:"column:name;type:varchar(50)"`
	Mobile        string `gorm:"column:mobile;type:varchar(20);index"`
	Email         string `gorm:"column:email;type:varchar(100);index"`
	Status        int    `gorm:"column:status;type:tinyint;default:1"`
	CreateTime    int64  `gorm:"column:create_time;type:bigint"`
	LastLoginTime int64  `gorm:"column:last_login_time;type:bigint"`
	Avatar        string `gorm:"column:avatar;type:text"`
	Description   string `gorm:"column:description;type:text"`
	Address       string `gorm:"column:address;type:varchar(100)"`
	Province      int    `gorm:"column:province;type:int;default:0"`
	City          int    `gorm:"column:city;type:int;default:0"`
	Area          int    `gorm:"column:area;type:int;default:0"`
}

// TableName 设置表名
func (MemberEntity) TableName() string {
	return "ms_member"
}

// ToModel 实体转领域模型
func (e *MemberEntity) ToModel() *model.Member {
	return &model.Member{
		Id:            e.ID,
		Account:       e.Account,
		Password:      e.Password,
		Name:          e.Name,
		Mobile:        e.Mobile,
		Email:         e.Email,
		Status:        e.Status,
		CreateTime:    e.CreateTime,
		LastLoginTime: e.LastLoginTime,
		Avatar:        e.Avatar,
		Description:   e.Description,
		Address:       e.Address,
		Province:      e.Province,
		City:          e.City,
		Area:          e.Area,
	}
}

// FromModel 领域模型转实体
func (e *MemberEntity) FromModel(m *model.Member) {
	e.ID = m.Id
	e.Account = m.Account
	e.Password = m.Password
	e.Name = m.Name
	e.Mobile = m.Mobile
	e.Email = m.Email
	e.Status = m.Status
	e.CreateTime = m.CreateTime
	e.LastLoginTime = m.LastLoginTime
	e.Avatar = m.Avatar
	e.Description = m.Description
	e.Address = m.Address
	e.Province = m.Province
	e.City = m.City
	e.Area = m.Area
}
