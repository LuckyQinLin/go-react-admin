package response

// MenuTree 菜单树
type MenuTree struct {
	Key      int64      `json:"key"`
	Label    string     `json:"label"`
	Children []MenuTree `json:"children"`
}
