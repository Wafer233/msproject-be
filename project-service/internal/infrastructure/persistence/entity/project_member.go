package entity

type ProjectMember struct {
	Id          int64  `gorm:"column:id;primaryKey;autoIncrement"`
	ProjectCode int64  `gorm:"column:project_code"`
	MemberCode  int64  `gorm:"column:member_code"`
	JoinTime    int64  `gorm:"column:join_time"`
	IsOwner     int64  `gorm:"column:is_owner"`
	Authorize   string `gorm:"column:authorize"`
}

func (*ProjectMember) TableName() string {
	return "ms_project_member"
}
