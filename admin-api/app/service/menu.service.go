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
func (m *MenuService) Tree() ([]*response.MenuTree, *response.BusinessError) {
	var (
		recursionBuild func(pId int64, data []entity.Menu) []*response.MenuTree
		tree           []*response.MenuTree
		all            []entity.Menu
		err            error
	)
	// 递归构建树
	recursionBuild = func(pId int64, data []entity.Menu) []*response.MenuTree {
		var children []*response.MenuTree
		for _, item := range data {
			if pId == item.ParentId {
				child := &response.MenuTree{
					Key:   item.MenuId,
					Label: item.MenuName,
				}
				child.Children = recursionBuild(item.MenuId, data)
				children = append(children, child)
			}
		}
		return children
	}
	if all, err = dao.Menu.All(); err != nil {
		core.Log.Error("获取菜单数据发生异常：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取菜单数据失败")
	}
	if len(all) <= 0 {
		return make([]*response.MenuTree, 0), nil
	}
	tree = recursionBuild(0, all)
	return tree, nil
}
