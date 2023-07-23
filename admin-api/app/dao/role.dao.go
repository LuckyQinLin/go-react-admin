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
