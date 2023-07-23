package response

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

// UserInfoResponse 用户信息
type UserInfoResponse struct {
	Id         int64          `json:"id"`         // 用户ID
	UserName   string         `json:"userName"`   // 姓名
	NickName   string         `json:"nickName"`   // 昵称
	Sex        int            `json:"sex"`        // 性别
	Avatar     string         `json:"avatar"`     // 头像
	DeptId     int64          `json:"deptId"`     // 部门ID
	Email      string         `json:"email"`      // 邮箱
	Phone      string         `json:"phone"`      // 手机号
	Remark     string         `json:"remark"`     // 备注
	Token      string         `json:"token"`      // token信息
	ExpireTime int64          `json:"expireTime"` // 到期时间
	Roles      []UserRoleProp `json:"roles"`      // 用户角色
}
