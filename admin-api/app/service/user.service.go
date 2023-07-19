package service

import (
	"admin-api/app/dao"
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/core"
	"admin-api/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mojocn/base64Captcha"
	"sync"
	"time"
)

var User = new(UserService)

type UserService struct{}

type redisStore struct {
	sync.RWMutex
	expiration time.Duration
}

func (r *redisStore) Set(id string, value string) error {
	r.Lock()
	defer r.Unlock()
	_, err := core.Cache.SetKeyValue(vo.CaptchaPrefix, id, value, time.Minute*5)
	return err
}

func (r *redisStore) Get(id string, clear bool) string {
	r.Lock()
	defer r.Unlock()
	if result, err := core.Cache.GetKey(vo.CaptchaPrefix, id); err == nil {
		return result
	}
	if clear {
		_ = core.Cache.Delete(vo.CaptchaPrefix, id)
	}
	return ""
}

func (r *redisStore) Verify(id, answer string, clear bool) bool {
	return r.Get(id, clear) == answer
}

// CaptchaImage 获取验证码
// https://mojotv.cn/go/refactor-base64-captcha 验证码文档地址
func (u *UserService) CaptchaImage() (*response.CaptchaImageResponse, *response.BusinessError) {
	var (
		captcha *base64Captcha.Captcha
		driver  base64Captcha.Driver
		store   base64Captcha.Store
		base64  string
		err     error
	)
	driver = base64Captcha.NewDriverDigit(40, 135, 5, 0.4, 72)
	store = &redisStore{expiration: time.Minute * 5}
	captcha = base64Captcha.NewCaptcha(driver, store)
	if _, base64, err = captcha.Generate(); err != nil {
		return nil, response.NewBusinessError(response.CaptchaImageError)
	}
	// 生成验证码
	return &response.CaptchaImageResponse{
		Uuid:       "",
		Image:      base64,
		ExpireTime: time.Now().Add(time.Minute * 5).Unix(),
	}, nil
}

// UserLogin 用户登陆
func (u *UserService) UserLogin(param *request.UserLoginRequest) (string, *response.BusinessError) {
	var (
		generateToken = func(user *entity.User) (string, error) {
			claims := vo.UserClaims{
				UserId:   user.UserId,
				Username: user.UserName,
				Email:    user.Email,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(core.Config.Jwt.ExpiresTime) * time.Hour)), // 过期时间
					IssuedAt:  jwt.NewNumericDate(time.Now()),                                                             // 签发时间
					NotBefore: jwt.NewNumericDate(time.Now()),                                                             // 生效时间
					Issuer:    core.Config.Jwt.Issuer,
				},
			}
			t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			return t.SignedString([]byte(core.Config.Jwt.SecretKey))
		} // 生成token信息
		token string
		user  *entity.User
		err   error
	)
	// 获取用户信息
	if user, err = dao.User.GetUserByUserName(param.Username); err != nil {
		core.Log.Error("当前用户[%s]不存在: [%s]", param.Username, err.Error())
		return "", response.NewBusinessError(response.DataNotExist)
	}
	// 密码不正确
	if utils.TransformMd5(param.Password+vo.UserSalt) != user.Password {
		core.Log.Error("当前用户密码[%s]不正确", param.Password)
		return "", response.NewBusinessError(response.UserPasswordError)
	}
	// 构建jwt
	if token, err = generateToken(user); err != nil {
		core.Log.Error("生成认证Token错误:%s", err.Error())
		return "", response.NewBusinessError(response.TokenBuildError)
	}
	_, err = core.Cache.SetKeyValue(vo.UserKey, user.UserId, token, time.Duration(core.Config.Jwt.ExpiresTime)*time.Hour)
	return token, nil
}
