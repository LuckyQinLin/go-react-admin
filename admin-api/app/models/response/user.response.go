package response

// CaptchaImageResponse 验证码返回
type CaptchaImageResponse struct {
	Uuid       string `json:"uuid"`       // uuid码
	Image      string `json:"image"`      // base64图片
	ExpireTime int64  `json:"expireTime"` // 到期时间
}
