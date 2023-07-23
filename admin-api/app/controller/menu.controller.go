package controller

import (
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
		tree      []response.MenuTree
		customErr *response.BusinessError
	)
	if tree, customErr = service.Menu.Tree(); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(tree))
}
