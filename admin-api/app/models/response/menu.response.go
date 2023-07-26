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
// id: number;       // 主键
//
//	title: string;    // 菜单名称
//	code: string;     // 权限字符
//	icon: string;     // 图标
//	path?: string;    // 路由
//	parentId: number; // 上级ID
//	status: number;   // 状态
//	order: number;    // 排序
//	createTime: string; // 创建时间
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
