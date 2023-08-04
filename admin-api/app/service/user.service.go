package service

import (
	"admin-api/app/dao"
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/core"
	"admin-api/internal/gin"
	"admin-api/internal/gorm"
	"admin-api/utils"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mojocn/base64Captcha"
	"github.com/mssola/useragent"
	"sync"
	"time"
)

var User = NewUserService()

type UserService struct {
	captcha *base64Captcha.Captcha
}

func NewUserService() *UserService {
	driver := base64Captcha.NewDriverDigit(40, 135, 5, 0.4, 72)
	store := &redisStore{expiration: time.Minute * 5}
	return &UserService{captcha: base64Captcha.NewCaptcha(driver, store)}
}

type redisStore struct {
	sync.RWMutex
	expiration time.Duration
}

func (r *redisStore) Set(id string, value string) error {
	r.Lock()
	defer r.Unlock()
	_, err := core.Cache.SetKeyValue(fmt.Sprintf("%s:%v", vo.RedisCaptcha, id), value, time.Minute*5)
	return err
}

func (r *redisStore) Get(id string, clear bool) string {
	r.Lock()
	defer r.Unlock()
	if result, err := core.Cache.GetKey(fmt.Sprintf("%s:%v", vo.RedisCaptcha, id)); err == nil {
		return result
	}
	if clear {
		_ = core.Cache.Delete(fmt.Sprintf("%s:%v", vo.RedisCaptcha, id))
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
		id     string
		base64 string
		err    error
	)
	if id, base64, err = u.captcha.Generate(); err != nil {
		return nil, response.NewBusinessError(response.CaptchaImageError)
	}
	// 生成验证码
	return &response.CaptchaImageResponse{
		Uuid:       id,
		Image:      base64,
		ExpireTime: time.Now().Add(time.Minute * 5).Unix(),
	}, nil
}

// UserLogin 用户登陆
func (u *UserService) UserLogin(param *request.UserLoginRequest, ctx *gin.Context) (*response.UserInfoResponse, *response.BusinessError) {
	var (
		captchaVerify = func(id, code string) error {
			if !core.Cache.Exist(vo.RedisCaptcha + ":" + id) {
				return errors.New("验证码已经过期")
			}
			if !u.captcha.Verify(id, code, true) {
				return errors.New("验证码不正确")
			}
			return nil
		} // 验证码验证
		generateToken = func(user entity.User) (claims vo.UserClaims, token string, err error) {
			claims = vo.UserClaims{
				UserId:   user.UserId,
				DeptId:   user.DeptId,
				Username: user.UserName,
				Email:    user.Email,
				Phone:    user.Phone,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(core.Config.Jwt.ExpiresTime) * time.Minute)), // 这里不配置过期时间，放到Redis中管理Token的过期时间，方便后面做续期
					IssuedAt:  jwt.NewNumericDate(time.Now()),                                                               // 签发时间
					NotBefore: jwt.NewNumericDate(time.Now()),                                                               // 生效时间
					Issuer:    core.Config.Jwt.Issuer,
				},
			}
			t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			token, err = t.SignedString([]byte(core.Config.Jwt.SecretKey))
			return
		} // 生成token信息
		getUserInfoWithVerity = func(username, password string) (user entity.User, err error) {
			if user, err = dao.User.GetUserByUserName(username); err != nil {
				err = errors.New("登录用户" + username + "不存在")
				return
			}
			// 删除
			if user.DelFlag == 0 {
				err = errors.New("对不起，您的账号：" + username + " 已被删除")
				return
			}
			// 状态
			if user.Status == 0 {
				err = errors.New("对不起，您的账号：" + username + " 已停用")
				return
			}
			// 密码
			pd, _ := utils.BcryptEncode(password)
			core.Log.Info("密码：%s", pd)
			if !utils.BcryptVerify(user.Password, password) {
				err = errors.New("账号密码不正确")
				return
			}
			return
		} // 验证用户信息
		loginLogger = func(c *gin.Context, username, msg string, status int) {
			userAgent := useragent.New(c.GetHeader("User-Agent"))
			browser, version := userAgent.Browser()
			ip := c.ClientIP()
			now := time.Now()
			_ = dao.Visit.Save(&entity.Visit{
				UserName:      username,
				IpAddr:        ip,
				LoginLocation: utils.IpAddress(ip),
				Browser:       fmt.Sprintf("%sv%s", browser, version),
				Os:            userAgent.OS(),
				Msg:           msg,
				LoginTime:     &now,
				Status:        status,
			})
		}
		token  string
		user   entity.User
		claims vo.UserClaims
		err    error
	)
	// 验证码校验
	if err = captchaVerify(param.Uuid, param.Captcha); err != nil {
		// 更新用户登录信息并记录登录信息
		go loginLogger(ctx, param.Username, err.Error(), 0)
		return nil, response.LoginBusinessError(err.Error())
	}
	// 获取用户信息
	if user, err = getUserInfoWithVerity(param.Username, param.Password); err != nil {
		// 更新用户登录信息并记录登录信息
		go loginLogger(ctx, param.Username, err.Error(), 0)
		return nil, response.LoginBusinessError(err.Error())
	}
	go loginLogger(ctx, user.UserName, "登录成功", 1)
	// 构建jwt
	if claims, token, err = generateToken(user); err != nil {
		core.Log.Error("生成认证Token错误:%s", err.Error())
		return nil, response.NewBusinessError(response.TokenBuildError)
	}
	if _, err = core.Cache.SetKeyValue(
		fmt.Sprintf("%s:%d", vo.RedisToken, user.UserId),
		token,
		time.Duration(core.Config.Jwt.ExpiresTime)*time.Minute,
	); err != nil {
		core.Log.Error("写入用户Token[%s]失败: %s", token, err.Error())
	}
	return &response.UserInfoResponse{
		Id:         user.UserId,
		UserName:   user.UserName,
		NickName:   user.NickName,
		Sex:        user.Sex,
		Avatar:     user.Avatar,
		DeptId:     user.DeptId,
		Email:      user.Email,
		Phone:      user.Phone,
		Remark:     user.Remark,
		Token:      token,
		ExpireTime: claims.ExpiresAt.Unix(),
		Roles:      nil,
	}, nil
}

