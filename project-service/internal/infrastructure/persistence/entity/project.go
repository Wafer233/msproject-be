package entity

type Project struct {
	Id                 int64   `gorm:"column:id;primaryKey;autoIncrement"`
	Cover              string  `gorm:"column:cover"`
	Name               string  `gorm:"column:name"`
	Description        string  `gorm:"column:description"`
	AccessControlType  int     `gorm:"column:access_control_type"`
	WhiteList          string  `gorm:"column:white_list"`
	Sort               int     `gorm:"column:sort"`
	Deleted            int     `gorm:"column:deleted"`
	TemplateCode       int     `gorm:"column:template_code"`
	Schedule           float64 `gorm:"column:schedule"`
	CreateTime         int64   `gorm:"column:create_time"`
	OrganizationCode   int64   `gorm:"column:organization_code"`
	DeletedTime        string  `gorm:"column:deleted_time"`
	Private            int     `gorm:"column:private"`
	Prefix             string  `gorm:"column:prefix"`
	OpenPrefix         int     `gorm:"column:open_prefix"`
	Archive            int     `gorm:"column:archive"`
	ArchiveTime        int64   `gorm:"column:archive_time"`
	OpenBeginTime      int     `gorm:"column:open_begin_time"`
	OpenTaskPrivate    int     `gorm:"column:open_task_private"`
	TaskBoardTheme     string  `gorm:"column:task_board_theme"`
	BeginTime          int64   `gorm:"column:begin_time"`
	EndTime            int64   `gorm:"column:end_time"`
	AutoUpdateSchedule int     `gorm:"column:auto_update_schedule"`
}

func (*Project) TableName() string {
	return "ms_project"
}
