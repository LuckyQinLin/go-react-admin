package request

// UserLoginRequest 用户登陆
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"` // 账号
	Password string `json:"password" binding:"required"` // 密码
	Captcha  string `json:"captcha" binding:"required"`  // 验证码
	Uuid     string `json:"uuid" binding:"required"`     // 验证码ID
}

// UserPageRequest 用户分页查询
type UserPageRequest struct {
	CommonPage
	DeptId   *int64 `json:"deptId"`   // 部门
	Status   *int   `json:"status"`   // 状态
	UserName string `json:"userName"` // 用户名称
	Phone    string `json:"phone"`    // 手机号
}

// UserCreateRequest 用户创建
type UserCreateRequest struct {
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

// UserUpdateRequest 用户更新
type UserUpdateRequest struct {
	UserCreateRequest
	UserId int64 `json:"userId"`
}
