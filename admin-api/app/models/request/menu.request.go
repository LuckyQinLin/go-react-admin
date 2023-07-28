package request

// MenuTableQueryRequest 菜单表格查询
type MenuTableQueryRequest struct {
	Name   string `json:"name"`   // 菜单名称
	Status int    `json:"status"` // 菜单状态
}

// MenuCreateRequest 菜单创建
type MenuCreateRequest struct {
	ParentId  int64  `json:"parentId"`                     // 上级菜单
	MenuType  string `json:"menuType" binding:"required"`  // 菜单类型 (M目录 C菜单 F按钮)
	Icon      string `json:"icon"`                         // 图标
	MenuName  string `json:"menuName"  binding:"required"` // 菜单名称
	Perms     string `json:"perms"`                        // 权限字符
	MenuSort  int    `json:"menuSort"`                     // 显示顺序
	Path      string `json:"path"`                         // 路由地址
	Component string `json:"component"`                    // 组件路由
	Param     string `json:"param"`                        // 路由参数
	IsLink    bool   `json:"isLink"`                       // 是否外链
	IsShow    bool   `json:"isShow"`                       // 显示状态
	IsCache   bool   `json:"isCache"`                      // 缓冲状态
	Status    int    `json:"status"`                       // 菜单状态
	UserName  string `json:"userName"`                     // 用户名称
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
