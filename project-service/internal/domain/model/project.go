package model

type Project struct {
	Id                 int64
	Cover              string
	Name               string
	Code               string
	Description        string
	AccessControlType  string
	WhiteList          string
	Order              int
	Deleted            int
	TemplateCode       string
	Schedule           float64
	CreateTime         string
	OrganizationCode   string
	DeletedTime        string
	Private            int
	Prefix             string
	OpenPrefix         int
	Archive            int
	ArchiveTime        int64
	OpenBeginTime      int
	OpenTaskPrivate    int
	TaskBoardTheme     string
	BeginTime          int64
	EndTime            int64
	AutoUpdateSchedule int
}
