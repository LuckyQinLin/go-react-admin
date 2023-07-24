package request

import "time"

// RolePageRequest 角色分页查询
type RolePageRequest struct {
	CommonPage
	Name      string    `json:"name"`      // 角色名称或者权限字符
	Status    int       `json:"status"`    // 角色状态
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
}

// RoleCreateRequest 角色创建
type RoleCreateRequest struct {
	RoleName string  `json:"roleName" binding:"required"` // 角色名称
	RoleKey  string  `json:"roleKey" binding:"required"`  // 权限字符
	RoleSort int     `json:"roleSort"`                    // 显示顺序
	Status   int     `json:"status"`                      // 状态
	Remark   string  `json:"remark"`                      // 备注
	MenuIds  []int64 `json:"menuIds"`                     // 菜单IDs
	UserName string  `json:"userName"`                    // 创建用户名称
}
