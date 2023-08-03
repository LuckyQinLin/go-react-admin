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
