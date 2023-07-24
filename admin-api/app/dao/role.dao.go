package dao

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/response"
	"admin-api/internal/gorm"
)

var Role = new(RoleDao)

type RoleDao struct{}

// Total 查询获取总条数
func (r *RoleDao) Total(condition *gorm.DB) (total int64, err error) {
	err = condition.Model(&entity.Role{}).Debug().Count(&total).Error
	return
}

// Limit 角色获取数据
func (r *RoleDao) Limit(condition *gorm.DB, offset int, limit int) (list []response.RolePageResponse, err error) {
	err = condition.Limit(limit).Offset(offset).Model(&entity.Role{}).Debug().Find(&list).Error
	return
}

// Exist 条件查询角色信息是否存在
func (r *RoleDao) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = condition.Model(&entity.Role{}).Debug().Count(&total).Error; err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}

// Create 创建角色
func (r *RoleDao) Create(tx *gorm.DB, role *entity.Role) error {
	return tx.Create(role).Error
}

// RoleMenuMapping 保存映射关系
func (r *RoleDao) RoleMenuMapping(tx *gorm.DB, maps []*entity.RoleMenu) error {
	return tx.Create(maps).Error
}
