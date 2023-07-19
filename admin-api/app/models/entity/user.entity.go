package entity

import "time"

// User 用户表
type User struct {
	UserId    int64     `gorm:"column:user_id;primaryKey;not null;autoIncrement;comment:用户id" json:"userId"`
	DeptId    int64     `gorm:"column:dept_id;default:null;comment:部门ID" json:"deptId"`
	UserName  string    `gorm:"column:user_name;not null;size:30;comment:用户账号" json:"userName"`
	NickName  string    `gorm:"column:nick_name;not null;size:30;comment:用户昵称" json:"nickName"`
	UserType  string    `gorm:"column:user_type;default:00;size:2;comment:用户类型(00系统用户)" json:"userType"`
	Email     string    `gorm:"size:50;default:'';comment:邮箱" json:"email"`
	Phone     string    `gorm:"size:11;default:'';comment:手机" json:"phone"`
	Sex       int       `gorm:"size:1;default:0;comment:用户性别(0:女 1:男 2:未知)" json:"sex"`
	Avatar    string    `gorm:"size:100;default:'';comment:头像地址" json:"avatar"`
	Password  string    `gorm:"size:100;default:'';comment:密码" json:"password"`
	Status    int       `gorm:"size:1;default:1;comment:账号状态(1:正常 0:停用)" json:"status"`
	LoginIp   string    `gorm:"size:128;default:'';comment:最后登录IP" json:"loginIp"`
	LoginDate time.Time `gorm:"comment:最后登录时间" json:"loginDate"`
	Remark    string    `gorm:"size:500;default:null;comment:备注" json:"remark"`
	DelFlag   int       `gorm:"column:del_flag;default:1;comment:删除标志（1代表存在 0代表删除）" json:"delFlag"`
	BaseField
}

// UserRole 用户和角色关联表  用户N-1角色
type UserRole struct {
	UserId int64 `gorm:"column:user_id;primaryKey;not null;autoIncrement:false;comment:用户id" json:"userId"`
	RoleId int64 `gorm:"column:role_id;primaryKey;not null;autoIncrement:false;comment:角色id" json:"roleId"`
}

// UserPost 用户与岗位关联表  用户1-N岗位
type UserPost struct {
	UserId int64 `gorm:"column:user_id;primaryKey;not null;autoIncrement:false;comment:用户id" json:"userId"`
	PostId int64 `gorm:"column:post_id;primaryKey;not null;autoIncrement:false;comment:岗位id" json:"postId"`
}
