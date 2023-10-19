package dao

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/response"
	"admin-api/core"
	"gitee.com/molonglove/goboot/gorm"
)

var Dept = new(DeptDao)

type DeptDao struct{}

// All 获取全部部门
func (d *DeptDao) All() (depts []entity.Dept, err error) {
	err = core.DB.Template(`del_flag = {{id}}`, map[string]any{"id": 1}).Find(&depts).Error
	return
}

// GetDeptById 通过部门ID获取部门数据
func (d *DeptDao) GetDeptById(deptId int64) (dept entity.Dept, err error) {
	err = core.DB.Model(&entity.Dept{}).
		Where("dept_id = ?", deptId).
		First(&dept).
		Error
	return
}

// Exist 条件查询菜单信息是否存在
func (d *DeptDao) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = condition.Model(&entity.Dept{}).Count(&total).Debug().Error; err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}

// Create 创建菜单
func (d *DeptDao) Create(tx *gorm.DB, dept *entity.Dept) error {
	return tx.Create(dept).Debug().Error
}

// UpdateById 更新菜单
func (d *DeptDao) UpdateById(tx *gorm.DB, dept *entity.Dept) error {
	return tx.Save(dept).Error
}

// Delete 删除数据
func (d *DeptDao) Delete(tx *gorm.DB, deptId ...int64) error {
	return tx.Where("dept_id in ?", deptId).Delete(&entity.Dept{}).Error
}

func (d *DeptDao) GetDeptByUserId(userId int64) (result response.UserDeptProp, err error) {
	err = core.DB.Model(entity.Dept{}).Template("dept.selectByUserId", map[string]any{"userId": userId}).First(&result).Error
	return
}
