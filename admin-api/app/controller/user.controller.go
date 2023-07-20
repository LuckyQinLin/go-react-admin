package controller

import (
	"admin-api/app/models/request"
	"admin-api/app/models/response"
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
		token     string
	)
	if err = c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if token, customErr = service.User.UserLogin(&param, c); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(token))
}
