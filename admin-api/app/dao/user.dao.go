package dao

import "admin-api/app/models/entity"

var User = new(UserDao)

type UserDao struct{}

// GetUserByUserName 获取用户信息通过用户名称
func (u *UserDao) GetUserByUserName(username string) (*entity.User, error) {
	return nil, nil
}
