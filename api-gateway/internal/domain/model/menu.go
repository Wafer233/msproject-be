package model

// ProjectMenu 表示项目导航菜单项
type ProjectMenu struct {
	Id         int64
	Pid        int64
	Title      string
	Icon       string
	Url        string
	FilePath   string
	Params     string
	Node       string
	Sort       int
	Status     int
	CreateBy   int64
	IsInner    int
	Values     string
	ShowSlider int
}

// TableName 设置表名
func (*ProjectMenu) TableName() string {
	return "ms_project_menu"
}

//// ProjectMenuChild 是带有子菜单的ProjectMenu
//type ProjectMenuChild struct {
//	ProjectMenu
//	Children []*ProjectMenuChild
//}
//
//// ConvertToMenuTree 将平面菜单列表转换为树状结构
//func ConvertToMenuTree(menus []*ProjectMenu) []*ProjectMenuChild {
//	var menuChildren []*ProjectMenuChild
//
//	// 从菜单列表创建子菜单项
//	for _, menu := range menus {
//		child := &ProjectMenuChild{
//			ProjectMenu: *menu,
//			Children:    []*ProjectMenuChild{},
//		}
//		menuChildren = append(menuChildren, child)
//	}
//
//	// 找出根菜单（pid = 0）
//	var rootMenus []*ProjectMenuChild
//	for _, menu := range menuChildren {
//		if menu.Pid == 0 {
//			rootMenus = append(rootMenus, menu)
//		}
//	}
//
//	// 构建菜单树
//	buildMenuTree(rootMenus, menuChildren)
//
//	return rootMenus
//}
//
//// buildMenuTree 递归构建菜单树
//func buildMenuTree(parents []*ProjectMenuChild, allMenus []*ProjectMenuChild) {
//	for _, parent := range parents {
//		for _, menu := range allMenus {
//			if parent.Id == menu.Pid {
//				parent.Children = append(parent.Children, menu)
//			}
//		}
//
//		// 递归处理子菜单
//		if len(parent.Children) > 0 {
//			buildMenuTree(parent.Children, allMenus)
//		}
//	}
//}
