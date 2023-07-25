package request

// MenuTableQueryRequest 菜单表格查询
type MenuTableQueryRequest struct {
	Name   string `json:"name"`   // 菜单名称
	Status int    `json:"status"` // 菜单状态
}
