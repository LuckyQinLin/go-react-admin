package dao

import (
	"admin-api/app/models/entity"
	"admin-api/core"
	"admin-api/internal/gorm"
)

var Menu = new(MenuDao)

type MenuDao struct{}

// All 获取全部菜单
func (m *MenuDao) All() (menus []entity.Menu, err error) {
	err = core.DB.Model(&entity.Menu{}).Find(&menus).Error
	return
}

// GetMenuById 通过菜单ID获取菜单数据
func (m *MenuDao) GetMenuById(menuId int64) (menu entity.Menu, err error) {
	err = core.DB.Model(&entity.Menu{}).
		Where("menu_id = ?", menuId).
		First(&menu).
		Error
	return
}

// Exist 条件查询菜单信息是否存在
func (m *MenuDao) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = condition.Model(&entity.Menu{}).Count(&total).Error; err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}

// Create 创建菜单
func (m *MenuDao) Create(tx *gorm.DB, menu *entity.Menu) error {
	return tx.Create(menu).Error
}

// UpdateById 更新菜单
func (m *MenuDao) UpdateById(tx *gorm.DB, menu *entity.Menu) error {
	return tx.Save(menu).Error
}

// Delete 删除数据
func (m *MenuDao) Delete(tx *gorm.DB, menuId ...int64) error {
	return tx.Where("menu_id in ?", menuId).Delete(&entity.Menu{}).Error
}
