package entity

import "github.com/Wafer233/msproject-be/project-service/internal/domain/model"

type ProjectMemberEntity struct {
	Id          int64  `gorm:"column:id;primaryKey"`
	ProjectCode int64  `gorm:"column:project_code"`
	MemberCode  int64  `gorm:"column:member_code"`
	JoinTime    string `gorm:"column:join_time"`
	IsOwner     int    `gorm:"column:is_owner"`
	Authorize   string `gorm:"column:authorize"`
}

// TableName 设置GORM的表名
func (ProjectMemberEntity) TableName() string {
	return "ms_project_member"
}

// ToModel 将实体转换为领域模型
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
