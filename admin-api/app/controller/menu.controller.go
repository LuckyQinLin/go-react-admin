package controller

import (
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/service"
	"admin-api/internal/gin"
	"net/http"
)

var Menu = new(MenuController)

type MenuController struct{}

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

// MenuTable 表格查询
func (m *MenuController) MenuTable(ctx *gin.Context) {
	var (
		err       error
		result    []*response.MenuTableResponse
		param     *request.MenuTableQueryRequest
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
