package controller

import (
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/service"
	"admin-api/internal/gin"
	"net/http"
)

var Role = new(RoleController)

type RoleController struct{}

// Page 分页
func (r *RoleController) Page(c *gin.Context) {
	var (
		param     request.RolePageRequest
		result    *response.PageData
		customErr *response.BusinessError
		err       error
	)
	if err = c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.Role.Page(&param); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// RoleAll 获取所有角色
func (r *RoleController) RoleAll(c *gin.Context) {
	var (
		list      []response.RoleKeyValueResponse
		customErr *response.BusinessError
	)
	if list, customErr = service.Role.RoleAll(); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(list))
}
