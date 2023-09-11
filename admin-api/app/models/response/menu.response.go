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
	MenuId    int64  `json:"menuId"`                       // 菜单ID
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

// UserRouterResponse 用户路由信息
type UserRouterResponse struct {
	MenuId    int64                `json:"id"`                // 菜单ID
	ParentId  int64                `json:"pId"`               // 上级菜单
	Icon      string               `json:"icon"`              // 图标
	MenuName  string               `json:"title"`             // 菜单名称
	Perms     string               `json:"perms"`             // 权限字符
	MenuSort  int                  `json:"sort"`              // 显示顺序
	Path      string               `json:"path"`              // 路由地址
	Component string               `json:"component"`         // 组件路由
	Children  []UserRouterResponse `json:"children" gorm:"-"` // 下级路由
}
