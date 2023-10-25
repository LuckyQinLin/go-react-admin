package response

import (
	"admin-api/utils"
	"encoding/json"
	"time"
)

// CaptchaImageResponse 验证码返回
type CaptchaImageResponse struct {
	Uuid       string `json:"uuid"`       // uuid码
	Image      string `json:"image"`      // base64图片
	ExpireTime int64  `json:"expireTime"` // 到期时间
}

// UserLoginResponse 用户登录返回
type UserLoginResponse struct {
	Token      string `json:"token"`      // token信息
	ExpireTime int64  `json:"expireTime"` // 到期时间
}

// UserInfoResponse 用户信息
type UserInfoResponse struct {
	UserId   int64   `json:"userId"`   // 用户ID
	UserName string  `json:"userName"` // 用户名称
	NickName string  `json:"nickName"` // 用户昵称
	DeptId   int64   `json:"deptId"`   // 部门ID
	Phone    string  `json:"phone"`    // 手机号
	Email    string  `json:"email"`    // 邮箱
	Sex      int     `json:"sex"`      // 性别
	Status   int     `json:"status"`   // 状态
	PostId   []int64 `json:"postId"`   // 岗位
	RoleId   []int64 `json:"roleId"`   // 角色
	Remark   string  `json:"remark"`   // 备注
}

// UserLoginInfoResponse 用户登录信息
type UserLoginInfoResponse struct {
	IsSuper  bool           `json:"isSuper"`  // 是否超级管理员
	Avatar   string         `json:"avatar"`   // 头像
	UserId   int64          `json:"userId"`   // 用户ID
	UserName string         `json:"userName"` // 用户名称
	Sex      int            `json:"sex"`      // 性别
	Phone    string         `json:"phone"`    // 手机号
	NickName string         `json:"nickName"` // 昵称
	Email    string         `json:"email"`    // 邮箱
	Dept     UserDeptProp   `json:"dept"`     // 部门
	Posts    []UserPostProp `json:"posts"`    // 岗位
	Roles    []UserRoleProp `json:"roles"`    // 角色信息
	Operates []string       `json:"operates"` // 按钮等页面操作权限字符
}

// UserInfoProp 用户信息
type UserInfoProp struct {
	Admin    bool           `json:"admin"`    // 是否管理员
	Avatar   string         `json:"avatar"`   // 头像
	UserId   int64          `json:"userId"`   // 用户ID
	UserName string         `json:"userName"` // 用户名称
	Sex      int            `json:"sex"`      // 性别
	Phone    string         `json:"phone"`    // 手机号
	NickName string         `json:"nickName"` // 昵称
	Email    string         `json:"email"`    // 邮箱
	DeptId   int64          `json:"deptId"`   // 部门ID
	Dept     UserDeptProp   `json:"dept"`     // 部门
	Roles    []UserRoleProp `json:"roles"`    // 角色信息
	Posts    []UserPostProp `json:"posts"`    // 岗位
}

// UserDeptProp 用户部门
type UserDeptProp struct {
	DeptId    int64  `json:"deptId"`    // 部门ID
	ParentId  int64  `json:"parentId"`  // 上级部门ID
	DeptName  string `json:"deptName"`  // 部门名称
	Leader    string `json:"leader"`    // 部门领导
	Ancestors string `json:"ancestors"` // 部门路径
	OrderNum  int    `json:"orderNum"`  // 排序
	Status    int    `json:"status"`    // 状态
}

// UserPostProp 用户岗位
type UserPostProp struct {
	PostId   int64  `json:"postId"`   // 岗位ID
	PostName string `json:"postName"` // 岗位名称
	PostCode string `json:"postCode"` // 岗位编码
}

// UserRoleProp 用户角色
type UserRoleProp struct {
	RoleId   int64  `json:"roleId"`
	RoleName string `json:"roleName"`
	RoleCode string `json:"roleCode"`
}

// UserPageResponse 用户分页
type UserPageResponse struct {
	UserId     int64      `json:"userId"`     // 用户ID
	UserName   string     `json:"userName"`   // 用户名称
	NickName   string     `json:"nickName"`   // 昵称
	DeptName   string     `json:"deptName"`   // 部门名称
	Phone      string     `json:"phone"`      // 手机号
	Status     int        `json:"status"`     // 状态
	CreateTime *time.Time `json:"createTime"` // 创建时间
	IsSuper    bool       `json:"isSuper"`    // 是否为超级用户
}

func (s UserPageResponse) MarshalJSON() ([]byte, error) {
	type temp UserPageResponse
	return json.Marshal(&struct {
		temp
		CreateTime utils.DateTime `json:"createTime"`
	}{
		temp:       (temp)(s),
		CreateTime: utils.DateTime(*s.CreateTime),
	})
}
