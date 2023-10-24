package response

import (
	"admin-api/utils"
	"encoding/json"
	"time"
)

// RolePageResponse 角色分页返回数据
type RolePageResponse struct {
	RoleId     int64     `json:"roleId"`     // 角色Id
	RoleKey    string    `json:"roleKey"`    // 角色编码
	RoleName   string    `json:"roleName"`   // 角色名称
	RoleSort   int       `json:"roleSort"`   // 角色排序
	Status     int       `json:"status"`     // 角色状态
	IsSuper    bool      `json:"isSuper"`    // 是否超级角色
	CreateTime time.Time `json:"createTime"` // 创建时间
}

func (s RolePageResponse) MarshalJSON() ([]byte, error) {
	type temp RolePageResponse
	return json.Marshal(&struct {
		temp
		CreateTime utils.DateTime `json:"createTime"`
	}{
		temp:       (temp)(s),
		CreateTime: utils.DateTime(s.CreateTime),
	})
}

// RoleKeyValueResponse 角色key/value返回数据
type RoleKeyValueResponse struct {
	RoleId   int64  `json:"value"` // 角色ID
	RoleName string `json:"label"` // 角色名称
}

// RoleInfoResponse 角色详情
type RoleInfoResponse struct {
	RoleId   int64   `json:"roleId"`   // 角色Id
	RoleKey  string  `json:"roleKey"`  // 角色编码
	RoleName string  `json:"roleName"` // 角色名称
	RoleSort int     `json:"roleSort"` // 角色排序
	Status   int     `json:"status"`   // 角色状态
	Remark   string  `json:"remark"`   // 备注
	MenuIds  []int64 `json:"menuIds"`  // 菜单ID
}

// RoleListResponse 角色信息
type RoleListResponse struct {
	RoleId   int64  `json:"value"`
	RoleName string `json:"label"`
}
