package model

// ProjectMember 项目与成员的关联
type ProjectMember struct {
	Id          int64
	ProjectCode int64
	MemberCode  int64
	JoinTime    string
	IsOwner     int
	Authorize   string
}

// TableName 返回数据库表名
func (*ProjectMember) TableName() string {
	return "ms_project_member"
}
