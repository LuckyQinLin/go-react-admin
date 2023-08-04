package controller

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/app/service"
	"admin-api/internal/gin"
	"net/http"
)

var User = new(UserController)

type UserController struct {
	BaseController
}

// CaptchaImage 获取验证码
func (u *UserController) CaptchaImage(c *gin.Context) {
	var (
		result *response.CaptchaImageResponse
		err    *response.BusinessError
	)
	if result, err = service.User.CaptchaImage(); err != nil {
		c.JSON(http.StatusOK, response.ResultCustom(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// Login 登陆
func (u *UserController) Login(c *gin.Context) {
	var (
		err       error
		customErr *response.BusinessError
		param     request.UserLoginRequest
		result    *response.UserInfoResponse
	)
	if err = c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.User.UserLogin(&param, c); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// GetUserInfo 获取用户信息
func (u *UserController) GetUserInfo(c *gin.Context) {
	var (
		value     any
		claims    *vo.UserClaims
		customErr *response.BusinessError
		result    *response.UserInfoResponse
	)
	value, _ = c.Get(vo.ClaimsInfo)
	claims = value.(*vo.UserClaims)
	if result, customErr = service.User.GetUserInfo(claims); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// AllotRole 用户分配角色
func (u *UserController) AllotRole(ctx *gin.Context) {

}

// Page 分页
func (u *UserController) Page(c *gin.Context) {
	var (
		param     request.UserPageRequest
		result    *response.PageData
		customErr *response.BusinessError
		err       error
	)
	if err = c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.User.Page(&param); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// UserCreate 用户创建
func (u *UserController) UserCreate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.UserCreateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = u.Parse(ctx, "用户创建", vo.Add, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		u.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.User.Create(&param); customErr != nil {
		u.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	u.Success(ctx, operate, response.Ok("用户创建成功"))
}

// UserUpdate 用户更新
func (u *UserController) UserUpdate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.UserUpdateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = u.Parse(ctx, "用户修改", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		u.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.User.Update(&param); customErr != nil {
		u.Failed(ctx, operate, response.ResultCustom(customErr))
	}
	u.Success(ctx, operate, response.Ok("用户修改成功"))
}
