package dao

import (
	"admin-api/app/models/entity"
	"admin-api/core"
	"admin-api/internal/gorm"
)

var Post = new(PostDao)

type PostDao struct{}

// Exist 条件查询岗位信息是否存在
func (p *PostDao) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = condition.Model(&entity.Post{}).Count(&total).Error; err != nil {
		return false, err
	}
	if total > 0 {
		return true, nil
	}
	return false, nil
}

// Create 创建岗位
func (p *PostDao) Create(tx *gorm.DB, post *entity.Post) error {
	return tx.Create(post).Error
}

// UpdateById 更新菜单
func (p *PostDao) UpdateById(tx *gorm.DB, post *entity.Post) error {
	return tx.Save(post).Error
}

// Delete 删除数据
func (p *PostDao) Delete(tx *gorm.DB, postId ...int64) error {
	return tx.Where("post_id in ?", postId).Delete(&entity.Post{}).Error
}

// GetById 通过ID获取岗位数据
func (p *PostDao) GetById(postId int64) (post entity.Post, err error) {
	err = core.DB.Model(&entity.Post{}).
		Where("post_id = ?", postId).
		First(&post).
		Error
	return
}

// List 岗位列表
func (p *PostDao) List(condition *gorm.DB) (post []entity.Post, err error) {
	err = condition.Model(&entity.Post{}).Find(&post).Error
	return
}
