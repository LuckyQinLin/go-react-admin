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
	"database/sql"
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
func (u *UserService) UserLogin(param *request.UserLoginRequest, ctx *gin.Context) (*response.UserLoginResponse, *response.BusinessError) {
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
	return &response.UserLoginResponse{
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
func (u *UserService) GetUserInfo(userId int64) (*response.UserInfoResponse, *response.BusinessError) {
	var (
		postId []int64
		roleId []int64
		user   entity.User
		result *response.UserInfoResponse
		err    error
	)
	if user, err = dao.User.GetUserById(userId); err != nil {
		return nil, response.NewBusinessError(response.DataNotExist)
	}
	result = &response.UserInfoResponse{
		UserId:   user.UserId,
		UserName: user.UserName,
		NickName: user.NickName,
		DeptId:   user.DeptId,
		Email:    user.Email,
		Phone:    user.Phone,
		Remark:   user.Remark,
		Status:   user.Status,
		Sex:      user.Sex,
	}
	// 用户岗位
	if postId, err = dao.User.UserPostId(user.UserId); err == nil && len(postId) > 0 {
		result.PostId = postId
	}
	// 用户角色
	if roleId, err = dao.User.UserRoleId(user.UserId); err == nil && len(roleId) > 0 {
		result.RoleId = roleId
	}
	return result, nil
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
					db.Where("u.status = @status", sql.Named("status", param.Status))
				}
				if param.UserName != "" {
					db.Where("u.user_name like concat('%', @userName, '%')", sql.Named("userName", param.UserName))
				}
				if param.Phone != "" {
					db.Where("u.phone like concat('%', @phone, '%')", sql.Named("phone", param.Phone))
				}
				if param.DeptId != nil && *param.DeptId != 0 {
					db.Where("(u.dept_id = @deptId or u.dept_id in (select t.dept_id from sys_dept t where @deptId = any (string_to_array(t.ancestors, ',')::integer[])))", sql.Named("deptId", param.DeptId))
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
	var (
		userPosts []*entity.UserPost
		userRoles []*entity.UserRole
		user      entity.User
		condition *gorm.DB
		now       time.Time
		encode    string
		err       error
		isExist   bool
	)
	// 判断是否存在相同账号
	condition = core.DB.Where("user_name = ? and del_flag = 1", param.UserName)
	if isExist, err = dao.User.Exist(condition); err != nil || isExist {
		core.Log.Error("新增用户'%s'失败，登录账号已存在", param.UserName)
		return response.CustomBusinessError(response.Failed, "新增用户'"+param.UserName+"'失败，登录账号已存在")
	}
	// 判断是否存在相同手机
	condition = core.DB.Where("phone = ? and del_flag = 1", param.Phone)
	if isExist, err = dao.User.Exist(condition); err != nil || isExist {
		core.Log.Error("新增用户'%s'失败，手机账号已存在", param.UserName)
		return response.CustomBusinessError(response.Failed, "新增用户'"+param.UserName+"'失败，手机账号已存在")
	}
	// 判断是否存在相同的邮箱
	condition = core.DB.Where("email = ? and del_flag = 1", param.Email)
	if isExist, err = dao.User.Exist(condition); err != nil || isExist {
		core.Log.Error("新增用户'%s'失败，邮箱账号已存在", param.UserName)
		return response.CustomBusinessError(response.Failed, "新增用户'"+param.UserName+"'失败，邮箱账号已存在")
	}
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		now = time.Now()
		encode, _ = utils.BcryptEncode(param.Password)
		// 创建角色
		user = entity.User{
			UserName:   param.UserName,
			NickName:   param.NickName,
			Password:   encode,
			Email:      param.Email,
			Phone:      param.Phone,
			Sex:        param.Sex,
			DeptId:     param.DeptId,
			Status:     param.Status,
			CreateBy:   param.CreateName,
			Remark:     param.Remark,
			DelFlag:    1,
			CreateTime: &now,
		}
		if err = dao.User.Create(tx, &user); err != nil {
			core.Log.Error("创建用户[%s]失败：%s", param.UserName, err.Error())
			return err
		}
		// 创建角色菜单映射关系
		if len(param.PostId) > 0 {
			for _, id := range param.PostId {
				userPosts = append(userPosts, &entity.UserPost{UserId: user.UserId, PostId: id})
			}
			if err = tx.Model(&entity.UserPost{}).Create(userPosts).Error; err != nil {
				core.Log.Error("创建用户[%s]和岗位[%v]映射关系失败：%s", user.UserName, param.PostId, err.Error())
				return err
			}
		}
		if len(param.RoleId) > 0 {
			for _, id := range param.RoleId {
				userRoles = append(userRoles, &entity.UserRole{UserId: user.UserId, RoleId: id})
			}
			if err = tx.Model(&entity.UserRole{}).Create(userRoles).Error; err != nil {
				core.Log.Error("创建用户[%s]和角色[%v]映射关系失败：%s", user.UserName, param.RoleId, err.Error())
				return err
			}
		}
		return nil
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "创建角色失败")
	}
	core.Log.Info("创建用户[%d:%s]成功", user.UserId, user.UserName)
	return nil
}

