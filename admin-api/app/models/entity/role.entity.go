package entity

import "time"

// Role 角色表
type Role struct {
	RoleId            int64      `gorm:"column:role_id;primaryKey;not null;autoIncrement;comment:角色id" json:"roleId"`
	RoleName          string     `gorm:"column:role_name;not null;size:64;comment:角色名称" json:"roleName"`
	RoleKey           string     `gorm:"column:role_key;not null;size:50;comment:角色权限字符串" json:"roleKey"`
	RoleSort          int        `gorm:"column:role_sort;default:0;comment:显示顺序" json:"roleSort"`
	DataScope         int        `gorm:"column:data_scope;default:1;comment:数据范围(1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限)" json:"dataScope"`
	MenuCheckStrictly bool       `gorm:"column:menu_check_strictly;type:boolean;default:true;comment:菜单树选择项是否关联显示" json:"menuCheckStrictly"`
	DeptCheckStrictly bool       `gorm:"column:dept_check_strictly;type:boolean;default:true;comment:部门树选择项是否关联显示" json:"deptCheckStrictly"`
	Status            int        `gorm:"column:status;default:1;comment:角色状态（1正常 0停用）" json:"status"`
	Remark            string     `gorm:"size:500;default:'';comment:备注" json:"remark"`
	DelFlag           int        `gorm:"column:del_flag;default:1;comment:删除标志（1代表存在 0代表删除）" json:"delFlag"`
	CreateBy          string     `gorm:"column:create_by;default:'';size:64;comment:创建者" json:"createBy"`
	CreateTime        *time.Time `gorm:"column:create_time;default:null;comment:创建时间" json:"createTime"`
	UpdateBy          string     `gorm:"column:update_by;default:'';size:64;comment:更新者" json:"updateBy"`
	UpdateTime        *time.Time `gorm:"column:update_time;default:null;comment:更新时间" json:"updateTime"`
}

// RoleDept 角色和部门关联表  角色1-N部门
type RoleDept struct {
	RoleId int64 `gorm:"column:role_id;primaryKey;not null;autoIncrement:false;comment:角色id" json:"roleId"`
	DeptId int64 `gorm:"column:dept_id;primaryKey;not null;autoIncrement:false;comment:部门id" json:"deptId"`
}

// RoleMenu 角色和菜单关联表  角色1-N菜单
type RoleMenu struct {
	RoleId int64 `gorm:"column:role_id;primaryKey;not null;autoIncrement:false;comment:角色id" json:"roleId"`
	MenuId int64 `gorm:"column:menu_id;primaryKey;not null;autoIncrement:false;comment:菜单id" json:"menuId"`
}
