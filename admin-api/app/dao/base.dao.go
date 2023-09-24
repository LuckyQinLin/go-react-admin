package dao

import (
	"admin-api/app/models/entity"
	"admin-api/core"
	"gitee.com/molonglove/goboot/gorm"
)

type BaseRestraint interface {
	entity.Dept | entity.User | entity.Role | entity.UserRole
}

type BaseDao[T BaseRestraint] struct {
	model T
}

// Update 更新
func (b *BaseDao[T]) Update(condition *gorm.DB, data map[string]any) error {
	return condition.Model(&b.model).Updates(data).Error
}

// TUpdate 更新
func (b *BaseDao[T]) TUpdate(tx *gorm.DB, condition *gorm.DB, data map[string]any) error {
	return tx.Model(&b.model).Where(condition).Updates(data).Error
}

// Find 条件查询
func (b *BaseDao[T]) Find(condition *gorm.DB) (data []any, err error) {
	err = condition.Model(&b.model).Find(data).Error
	return
}

// FindByMap 条件查询
func (b *BaseDao[T]) FindByMap(condition map[string]any, data any) error {
	return core.DB.Model(&b.model).Where(condition).Find(data).Error
}

// Limit 分页查询
func (b *BaseDao[T]) Limit(condition *gorm.DB, offset, limit int) (data []any, err error) {
	err = condition.Limit(limit).Offset(offset).Model(&b.model).Find(data).Error
	return
}

// CreateMany 创建多条
func (b *BaseDao[T]) CreateMany(data []T) error {
	return core.DB.Model(&b.model).Create(data).Error
}

// TCreateMany 带事务创建多条
func (b *BaseDao[T]) TCreateMany(tx *gorm.DB, data []T) error {
	return tx.Model(&b.model).Create(data).Error
}

// Create 创建
func (b *BaseDao[T]) Create(data T) error {
	return core.DB.Model(&b.model).Create(data).Error
}

// TCreate 创建(事务)
func (b *BaseDao[T]) TCreate(tx *gorm.DB, data T) error {
	return tx.Model(&b.model).Create(data).Error
}

// DeleteById 删除通过ID
func (b *BaseDao[T]) DeleteById(id any) error {
	return core.DB.Delete(&b.model, id).Error
}

// TDeleteById 删除通过ID
func (b *BaseDao[T]) TDeleteById(tx *gorm.DB, id any) error {
	return tx.Delete(&b.model, id).Error
}

// ConditionDelete 条件删除
func (b *BaseDao[T]) ConditionDelete(condition *gorm.DB) error {
	return condition.Delete(&b.model).Error
}

// TConditionDelete 条件删除
func (b *BaseDao[T]) TConditionDelete(tx *gorm.DB, condition *gorm.DB) error {
	return tx.Where(condition).Delete(&b.model).Error
}

// GetById 通过ID获取数据
func (b *BaseDao[T]) GetById(id any) (model T, err error) {
	err = core.DB.First(&model, id).Error
	return
}

// GetOne 获取一个
func (b *BaseDao[T]) GetOne(condition *gorm.DB) (data T, err error) {
	err = condition.Model(&b.model).First(&data).Error
	return
}

// Count 统计数量
func (b *BaseDao[T]) Count(condition *gorm.DB) (num int64, err error) {
	err = condition.Model(&b.model).Count(&num).Error
	return
}

// Exist 存在
func (b *BaseDao[T]) Exist(condition *gorm.DB) (bool, error) {
	var (
		total int64
		err   error
	)
	if err = core.DB.Model(&b.model).Where(condition).Count(&total).Debug().Error; err != nil || total <= 0 {
		return false, err
	}
	return true, nil
}

// Page 分页查询
func (b *BaseDao[T]) Page(handler func(*gorm.DB) *gorm.DB) (data []any, total int64, err error) {
	if err = core.DB.Scopes(handler).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return []any{}, 0, nil
	}
	if err = core.DB.Scopes(handler).Find(&data).Error; err != nil {
		return nil, total, err
	}
	return
}
