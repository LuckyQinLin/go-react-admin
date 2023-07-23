package dao

import (
	"admin-api/app/models/entity"
	"admin-api/core"
)

var Menu = new(MenuDao)

type MenuDao struct{}

// All 获取全部菜单
func (m *MenuDao) All() (menus []entity.Menu, err error) {
	err = core.DB.Model(&entity.Menu{}).Find(&menus).Error
	return
}
