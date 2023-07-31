package dao

import (
	"admin-api/app/models/entity"
	"admin-api/core"
	"admin-api/internal/gorm"
)

var Notice = new(NoticeDao)

type NoticeDao struct{}

// All 获取全部参数
func (d *NoticeDao) All() (notices []entity.Notice, err error) {
	err = core.DB.Model(&entity.Notice{}).Find(&notices).Error
	return
}

// GetById 通过参数ID获取参数数据
func (d *NoticeDao) GetById(noticeId int64) (notice entity.Notice, err error) {
	err = core.DB.Model(&entity.Notice{}).
		Where("notice_id = ?", noticeId).
		First(&notice).
		Error
	return
}

// Exist 条件查询菜单信息是否存在
func (d *NoticeDao) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = condition.Model(&entity.Notice{}).Count(&total).Debug().Error; err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}

// Create 创建菜单
func (d *NoticeDao) Create(tx *gorm.DB, notice *entity.Notice) error {
	return tx.Create(notice).Debug().Error
}

// UpdateById 更新菜单
func (d *NoticeDao) UpdateById(tx *gorm.DB, notice *entity.Notice) error {
	return tx.Save(notice).Error
}

// Delete 删除数据
func (d *NoticeDao) Delete(tx *gorm.DB, noticeId ...int64) error {
	return tx.Where("config_id in ?", noticeId).Delete(&entity.Notice{}).Error
}
