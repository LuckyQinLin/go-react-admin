package controller

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/app/service"
	"admin-api/internal/gin"
	"admin-api/utils"
	"github.com/xuri/excelize/v2"
	"net/http"
)

var Role = new(RoleController)

type RoleController struct{ BaseController }

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
	claims, operate = r.Parse(c, "角色创建", vo.Add, nil)
	if err = c.ShouldBind(&param); err != nil {
		r.Failed(c, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Role.Create(&param); customErr != nil {
		r.Failed(c, operate, response.ResultCustom(customErr))
		return
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
	claims, operate = r.Parse(ctx, "角色修改", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		r.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Role.Update(&param); customErr != nil {
		r.Failed(ctx, operate, response.ResultCustom(customErr))
	}
	r.Success(ctx, operate, response.Ok("角色修改成功"))
}

// RoleStatus 角色状态
func (r *RoleController) RoleStatus(ctx *gin.Context) {
	var (
		err       error
		claims    *vo.UserClaims
		operate   *entity.Operate
		customErr *response.BusinessError
		param     request.RoleStatusRequest
	)
	claims, operate = r.Parse(ctx, "角色状态修改", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		r.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Role.ChangeStatus(&param); customErr != nil {
		r.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	r.Success(ctx, operate, response.Ok("修改角色状态成功"))
}

// RoleInfo 角色详情
func (r *RoleController) RoleInfo(ctx *gin.Context) {
	var (
		roleId    int64
		customErr *response.BusinessError
		result    *response.RoleInfoResponse
		err       error
	)
	if roleId, err = ctx.QueryInt64("id"); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.Role.Info(roleId); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}

// RoleDelete 角色删除
func (r *RoleController) RoleDelete(ctx *gin.Context) {
	var (
		err       error
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.RoleDeleteRequest
		customErr *response.BusinessError
	)
	claims, operate = r.Parse(ctx, "角色删除", vo.Delete, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		r.Failed(ctx, operate, response.Fail("请求参数不存在"))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Role.Delete(&param); customErr != nil {
		r.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	r.Success(ctx, operate, response.Ok("删除角色数据成功"))
}

// RoleExport 角色导出
func (r *RoleController) RoleExport(ctx *gin.Context) {
	var (
		exist     bool
		ids       string
		roleIds   []int64
		err       error
		file      *excelize.File
		operate   *entity.Operate
		customErr *response.BusinessError
	)
	_, operate = r.Parse(ctx, "角色导出", vo.Other, nil)
	if ids, exist = ctx.GetQuery("ids"); !exist {
		r.Failed(ctx, operate, response.Fail("请求参数不存在"))
		return
	}
	operate.OperParam = ids
	if roleIds = utils.StrToArray(ids); len(roleIds) > 0 {
		r.Failed(ctx, operate, response.Fail("请求参数不存在"))
		return
	}
	if file, customErr = service.Role.DataExport(roleIds); customErr != nil {
		r.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	// 设置响应头
	ctx.Header("Content-Disposition", "filename=角色数据.xlsx")
	ctx.Header("Content-Type", "application/octet-stream")

	// 将文件内容写入ResponseWriter
	if err = file.Write(ctx.Writer); err != nil {
		r.Failed(ctx, operate, response.Fail("导出文件失败"))
		return
	}
}

// UserRole 用户拥有的角色
func (r *RoleController) UserRole(ctx *gin.Context) {
	var (
		customErr *response.BusinessError
		roleIds   []int64
		userId    int64
		err       error
	)
	if userId, err = ctx.QueryInt64("userId"); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if roleIds, customErr = service.Role.UserRole(userId); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(roleIds))
}

// RoleDataAuth 角色分配数据权限
func (r *RoleController) RoleDataAuth(ctx *gin.Context) {

}
