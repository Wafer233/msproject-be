package model

import "github.com/jinzhu/copier"

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

type ProjectMenuChild struct {
	ProjectMenu
	StatusText string
	InnerText  string
	FullUrl    string
	Children   []*ProjectMenuChild
}

func CovertChild(pms []*ProjectMenu) []*ProjectMenuChild {
	var pmcs []*ProjectMenuChild
	copier.Copy(&pmcs, pms)

	// 添加额外字段
	for _, v := range pmcs {
		v.StatusText = GetStatus(v.Status)
		v.InnerText = GetInnerText(v.IsInner)
		v.FullUrl = GetFullUrl(v.Url, v.Params, v.Values)
	}

	// 提取顶层菜单
	var childPmcs []*ProjectMenuChild
	for _, v := range pmcs {
		if v.Pid == 0 {
			childPmcs = append(childPmcs, v)
		}
	}

	// 构建子树
	ToChild(childPmcs, pmcs)
	return childPmcs
}

func ToChild(childPmcs []*ProjectMenuChild, pmcs []*ProjectMenuChild) {
	for _, pmc := range childPmcs {
		for _, pm := range pmcs {
			if pmc.Id == pm.Pid {
				child := &ProjectMenuChild{}
				copier.Copy(child, pm)
				pmc.Children = append(pmc.Children, child)
			}
		}
		ToChild(pmc.Children, pmcs)
	}
}
func GetFullUrl(url string, params string, values string) string {
	if (params != "" && values != "") || values != "" {
		return url + "/" + values
	}
	return url
}

func GetInnerText(inner int) string {
	if inner == 0 {
		return "导航"

	}
	if inner == 1 {
		return "内页"
	}
	return ""
}

func GetStatus(status int) string {
	if status == 0 {
		return "禁用"
	}
	if status == 1 {
		return "使用中"
	}
	return ""
}
