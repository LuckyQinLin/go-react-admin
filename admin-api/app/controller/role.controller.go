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

var Role = new(RoleController)

type RoleController struct {
	BaseController
}

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

// RoleCreate 角色创建
func (r *RoleController) RoleCreate(c *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.RoleCreateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = r.Parse(c, "角色创建", vo.Add, param)
	if err = c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	if customErr = service.Role.Create(&param); customErr != nil {
		r.Failed(c, operate, response.ResultCustom(customErr))
	}
	r.Success(c, operate, response.Ok("角色创建成功"))
}

// RoleUpdate 角色修改
func (r *RoleController) RoleUpdate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.RoleUpdateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = r.Parse(ctx, "角色修改", vo.Update, param)
	if err = ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	if customErr = service.Role.Update(&param); customErr != nil {
		r.Failed(ctx, operate, response.ResultCustom(customErr))
	}
	r.Success(ctx, operate, response.Ok("角色修改成功"))
}
