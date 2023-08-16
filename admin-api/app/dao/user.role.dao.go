package dao

import (
	"admin-api/app/models/entity"
	"admin-api/internal/gorm"
)

var UserRole = NewUserRoleDao()

type UserRoleDao struct{}

func NewUserRoleDao() *UserRoleDao {
	return &UserRoleDao{}
}

// DeleteByRoleId 删除通过角色ID
func (u *UserRoleDao) DeleteByRoleId(tx *gorm.DB, roleId int64) error {
	return tx.Delete(&entity.UserRole{}, "role_id = ?", roleId).Error
}

// InsertBatch 插入多条数据
func (u *UserRoleDao) InsertBatch(tx *gorm.DB, maps []*entity.UserRole) error {
	return tx.Create(maps).Error
}