// GetUserInfo 获取用户信息
func (u *UserService) GetUserInfo(claims *vo.UserClaims) (*response.UserInfoResponse, *response.BusinessError) {
	var (
		user entity.User
		err  error
	)
	if user, err = dao.User.GetUserById(claims.UserId); err != nil {
		return nil, response.NewBusinessError(response.DataNotExist)
	}

	return &response.UserInfoResponse{
		Id:         claims.UserId,
		UserName:   user.UserName,
		NickName:   user.NickName,
		Sex:        user.Sex,
		Avatar:     user.Avatar,
		DeptId:     user.DeptId,
		Email:      user.Email,
		Phone:      user.Phone,
		Remark:     user.Remark,
		ExpireTime: claims.ExpiresAt.Unix(),
		Roles:      nil,
	}, nil

}

// Page 用户分页
func (u *UserService) Page(param *request.UserPageRequest) (*response.PageData, *response.BusinessError) {
	var (
		buildCondition = func(param *request.UserPageRequest) func(db *gorm.DB) *gorm.DB {
			return func(db *gorm.DB) *gorm.DB {
				db.Model(&entity.User{}).
					Alias("u").
					Where("u.del_flag = 1")
				if param.Status != nil {
					db.Where("u.status = ?", param.Status)
				}
				if param.UserName != "" {
					db.Where("u.user_name like concat('%', ?, '%')", param.UserName)
				}
				if param.Phone != "" {
					db.Where("u.phone like concat('%', ?, '%')", param.Phone)
				}
				if param.DeptId != nil && *param.DeptId != 0 {
					db.Where("(u.dept_id = ? or u.dept_id in (select t.dept_id from sys_dept t where ?::varchar = any (string_to_array(t.ancestors, ','))))", param.DeptId)
				}
				return db
			}
		}
		list  []response.UserPageResponse
		total int64
		err   error
	)
	if err = core.DB.Scopes(buildCondition(param)).Count(&total).Debug().Error; err != nil {
		core.Log.Error("统计用户数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取用户数据失败")
	}
	if err = core.DB.Debug().Scopes(buildCondition(param)).
		Select("u.user_id,u.dept_id,u.nick_name,u.user_name,u.email,u.avatar,u.phone,u.status,u.create_time,d.dept_name").
		Joins("left join sys_dept d on u.dept_id = d.dept_id").
		Find(&list).
		Error; err != nil {
		core.Log.Error("查询用户数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取用户数据失败")
	}
	return &response.PageData{
		Total: total,
		Page:  param.Page,
		Size:  param.Size,
		Data:  list,
	}, nil
}

// Create 创建用户
func (u *UserService) Create(param *request.UserCreateRequest) *response.BusinessError {
	return nil
}

// Update 更新用户
func (u *UserService) Update(param *request.UserUpdateRequest) *response.BusinessError {
	return nil
}
