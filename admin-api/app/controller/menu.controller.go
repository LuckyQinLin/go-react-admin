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

var Menu = new(MenuController)

type MenuController struct{ BaseController }

// MenuTree 菜单树
func (m *MenuController) MenuTree(c *gin.Context) {
	var (
		tree      []*response.MenuTree
		customErr *response.BusinessError
	)
	if tree, customErr = service.Menu.Tree(); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(tree))
}

// MenuTable 菜单表格查询
func (m *MenuController) MenuTable(ctx *gin.Context) {
	var (
		err       error
		result    []*response.MenuTableResponse
		param     request.MenuTableQueryRequest
		customErr *response.BusinessError
	)
	if err = ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.Menu.Table(&param); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}

// MenuCreate 菜单创建
func (m *MenuController) MenuCreate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.MenuCreateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = m.Parse(ctx, "菜单创建", vo.Add, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		m.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Menu.Create(&param); customErr != nil {
		m.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	m.Success(ctx, operate, response.Ok("菜单创建成功"))
}

// MenuUpdate 菜单修改
func (m *MenuController) MenuUpdate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.MenuUpdateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = m.Parse(ctx, "菜单修改", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		m.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Menu.Update(&param); customErr != nil {
		m.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	m.Success(ctx, operate, response.Ok("菜单修改成功"))
}

// MenuDelete 菜单删除
func (m *MenuController) MenuDelete(ctx *gin.Context) {
	var (
		err       error
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.MenuDeleteRequest
		customErr *response.BusinessError
	)
	claims, operate = m.Parse(ctx, "菜单删除", vo.Delete, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		m.Failed(ctx, operate, response.Fail("请求参数不存在"))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Menu.Delete(&param); customErr != nil {
		m.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	m.Success(ctx, operate, response.Ok("删除菜单数据成功"))
}
