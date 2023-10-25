package request

// RolePageRequest 角色分页查询
type RolePageRequest struct {
	CommonPage
	Name      string `json:"name"`      // 角色名称或者权限字符
	Status    int    `json:"status"`    // 角色状态
	StartTime string `json:"startTime"` // 开始时间
	EndTime   string `json:"endTime"`   // 结束时间
	SuperId   int64
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

// RoleUpdateRequest 角色修改
type RoleUpdateRequest struct {
	RoleId int64 `json:"roleId" binding:"required"` // 角色ID
	RoleCreateRequest
}

// RoleStatusRequest 角色状态修改
type RoleStatusRequest struct {
	RoleId   int64  `json:"roleId" binding:"required"` // 角色ID
	Status   int    `json:"status" binding:"required"` // 状态
	UserName string `json:"userName"`                  // 创建用户名称
}

// RoleDeleteRequest 角色删除
type RoleDeleteRequest struct {
	Ids      []int64 `json:"ids" binding:"required"`
	UserName string  `json:"userName"`
}

// RoleUserRequest 角色分配用户
type RoleUserRequest struct {
	RoleId  int64   `json:"roleId"`
	UserIds []int64 `json:"userIds"`
}
