package response

// MenuTree 菜单树
type MenuTree struct {
	Key      int64       `json:"key"`
	Label    string      `json:"title"`
	Children []*MenuTree `json:"children"`
}
