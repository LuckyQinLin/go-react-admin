package dao

import (
	"admin-api/app/models/entity"
	"admin-api/core"
	"admin-api/internal/gorm"
)

var Config = new(ConfigDao)

type ConfigDao struct{}

// All 获取全部参数
func (d *ConfigDao) All() (configs []entity.Setting, err error) {
	err = core.DB.Model(&entity.Setting{}).Find(&configs).Error
	return
}

// GetById 通过参数ID获取参数数据
func (d *ConfigDao) GetById(configId int64) (config entity.Setting, err error) {
	err = core.DB.Model(&entity.Setting{}).
		Where("config_id = ?", configId).
		First(&config).
		Error
	return
}

// Exist 条件查询菜单信息是否存在
func (d *ConfigDao) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = condition.Model(&entity.Setting{}).Count(&total).Debug().Error; err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}

// Create 创建菜单
func (d *ConfigDao) Create(tx *gorm.DB, dict *entity.Setting) error {
	return tx.Create(dict).Debug().Error
}

// UpdateById 更新菜单
func (d *ConfigDao) UpdateById(tx *gorm.DB, dict *entity.Setting) error {
	return tx.Save(dict).Error
}

// Delete 删除数据
func (d *ConfigDao) Delete(tx *gorm.DB, dictId ...int64) error {
	return tx.Where("config_id in ?", dictId).Delete(&entity.Setting{}).Error
}
