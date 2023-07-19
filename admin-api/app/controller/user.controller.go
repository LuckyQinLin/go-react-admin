package controller

import (
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
