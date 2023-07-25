package dao

import (
	"admin-api/app/models/entity"
	"admin-api/internal/gorm"
)

var RoleMenu = new(RoleMenuDao)

type RoleMenuDao struct{}

// Delete 删除数据
func (r *RoleMenuDao) Delete(tx *gorm.DB, roleIds ...int64) error {
	return tx.Where("role_id in ?", roleIds).Delete(&entity.RoleMenu{}).Error
}
