package convert

import (
	"github.com/Wafer233/msproject-be/api-gateway/internal/application/dto"
	"github.com/Wafer233/msproject-be/api-gateway/internal/domain/model"
)

func ToMenuDTOs(menuTree []*model.ProjectMenuChild) []*dto.MenuDTO {
	if menuTree == nil {
		return nil
	}

	result := make([]*dto.MenuDTO, 0, len(menuTree))

	for _, menu := range menuTree {
		result = append(result, ToMenuDTO(menu))
	}

	return result
}

func ToMenuDTO(menu *model.ProjectMenuChild) *dto.MenuDTO {
	if menu == nil {
		return nil
	}

	menuDTO := &dto.MenuDTO{
		Id:         menu.Id,
		Pid:        menu.Pid,
		Title:      menu.Title,
		Icon:       menu.Icon,
		Url:        menu.Url,
		FilePath:   menu.FilePath,
		Params:     menu.Params,
		Node:       menu.Node,
		Sort:       menu.Sort,
		Status:     menu.Status,
		CreateBy:   menu.CreateBy,
		IsInner:    menu.IsInner,
		Values:     menu.Values,
		ShowSlider: menu.ShowSlider,
	}

	// 递归处理子菜单
	if len(menu.Children) > 0 {
		menuDTO.Children = ToMenuDTOs(menu.Children)
	}

	return menuDTO
}
