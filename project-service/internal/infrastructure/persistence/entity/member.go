package entity

import "github.com/Wafer233/msproject-be/project-service/internal/domain/model"

type ProjectMemberEntity struct {
	Id          int64  `gorm:"column:id;primaryKey"`
	ProjectCode string `gorm:"column:project_code"`
	MemberCode  string `gorm:"column:member_code"`
	JoinTime    string `gorm:"column:join_time"`
	IsOwner     int    `gorm:"column:is_owner"`
	Authorize   string `gorm:"column:authorize"`
}

// TableName returns the table name for GORM
func (ProjectMemberEntity) TableName() string {
	return "ms_project_member"
}

// ToModel converts entity to domain model
func (e *ProjectMemberEntity) ToModel() *model.ProjectMember {
	return &model.ProjectMember{
		Id:          e.Id,
		ProjectCode: e.ProjectCode,
		MemberCode:  e.MemberCode,
		JoinTime:    e.JoinTime,
		IsOwner:     e.IsOwner,
		Authorize:   e.Authorize,
	}
}
