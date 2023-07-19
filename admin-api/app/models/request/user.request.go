package request

// UserLoginRequest 用户登陆
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"` // 账号
	Password string `json:"password" binding:"required"` // 密码
	Captcha  string `json:"captcha" binding:"required"`  // 验证码
}
