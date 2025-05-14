package dto

type ProjectDTO struct {
	Id                 int64   `json:"id"`
	Cover              string  `json:"cover"`
	Name               string  `json:"name"`
	Code               string  `json:"code"`
	Description        string  `json:"description"`
	AccessControlType  string  `json:"access_control_type"`
	WhiteList          string  `json:"white_list"`
	Order              int     `json:"order"`
	Deleted            int     `json:"deleted"`
	TemplateCode       string  `json:"template_code"`
	Schedule           float64 `json:"schedule"`
	CreateTime         string  `json:"create_time"`
	OrganizationCode   string  `json:"organization_code"`
	DeletedTime        string  `json:"deleted_time"`
	Private            int     `json:"private"`
	Prefix             string  `json:"prefix"`
	OpenPrefix         int     `json:"open_prefix"`
	Archive            int     `json:"archive"`
	ArchiveTime        int64   `json:"archive_time"`
	OpenBeginTime      int     `json:"open_begin_time"`
	OpenTaskPrivate    int     `json:"open_task_private"`
	TaskBoardTheme     string  `json:"task_board_theme"`
	BeginTime          int64   `json:"begin_time"`
	EndTime            int64   `json:"end_time"`
	AutoUpdateSchedule int     `json:"auto_update_schedule"`
	ProjectCode        string  `json:"project_code"`
	MemberCode         string  `json:"member_code"`
	JoinTime           string  `json:"join_time"`
	IsOwner            int     `json:"is_owner"`
	Authorize          string  `json:"authorize"`
	OwnerName          string  `json:"owner_name"`
	Collected          int     `json:"collected"`
}

// ProjectListResponse represents a response containing a list of projects
type ProjectListResponse struct {
	List  []*ProjectDTO `json:"list"`
	Total int64         `json:"total"`
}

// ProjectRequest represents a project list request
type ProjectRequest struct {
	MemberId int64 `json:"memberId"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}
