package service

import (
	"admin-api/app/dao"
	"admin-api/app/models/entity"
	"admin-api/app/models/response"
	"admin-api/core"
)

var Menu = new(MenuService)

type MenuService struct{}

// Tree 菜单树
func (m *MenuService) Tree() ([]response.MenuTree, *response.BusinessError) {
	var (
		recursionBuild = func(parent *response.MenuTree, []entity.Menu) {}
		all            []entity.Menu
		err            error
	)
	if all, err = dao.Menu.All(); err != nil {
		core.Log.Error("获取菜单数据发生异常：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取菜单数据失败")
	}
	if len(all) <= 0 {
		return []response.MenuTree{}, nil
	}

	return nil, nil
}
