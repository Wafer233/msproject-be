package dto

// MenuDTO 表示菜单项的数据传输对象
type MenuDTO struct {
	Id         int64      `json:"id"`
	Pid        int64      `json:"pid"`
	Title      string     `json:"title"`
	Icon       string     `json:"icon"`
	Url        string     `json:"url"`
	FilePath   string     `json:"filePath"`
	Params     string     `json:"params"`
	Node       string     `json:"node"`
	Sort       int        `json:"sort"`
	Status     int        `json:"status"`
	CreateBy   int64      `json:"createBy"`
	IsInner    int        `json:"isInner"`
	Values     string     `json:"values"`
	ShowSlider int        `json:"showSlider"`
	Children   []*MenuDTO `json:"children,omitempty"`
}

// MenuResponse 表示菜单响应
type MenuResponse struct {
	Menus []*MenuDTO `json:"menus"`
}
