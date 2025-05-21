package entity

type ProjectMenu struct {
	Id         int64  `gorm:"column:id;primaryKey;autoIncrement"`
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
	IsInner    int    `gorm:"column:is_inner"`
	Values     string `gorm:"column:values"`
	ShowSlider int    `gorm:"column:show_slider"`
}

func (*ProjectMenu) TableName() string {
	return "ms_project_menu"
}
