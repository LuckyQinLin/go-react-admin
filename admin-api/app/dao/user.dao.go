package dao

import (
	"admin-api/app/models/entity"
	"admin-api/core"
	"admin-api/internal/gorm"
)

var User = new(UserDao)

type UserDao struct{}

// GetUserByUserName 获取用户信息通过用户名称
func (u *UserDao) GetUserByUserName(username string) (user entity.User, err error) {
	err = core.DB.Model(&entity.User{}).Where("user_name = ?", username).First(&user).Error
	return
}

// GetUserById 获取用户信息
func (u *UserDao) GetUserById(id int64) (user entity.User, err error) {
	err = core.DB.Model(&entity.User{}).Where("user_id = ?", id).First(&user).Error
	return
}

// Exist 条件查询菜单信息是否存在
func (u *UserDao) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = condition.Model(&entity.User{}).Count(&total).Debug().Error; err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}