// Update 更新用户
func (u *UserService) Update(param *request.UserUpdateRequest) *response.BusinessError {
	var (
		err       error
		isNeed    bool
		posts     []*entity.UserPost
		roles     []*entity.UserRole
		now       time.Time
		old       entity.User
		customErr *response.BusinessError
		contrast  = func(old *entity.User, param *request.UserUpdateRequest) (isNeed bool, bErr *response.BusinessError) {
			var (
				condition *gorm.DB
				exist     bool
				err       error
			)
			isNeed = false
			if old.UserName != param.UserName {
				// 判断新的角色名称是否存在相同的角色信息
				condition = core.DB.Where("user_name = ? and del_flag = 1 and user_id != ?", param.UserName, old.UserId)
				if exist, err = dao.User.Exist(condition); err != nil || exist {
					return false, response.CustomBusinessError(response.Failed, "存在相同的账号名称["+param.UserName+"]")
				}
				old.UserName = param.UserName
				isNeed = true
			}
			if old.Phone != param.Phone {
				condition = core.DB.Where("phone = ? and del_flag = 1 and user_id != ?", param.Phone, old.UserId)
				if exist, err = dao.User.Exist(condition); err != nil || exist {
					return false, response.CustomBusinessError(response.Failed, "存在相同的手机号["+param.Phone+"]")
				}
				old.Phone = param.Phone
				isNeed = true
			}
			if old.Email != param.Email {
				condition = core.DB.Where("email = ? and del_flag = 1 and user_id != ?", param.Email, old.UserId)
				if exist, err = dao.User.Exist(condition); err != nil || exist {
					return false, response.CustomBusinessError(response.Failed, "存在相同的邮箱["+param.Phone+"]")
				}
				old.Email = param.Email
				isNeed = true
			}
			if old.DeptId != param.DeptId {
				old.DeptId = param.DeptId
				isNeed = true
			}
			if old.Sex != param.Sex {
				old.Status = param.Status
				isNeed = true
			}
			if old.Status != param.Status {
				old.Status = param.Status
				isNeed = true
			}
			if old.NickName != param.NickName {
				old.NickName = param.NickName
				isNeed = true
			}
			if old.Remark != param.Remark {
				old.Remark = param.Remark
				isNeed = true
			}
			return isNeed, nil
		} // 对比是否需要更新数据
	)
	// 获取修改数据
	if old, err = dao.User.GetUserById(param.UserId); err != nil {
		core.Log.Error("当前用户[%d]不存在", param.RoleId)
		return response.CustomBusinessError(response.Failed, "当前用户不存在")
	}
	// 判断是否需要修改数据
	if isNeed, customErr = contrast(&old, param); customErr != nil || !isNeed {
		core.Log.Error("修改用户失败：%s", customErr.Error())
		return customErr
	}
	// 判断是否需要更新用户和岗位
	if len(param.PostId) > 0 {
		for _, id := range param.PostId {
			posts = append(posts, &entity.UserPost{UserId: param.UserId, PostId: id})
		}
	}
	// 判断是否需要更新用户和角色
	if len(param.RoleId) > 0 {
		for _, id := range param.RoleId {
			roles = append(roles, &entity.UserRole{UserId: param.UserId, RoleId: id})
		}
	}
	// 执行更新
	if err = core.DB.Transaction(func(tx *gorm.DB) (err error) {
		if isNeed {
			now = time.Now()
			old.UpdateBy = param.UserName
			old.UpdateTime = &now
			if err = tx.Save(&old).Error; err != nil {
				core.Log.Error("更新用户数据失败:%s", err.Error())
				return
			}
		}
		// 判断是否需要更新映射关系
		if len(roles) > 0 {
			if err = tx.Where("user_id = ?", old.UserId).Delete(&entity.UserRole{}).Error; err != nil {
				core.Log.Error("删除用户和角色映射数据失败:%s", err.Error())
				return
			}
			if err = tx.Save(&roles).Error; err != nil {
				core.Log.Error("创建用户[%s]和角色映射关系失败：%s", param.UserId, err.Error())
				return
			}
		}
		if len(posts) > 0 {
			if err = tx.Where("user_id = ?", old.UserId).Delete(&entity.UserPost{}).Error; err != nil {
				core.Log.Error("删除用户和岗位映射数据失败:%s", err.Error())
				return
			}
			if err = tx.Save(&posts).Error; err != nil {
				core.Log.Error("创建用户[%s]和岗位映射关系失败：%s", param.UserId, err.Error())
				return
			}
		}
		return
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "更新用户失败")
	}
	return nil
}

