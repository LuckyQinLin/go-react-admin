package controller

import (
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/app/service"
	"admin-api/internal/gin"
	"net/http"
)

var User = new(UserController)

type UserController struct{}

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
