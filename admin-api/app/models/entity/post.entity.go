package entity

// Post 岗位信息表
type Post struct {
	PostId   int64  `gorm:"column:post_id;primaryKey;not null;autoIncrement;comment:岗位id" json:"postId"`
	PostCode string `gorm:"column:post_code;not null;size:64;comment:岗位编码" json:"postCode"`
	PostName string `gorm:"column:post_name;not null;size:50;comment:岗位名称" json:"postName"`
	PostSort int    `gorm:"column:post_sort;default:0;comment:显示顺序" json:"postSort"`
	Status   int    `gorm:"column:status;default:1;comment:岗位状态（1正常 0停用）" json:"status"`
	Remark   string `gorm:"size:500;default:null;comment:备注" json:"remark"`
	BaseField
}
