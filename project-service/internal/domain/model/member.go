package model

// ProjectMember represents the relationship between a project and a member
type ProjectMember struct {
	Id          int64
	ProjectCode string
	MemberCode  string
	JoinTime    string
	IsOwner     int
	Authorize   string
}