// ResetPassword 重置用户密码
func (u *UserService) ResetPassword(param *request.UserPasswordRequest) *response.BusinessError {
	var (
		now      time.Time
		password string
		user     entity.User
		err      error
	)
	if user, err = dao.User.GetUserById(param.UserId); err != nil {
		core.Log.Error("当前用户[%d]不存在", param.UserId)
		return response.CustomBusinessError(response.Failed, "当前用户不存在")
	}
	now = time.Now()
	password, _ = utils.BcryptEncode(param.Password)
	user.Password = password
	user.UpdateBy = param.UpdateName
	user.UpdateTime = &now
	if err = core.DB.Model(&entity.User{}).Save(&user).Error; err != nil {
		return response.CustomBusinessError(response.Failed, "重置密码失败")
	}
	return nil
}

// ChangeStatus 修改用户状态
func (u *UserService) ChangeStatus(param *request.UserStatusRequest) *response.BusinessError {
	var (
		now   time.Time
		exist bool
		err   error
	)
	if exist, err = dao.User.Exist(core.DB.Where("user_id = ?", param.UserId)); err != nil || !exist {
		core.Log.Error("当前用户[%d]不存在", param.UserId)
		return response.CustomBusinessError(response.Failed, "当前用户不存在")
	}
	now = time.Now()
	if err = dao.User.Update(
		core.DB.Where("user_id = ?", param.UserId),
		map[string]any{"update_by": param.UpdateName, "update_time": &now, "status": param.Status},
	); err != nil {
		return response.CustomBusinessError(response.Failed, "重置密码失败")
	}
	return nil
}

// DeleteUser 删除用户
func (u *UserService) DeleteUser(param *request.UserDeleteRequest) *response.BusinessError {
	var (
		now time.Time
		err error
	)
	now = time.Now()
	if err = dao.User.Update(
		core.DB.Where("user_id in ?", param.Ids),
		map[string]any{"update_by": param.UpdateName, "update_time": &now, "del_flag": 0},
	); err != nil {
		return response.CustomBusinessError(response.Failed, "删除用户成功")
	}
	return nil
}
