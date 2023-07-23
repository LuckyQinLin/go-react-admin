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
