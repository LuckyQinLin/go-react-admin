package entity

import "time"

// Post 岗位信息表
type Post struct {
	PostId     int64      `gorm:"column:post_id;primaryKey;not null;autoIncrement;comment:岗位id" json:"postId"`
	PostCode   string     `gorm:"column:post_code;not null;size:64;comment:岗位编码" json:"postCode"`
	PostName   string     `gorm:"column:post_name;not null;size:50;comment:岗位名称" json:"postName"`
	PostSort   int        `gorm:"column:post_sort;default:0;comment:显示顺序" json:"postSort"`
	Status     int        `gorm:"column:status;default:1;comment:岗位状态（1正常 0停用）" json:"status"`
	Remark     string     `gorm:"size:500;default:null;comment:备注" json:"remark"`
	CreateBy   string     `gorm:"column:create_by;default:'';size:64;comment:创建者" json:"createBy"`
	CreateTime *time.Time `gorm:"column:create_time;default:null;comment:创建时间" json:"createTime"`
	UpdateBy   string     `gorm:"column:update_by;default:'';size:64;comment:更新者" json:"updateBy"`
	UpdateTime *time.Time `gorm:"column:update_time;default:null;comment:更新时间" json:"updateTime"`
}
