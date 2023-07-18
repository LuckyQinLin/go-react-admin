package entity

import "time"

// BaseField 基础属性
type BaseField struct {
	DelFlag    int       `gorm:"column:del_flag;default:1;comment:删除标志（1代表存在 0代表删除）" json:"delFlag"`
	CreateBy   string    `gorm:"column:create_by;default:'';size:64;comment:创建者" json:"createBy"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"createTime"`
	UpdateBy   string    `gorm:"column:update_by;default:'';size:64;comment:更新者" json:"updateBy"`
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"updateTime"`
}
