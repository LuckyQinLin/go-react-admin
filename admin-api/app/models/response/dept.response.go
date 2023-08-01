package response

import (
	"admin-api/utils"
	"encoding/json"
	"time"
)

// DeptTree 菜单树
type DeptTree struct {
	Key      int64       `json:"key"`
	Label    string      `json:"title"`
	Children []*DeptTree `json:"children"`
}

// DeptTableResponse 部门表属性
type DeptTableResponse struct {
	DeptId     int64                `json:"key"`        // 部门ID
	DeptName   string               `json:"title"`      // 部门名称
	DeptSort   int                  `json:"order"`      // 部门排序
	Status     int                  `json:"status"`     // 部门状态
	CreateTime *time.Time           `json:"createTime"` // 创建时间
	Children   []*DeptTableResponse `json:"children"`   // 子部门
}

func (s DeptTableResponse) MarshalJSON() ([]byte, error) {
	type temp DeptTableResponse
	return json.Marshal(&struct {
		temp
		CreateTime utils.DateTime `json:"createTime"`
	}{
		temp:       (temp)(s),
		CreateTime: utils.DateTime(*s.CreateTime),
	})
}

// DeptInfoResponse 部门详情
type DeptInfoResponse struct {
	DeptId   int64  `json:"deptId"`                      // 部门ID
	ParentId int64  `json:"parentId"`                    // 上级部门
	DeptName string `json:"deptName" binding:"required"` // 部门名称
	OrderNum int    `json:"orderNum"`                    // 显示顺序
	Leader   string `json:"leader"`                      // 负责人
	Phone    string `json:"phone"`                       // 联系电话
	Email    string `json:"email"`                       // 邮箱
	Status   int    `json:"status"`                      // 部门状态
}
