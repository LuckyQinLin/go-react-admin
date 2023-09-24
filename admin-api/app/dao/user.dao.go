package dao

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/response"
	"admin-api/core"
	"gitee.com/molonglove/goboot/gorm"
)

var User = NewUserDao()

type UserDao struct {
	BaseDao[entity.User]
}

func NewUserDao() *UserDao {
	return &UserDao{struct {
		model entity.User
	}{
		model: entity.User{},
	},
	}
}

// Create 创建用户
func (u *UserDao) Create(tx *gorm.DB, user *entity.User) error {
	return tx.Create(user).Error
}

// GetUserByUserName 获取用户信息通过用户名称
func (u *UserDao) GetUserByUserName(username string) (user entity.User, err error) {
	return u.GetOne(core.DB.Where("user_name = ?", username))
	//err = core.DB.Model(&entity.User{}).Where("user_name = ?", username).First(&user).Error
	//return
}

// GetUserById 获取用户信息
func (u *UserDao) GetUserById(id int64) (user entity.User, err error) {
	return u.GetById(id)
}

// Total 查询获取总条数
func (u *UserDao) Total(condition *gorm.DB) (total int64, err error) {
	err = condition.Model(&entity.User{}).Count(&total).Error
	return
}

// Limit 用户获取数据
func (u *UserDao) Limit(condition *gorm.DB, offset int, limit int) (list []response.RolePageResponse, err error) {
	err = condition.Limit(limit).Offset(offset).Model(&entity.User{}).Find(&list).Error
	return
}

// UserPostId 用户岗位
func (u *UserDao) UserPostId(userId int64) (data []int64, err error) {
	err = core.DB.Model(&entity.UserPost{}).
		Where("user_id = ?", userId).
		Distinct().
		Pluck("post_id", &data).
		Error
	return
}

// UserRoleId 用户角色
func (u *UserDao) UserRoleId(userId int64) (data []int64, err error) {
	err = core.DB.Model(&entity.UserRole{}).
		Where("user_id = ?", userId).
		Distinct().
		Pluck("role_id", &data).
		Error
	return
}
