package dao

import (
	"admin-api/app/models/entity"
	"admin-api/core"
	"gitee.com/molonglove/goboot/gorm"
)

var Dict = new(DictDao)

type DictDao struct{}

// All 获取全部部门
func (d *DictDao) All() (dicts []entity.DictType, err error) {
	err = core.DB.Model(&entity.DictType{}).Find(&dicts).Error
	return
}

// GetDictById 通过部门ID获取部门数据
func (d *DictDao) GetDictById(dictId int64) (dict entity.DictType, err error) {
	err = core.DB.Model(&entity.DictType{}).
		Where("dict_id = ?", dictId).
		First(&dict).
		Error
	return
}

// Exist 条件查询菜单信息是否存在
func (d *DictDao) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = condition.Model(&entity.DictType{}).Count(&total).Debug().Error; err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}

// Create 创建菜单
func (d *DictDao) Create(tx *gorm.DB, dict *entity.DictType) error {
	return tx.Create(dict).Debug().Error
}

// UpdateById 更新菜单
func (d *DictDao) UpdateById(tx *gorm.DB, dict *entity.DictType) error {
	return tx.Save(dict).Error
}

// Delete 删除数据
func (d *DictDao) Delete(tx *gorm.DB, dictId ...int64) error {
	return tx.Where("dept_id in ?", dictId).Delete(&entity.DictType{}).Error
}
