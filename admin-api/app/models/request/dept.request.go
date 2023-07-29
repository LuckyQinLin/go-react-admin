package request

// DeptTableQueryRequest 部门表查询
type DeptTableQueryRequest struct {
	Name   string `json:"name"`   // 部门名称
	Status int    `json:"status"` // 部门状态
}

// DeptCreateRequest 部门创建
type DeptCreateRequest struct {
	ParentId int64  `json:"parentId"`                    // 上级部门
	DeptName string `json:"deptName" binding:"required"` // 部门名称
	OrderNum int    `json:"orderNum"`                    // 显示顺序
	Leader   string `json:"leader"`                      // 负责人
	Phone    string `json:"phone"`                       // 联系电话
	Email    string `json:"email"`                       // 邮箱
	Status   int    `json:"status"`                      // 部门状态
	UserName string `json:"userName"`                    // 用户名称
}

// DeptUpdateRequest 部门更新
type DeptUpdateRequest struct {
	DeptCreateRequest
	DeptId int64 `json:"deptId" binding:"required"`
}
