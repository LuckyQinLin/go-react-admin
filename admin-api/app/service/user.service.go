package service

import (
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/core"
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
	return core.Cache.Set(vo.CaptchaPrefix+id, value, time.Minute*5).Err()
}

func (r *redisStore) Get(id string, clear bool) string {
	r.Lock()
	defer r.Unlock()
	if result, err := core.Cache.Get(vo.CaptchaPrefix + id).Result(); err == nil {
		return result
	}
	if clear {
		core.Cache.Del(vo.CaptchaPrefix + id)
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
