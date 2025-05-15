package entity

import "github.com/Wafer233/msproject-be/project-service/internal/domain/model"

type ProjectEntity struct {
	Id                 int64   `gorm:"column:id;primaryKey"`
	Cover              string  `gorm:"column:cover"`
	Name               string  `gorm:"column:name"`
	Description        string  `gorm:"column:description"`
	AccessControlType  string  `gorm:"column:access_control_type"`
	WhiteList          string  `gorm:"column:white_list"`
	Order              int     `gorm:"column:order"`
	Deleted            int     `gorm:"column:deleted"`
	TemplateCode       string  `gorm:"column:template_code"`
	Schedule           float64 `gorm:"column:schedule"`
	CreateTime         string  `gorm:"column:create_time"`
	OrganizationCode   string  `gorm:"column:organization_code"`
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

// TableName 设置GORM的表名
func (ProjectEntity) TableName() string {
	return "ms_project"
}

// ToModel 将实体转换为领域模型
func (e *ProjectEntity) ToModel() *model.Project {
	return &model.Project{
		Id:                 e.Id,
		Cover:              e.Cover,
		Name:               e.Name,
		Description:        e.Description,
		AccessControlType:  e.AccessControlType,
		WhiteList:          e.WhiteList,
		Order:              e.Order,
		Deleted:            e.Deleted,
		TemplateCode:       e.TemplateCode,
		Schedule:           e.Schedule,
		CreateTime:         e.CreateTime,
		OrganizationCode:   e.OrganizationCode,
		DeletedTime:        e.DeletedTime,
		Private:            e.Private,
		Prefix:             e.Prefix,
		OpenPrefix:         e.OpenPrefix,
		Archive:            e.Archive,
		ArchiveTime:        e.ArchiveTime,
		OpenBeginTime:      e.OpenBeginTime,
		OpenTaskPrivate:    e.OpenTaskPrivate,
		TaskBoardTheme:     e.TaskBoardTheme,
		BeginTime:          e.BeginTime,
		EndTime:            e.EndTime,
		AutoUpdateSchedule: e.AutoUpdateSchedule,
	}
}
