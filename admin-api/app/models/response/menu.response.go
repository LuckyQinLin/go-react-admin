package response

import (
	"admin-api/utils"
	"encoding/json"
	"time"
)

// MenuTree 菜单树
type MenuTree struct {
	Key      int64       `json:"key"`
	Label    string      `json:"title"`
	Children []*MenuTree `json:"children"`
}

// MenuTableResponse 菜单表
type MenuTableResponse struct {
	MenuId     int64                `json:"key"`        // 主键
	MenuName   string               `json:"title"`      // 菜单名称
	Perms      string               `json:"code"`       // 权限标识
	Icon       string               `json:"icon"`       // 图标
	Path       string               `json:"path"`       // 路由
	ParentId   int64                `json:"parentId"`   // 上级ID
	Status     int                  `json:"status"`     // 状态
	OrderNum   int                  `json:"order"`      // 排序
	CreateTime *time.Time           `json:"createTime"` // 创建时间
	Children   []*MenuTableResponse `json:"children"`   // 子数据
}

func (s MenuTableResponse) MarshalJSON() ([]byte, error) {
	type temp MenuTableResponse
	return json.Marshal(&struct {
		temp
		CreateTime utils.DateTime `json:"createTime"`
	}{
		temp:       (temp)(s),
		CreateTime: utils.DateTime(*s.CreateTime),
	})
}

// MenuInfoResponse 菜单详情
type MenuInfoResponse struct {
	MenuId   int64  `json:"menuId"`   // 菜单ID
	ParentId int64  `json:"parentId"` // 上级菜单
	MenuType string `json:"menuType"` // 菜单类型 (M目录 C菜单 F按钮)
	Icon     string `json:"icon"`     // 图标
	MenuName string `json:"menuName"` // 菜单名称
	MenuSort int    `json:"menuSort"` // 显示顺序
	IsLink   bool   `json:"isLink"`   // 是否外链
	Path     string `json:"path"`     // 路由地址
	Show     bool   `json:"show"`     // 显示状态
	Status   int    `json:"status"`   // 菜单状态
}
