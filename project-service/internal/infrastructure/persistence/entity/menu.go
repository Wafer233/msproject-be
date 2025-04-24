package entity

import (
	"github.com/Wafer233/msproject-be/project-service/internal/domain/model"
)

// ProjectMenuEntity 表示数据库中的菜单项
type ProjectMenuEntity struct {
	Id         int64  `gorm:"column:id;primaryKey"`
	Pid        int64  `gorm:"column:pid"`
	Title      string `gorm:"column:title"`
	Icon       string `gorm:"column:icon"`
	Url        string `gorm:"column:url"`
	FilePath   string `gorm:"column:file_path"`
	Params     string `gorm:"column:params"`
	Node       string `gorm:"column:node"`
	Sort       int    `gorm:"column:sort"`
	Status     int    `gorm:"column:status"`
	CreateBy   int64  `gorm:"column:create_by"`
	CreateAt   string `gorm:"column:create_at"`
	IsInner    int    `gorm:"column:is_inner"`
	Values     string `gorm:"column:values"`
	ShowSlider int    `gorm:"column:show_slider"`
}

// TableName 设置GORM的表名
func (ProjectMenuEntity) TableName() string {
	return "ms_project_menu"
}

// ToModel 将实体转换为领域模型
func (e *ProjectMenuEntity) ToModel() *model.ProjectMenu {
	return &model.ProjectMenu{
		Id:         e.Id,
		Pid:        e.Pid,
		Title:      e.Title,
		Icon:       e.Icon,
		Url:        e.Url,
		FilePath:   e.FilePath,
		Params:     e.Params,
		Node:       e.Node,
		Sort:       e.Sort,
		Status:     e.Status,
		CreateBy:   e.CreateBy,
		IsInner:    e.IsInner,
		Values:     e.Values,
		ShowSlider: e.ShowSlider,
	}
}

// FromModel 将领域模型转换为实体
func (e *ProjectMenuEntity) FromModel(m *model.ProjectMenu) {
	e.Id = m.Id
	e.Pid = m.Pid
	e.Title = m.Title
	e.Icon = m.Icon
	e.Url = m.Url
	e.FilePath = m.FilePath
	e.Params = m.Params
	e.Node = m.Node
	e.Sort = m.Sort
	e.Status = m.Status
	e.CreateBy = m.CreateBy
	e.IsInner = m.IsInner
	e.Values = m.Values
	e.ShowSlider = m.ShowSlider
}
