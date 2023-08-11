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

// UserRoleProp 用户角色属性
type UserRoleProp struct {
	RoleId   int64  `json:"roleId"`
	RoleName string `json:"roleName"`
	RoleCode string `json:"roleCode"`
}

// UserLoginResponse 用户登录返回
type UserLoginResponse struct {
	Id         int64  `json:"id"`         // 用户ID
	UserName   string `json:"userName"`   // 姓名
	NickName   string `json:"nickName"`   // 昵称
	Sex        int    `json:"sex"`        // 性别
	Avatar     string `json:"avatar"`     // 头像
	DeptId     int64  `json:"deptId"`     // 部门ID
	Email      string `json:"email"`      // 邮箱
	Phone      string `json:"phone"`      // 手机号
	Remark     string `json:"remark"`     // 备注
	Token      string `json:"token"`      // token信息
	ExpireTime int64  `json:"expireTime"` // 到期时间

}

// UserInfoResponse 用户信息
type UserInfoResponse struct {
	Menus      []any    `json:"menus"`      // 菜单
	Roles      []string `json:"roles"`      // 角色字符
	Permission []string `json:"permission"` // 权限字符
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
