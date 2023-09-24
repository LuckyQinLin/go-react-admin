package dao

import (
	"admin-api/app/models/entity"
	"gitee.com/molonglove/goboot/gorm"
)

var RoleMenu = new(RoleMenuDao)

type RoleMenuDao struct{}

// Delete 删除数据
func (r *RoleMenuDao) Delete(tx *gorm.DB, roleIds ...int64) error {
	return tx.Where("role_id in ?", roleIds).Delete(&entity.RoleMenu{}).Error
}

// Exist 条件查询菜单信息是否存在
func (r *RoleMenuDao) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = condition.Model(&entity.RoleMenu{}).Count(&total).Error; err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}
