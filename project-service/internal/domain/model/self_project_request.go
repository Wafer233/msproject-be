package model

type SelfProjectRequest struct {
	MemberID         int64
	MemberName       string
	Page             int64
	PageSize         int64
	SelectBy         string
	OrganizationCode string
	ViewType         int32
	Name             string
	TemplateCode     string
	Description      string
	ID               int64
	ProjectCode      string
	Deleted          bool
	CollectType      string
	TaskCode         string
}
