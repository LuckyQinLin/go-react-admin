package controller

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/app/service"
	"gitee.com/molonglove/goboot/gin"
	"net/http"
)

var Dept = new(DeptController)

type DeptController struct{ BaseController }

// DeptCreate 部门创建
func (d *DeptController) DeptCreate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.DeptCreateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = d.Parse(ctx, "部门创建", vo.Add, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		d.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Dept.Create(&param); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("部门创建成功"))
}

// DeptUpdate 部门更新
func (d *DeptController) DeptUpdate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.DeptUpdateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = d.Parse(ctx, "部门修改", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		d.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Dept.Update(&param); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("部门修改成功"))
}

// DeptDelete 部门删除
func (d *DeptController) DeptDelete(ctx *gin.Context) {
	var (
		err       error
		claims    *vo.UserClaims
		operate   *entity.Operate
		deptId    int64
		customErr *response.BusinessError
	)
	claims, operate = d.Parse(ctx, "部门删除", vo.Delete, nil)
	if deptId, err = ctx.QueryInt64("deptId"); err != nil {
		d.Failed(ctx, operate, response.Fail("请求参数不存在"))
		return
	}
	operate.ParamToJson(deptId)
	if customErr = service.Dept.Delete(deptId, claims.Username); customErr != nil {
		d.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	d.Success(ctx, operate, response.Ok("删除部门数据成功"))
}

// DeptTableTree 部门表格树
func (d *DeptController) DeptTableTree(ctx *gin.Context) {
	var (
		err       error
		result    []*response.DeptTableResponse
		param     request.DeptTableQueryRequest
		customErr *response.BusinessError
	)
	if err = ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.Dept.Table(&param); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}

// DeptTree 部门树
func (d *DeptController) DeptTree(ctx *gin.Context) {
	var (
		tree      []*response.DeptTree
		customErr *response.BusinessError
	)
	if tree, customErr = service.Dept.Tree(); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(tree))
}

// DeptInfo 部门详情
func (d *DeptController) DeptInfo(ctx *gin.Context) {
	var (
		err       error
		deptId    int64
		info      *response.DeptInfoResponse
		customErr *response.BusinessError
	)
	if deptId, err = ctx.QueryInt64("deptId"); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if info, customErr = service.Dept.Info(deptId); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(info))
}
