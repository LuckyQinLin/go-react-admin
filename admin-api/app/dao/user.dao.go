package dao

import (
	"admin-api/app/models/entity"
	"admin-api/core"
)

var User = new(UserDao)

type UserDao struct{}

// GetUserByUserName 获取用户信息通过用户名称
func (u *UserDao) GetUserByUserName(username string) (user entity.User, err error) {
	err = core.DB.Model(&entity.User{}).Where("user_name = ?", username).First(&user).Error
	return
}
