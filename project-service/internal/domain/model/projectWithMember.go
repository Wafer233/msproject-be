package model

// ProjectWithMember combines Project and ProjectMember data
type ProjectWithMember struct {
	Project
	MemberCode string
	JoinTime   string
	IsOwner    int
	Authorize  string
	OwnerName  string
	Collected  int
}
