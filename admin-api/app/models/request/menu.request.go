package request

// MenuTableQueryRequest 菜单表格查询
type MenuTableQueryRequest struct {
	Name   string `json:"name"`   // 菜单名称
	Status int    `json:"status"` // 菜单状态
}

// MenuCreateRequest 菜单创建
type MenuCreateRequest struct {
	ParentId int64  `json:"parentId" binding:"required"` // 上级菜单
	MenuType int    `json:"menuType"`                    // 菜单类型
	Icon     string `json:"icon"`                        // 图标
	MenuName string `json:"menuName"`                    // 菜单名称
	MenuSort int    `json:"menuSort"`                    // 显示顺序
	IsLink   bool   `json:"isLink"`                      // 是否外链
	Path     string `json:"path"`                        // 路由地址
	Show     bool   `json:"show"`                        // 显示状态
	Status   bool   `json:"status"`                      // 菜单状态
	UserName string `json:"userName"`                    // 用户名称
}

// MenuUpdateRequest 菜单更新
type MenuUpdateRequest struct {
	MenuCreateRequest
	MenuId int64 `json:"menuId" binding:"required"` // 菜单ID
}

// MenuDeleteRequest 菜单删除
type MenuDeleteRequest struct {
	Ids      []int64 `json:"ids" binding:"required"`
	UserName string  `json:"userName"`
}
