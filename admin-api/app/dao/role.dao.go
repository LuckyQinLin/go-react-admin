package dao

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/response"
	"admin-api/core"
	"admin-api/internal/gorm"
)

var Role = NewRoleDao()

type RoleDao struct {
	role     BaseDao[entity.Role]
	userRole BaseDao[entity.UserRole]
}

func NewRoleDao() *RoleDao {
	return &RoleDao{
		role:     BaseDao[entity.Role]{},
		userRole: BaseDao[entity.UserRole]{},
	}
}

// Total 查询获取总条数
func (r *RoleDao) Total(condition *gorm.DB) (total int64, err error) {
	err = condition.Model(&entity.Role{}).Count(&total).Error
	return
}

// Limit 角色获取数据
func (r *RoleDao) Limit(condition *gorm.DB, offset int, limit int) (list []response.RolePageResponse, err error) {
	err = condition.Limit(limit).Offset(offset).Model(&entity.Role{}).Find(&list).Error
	return
}

// Exist 条件查询角色信息是否存在
func (r *RoleDao) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = condition.Model(&entity.Role{}).Count(&total).Error; err != nil {
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

// GetRoleById 角色ID
func (r *RoleDao) GetRoleById(roleId int64) (role entity.Role, err error) {
	err = core.DB.Model(&entity.Role{}).
		Where("role_id = ? and del_flag = 1", roleId).
		First(&role).
		Error
	return
}

// GetRoleMenu 获取角色授权菜单数据
func (r *RoleDao) GetRoleMenu(roleId int64) (data []entity.RoleMenu, err error) {
	err = core.DB.Where("role_id = ?", roleId).Find(&data).Error
	return
}

// UpdateById 角色修改通过ID
func (r *RoleDao) UpdateById(tx *gorm.DB, update map[string]any, roleId ...int64) error {
	return tx.Model(&entity.Role{}).Where("role_id in ?", roleId).Updates(update).Error
}

// List 角色列表
func (r *RoleDao) List(condition *gorm.DB) (roles []entity.Role, err error) {
	err = condition.Model(&entity.Role{}).Find(&roles).Error
	return
}

// UserRole 获取用户角色
func (r *RoleDao) UserRole(userId int64) ([]int64, error) {
	var (
		result []int64
		data   []entity.UserRole
		err    error
	)
	if err = r.userRole.FindByMap(map[string]any{"user_id": userId}, &data); err != nil {
		return nil, err
	}
	for _, userRole := range data {
		result = append(result, userRole.RoleId)
	}
	return result, nil
}
