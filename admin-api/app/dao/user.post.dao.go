package dao

import (
	"admin-api/app/models/entity"
	"admin-api/internal/gorm"
)

var UserPost = new(UserPostDao)

type UserPostDao struct{}

// Exist 条件查询岗位信息是否存在
func (u *UserPostDao) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = condition.Model(&entity.UserPost{}).Count(&total).Error; err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}
