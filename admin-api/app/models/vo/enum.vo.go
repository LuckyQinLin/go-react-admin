package vo

type MongoName string

const (
	UserSalt = "Lucky.麒麟"
)

const (
	AuthHeader    = "Authorization"
	HeaderUserKey = "userId"
	WSHeaderKey   = "WsValid"
)

const (
	UserKey       = "user:"
	CaptchaPrefix = "captcha:"
)

const (
	RedisToken   = "token"   // 用户token
	RedisCaptcha = "captcha" // 验证码
)
