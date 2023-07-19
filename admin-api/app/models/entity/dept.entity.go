package entity

// Dept 部门表
type Dept struct {
	DeptId    int64  `gorm:"column:dept_id;primaryKey;not null;autoIncrement;comment:部门id" json:"deptId"`
	ParentId  int64  `gorm:"column:parent_id;default:0;comment:父部门id" json:"parentId"`
	Ancestors string `gorm:"column:ancestors;default:'';size:50;comment:祖级列表" json:"ancestors"`
	DeptName  string `gorm:"column:dept_name;default:'';size:30;comment:部门名称" json:"deptName"`
	OrderNum  int    `gorm:"column:order_num;default:0;comment:显示顺序" json:"orderNum"`
	Leader    string `gorm:"column:leader;default:null;size:20;comment:负责人" json:"leader"`
	Phone     string `gorm:"column:phone;default:null;size:11;comment:联系电话" json:"phone"`
	Email     string `gorm:"column:email;default:'';size:50;comment:邮箱" json:"email"`
	Status    int    `gorm:"column:status;default:1;comment:部门状态（1正常 0停用）" json:"status"`
	DelFlag   int    `gorm:"column:del_flag;default:1;comment:删除标志（1代表存在 0代表删除）" json:"delFlag"`
	BaseField
}
