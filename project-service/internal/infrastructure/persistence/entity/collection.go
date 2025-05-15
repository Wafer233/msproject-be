package entity

import "github.com/Wafer233/msproject-be/project-service/internal/domain/model"

type ProjectCollectionEntity struct {
	Id          int64 `gorm:"column:id;primaryKey;autoIncrement"`
	ProjectCode int64 `gorm:"column:project_code"`
	MemberCode  int64 `gorm:"column:member_code"`
	CreateTime  int64 `gorm:"column:create_time"`
}

func (ProjectCollectionEntity) TableName() string {
	return "ms_project_collection"
}

func (p *ProjectCollectionEntity) ToModel() *model.ProjectCollectionModel {
	return &model.ProjectCollectionModel{
		Id:          p.Id,
		ProjectCode: p.ProjectCode,
		MemberCode:  p.MemberCode,
		CreateTime:  p.CreateTime,
	}
}
