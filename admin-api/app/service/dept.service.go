package service

import (
	"admin-api/app/dao"
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/core"
	"admin-api/internal/gorm"
	"time"
)

var Dept = new(DeptService)

type DeptService struct{}

// Tree 部门数
func (s *DeptService) Tree() ([]*response.DeptTree, *response.BusinessError) {
	var (
		recursionBuild func(pId int64, data []entity.Dept) []*response.DeptTree
		tree           []*response.DeptTree
		all            []entity.Dept
		err            error
	)
	// 递归构建树
	recursionBuild = func(pId int64, data []entity.Dept) []*response.DeptTree {
		var children []*response.DeptTree
		for _, item := range data {
			if pId == item.ParentId {
				child := &response.DeptTree{
					Key:   item.DeptId,
					Label: item.DeptName,
				}
				child.Children = recursionBuild(item.DeptId, data)
				children = append(children, child)
			}
		}
		return children
	}
	if all, err = dao.Dept.All(); err != nil {
		core.Log.Error("获取部门数据发生异常：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取部门数据失败")
	}
	if len(all) <= 0 {
		return make([]*response.DeptTree, 0), nil
	}
	tree = recursionBuild(0, all)
	return tree, nil
}

// Table 部门树表
func (s *DeptService) Table(param *request.DeptTableQueryRequest) ([]*response.DeptTableResponse, *response.BusinessError) {
	var (
		recursionBuild func(pId int64, data []entity.Dept) []*response.DeptTableResponse
		result         []*response.DeptTableResponse
		all            []entity.Dept
		err            error
	)
	// 递归构建树
	recursionBuild = func(pId int64, data []entity.Dept) []*response.DeptTableResponse {
		var children []*response.DeptTableResponse
		for _, item := range data {
			if pId == item.ParentId {
				child := &response.DeptTableResponse{
					DeptId:     item.DeptId,
					DeptName:   item.DeptName,
					DeptSort:   item.OrderNum,
					Status:     item.Status,
					CreateTime: item.CreateTime,
				}
				child.Children = recursionBuild(item.DeptId, data)
				children = append(children, child)
			}
		}
		return children
	}
	if all, err = dao.Dept.All(); err != nil {
		core.Log.Error("获取部门数据发生异常：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取部门数据失败")
	}
	if len(all) <= 0 {
		return nil, nil
	}
	result = recursionBuild(0, all)
	return result, nil
}

// Create 部门创建
func (s *DeptService) Create(param *request.DeptCreateRequest) *response.BusinessError {
	var (
		dept      entity.Dept
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测菜单名称是否唯一
	condition = core.DB.Where("dept_name = ? and parent_id = ?", param.DeptName, param.ParentId)
	if exist, err = dao.Dept.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的部门名称")
		return response.CustomBusinessError(response.Failed, "存在相同的部门名称")
	}
	// 保存菜单数据
	now = time.Now()
	dept = entity.Dept{
		DeptName:   param.DeptName,
		ParentId:   param.ParentId,
		OrderNum:   param.OrderNum,
		CreateBy:   param.UserName,
		Leader:     param.Leader,
		Phone:      param.Phone,
		Email:      param.Email,
		Status:     param.Status,
		CreateTime: &now,
	}
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Dept.Create(tx, &dept)
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "创建部门失败")
	}
	core.Log.Info("创建部门[%d:%s]成功", dept.DeptId, dept.DeptName)
	return nil
}

// Update 部门更新
func (s *DeptService) Update(param *request.DeptUpdateRequest) *response.BusinessError {
	var (
		old       entity.Dept
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测部门名称是否唯一
	condition = core.DB.Where("dept_name = ? or parent_id = ?", param.DeptName, param.ParentId)
	if exist, err = dao.Dept.Exist(condition); err != nil || !exist {
		core.Log.Error("存在相同的部门名称")
		return response.CustomBusinessError(response.Failed, "存在相同的部门名称")
	}
	// 获取修改的数据
	if old, err = dao.Dept.GetDeptById(param.DeptId); err != nil {
		core.Log.Error("当前部门[%d]不存在", param.DeptId)
		return response.CustomBusinessError(response.Failed, "当前部门不存在")
	}
	// 保存部门数据
	now = time.Now()
	old.ParentId = param.ParentId
	old.DeptName = param.DeptName
	old.Leader = param.Leader
	old.Status = param.Status
	old.OrderNum = param.OrderNum
	old.Email = param.Email
	old.Phone = param.Phone
	old.UpdateBy = param.UserName
	old.UpdateTime = &now
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Dept.UpdateById(tx, &old)
	}); err != nil {
		core.Log.Error("更新部门[%d]:%s", param.DeptId, err.Error())
		return response.CustomBusinessError(response.Failed, "更新部门失败")
	}
	return nil
}

// Info 部门详情
func (s *DeptService) Info(deptId int64) (*response.DeptInfoResponse, *response.BusinessError) {
	var (
		menu entity.Dept
		err  error
	)
	if menu, err = dao.Dept.GetDeptById(deptId); err != nil {
		core.Log.Info("获取部门失败：%s", deptId)
		return nil, response.CustomBusinessError(response.Failed, "获取部门失败")
	}
	return &response.DeptInfoResponse{
		DeptId:   menu.DeptId,
		ParentId: menu.ParentId,
		DeptName: menu.DeptName,
		OrderNum: menu.OrderNum,
		Leader:   menu.Leader,
		Email:    menu.Email,
		Phone:    menu.Phone,
		Status:   menu.Status,
	}, nil
}

// Delete 部门删除
func (s *DeptService) Delete(deptId int64, username string) *response.BusinessError {
	var (
		exist     bool
		condition *gorm.DB
		err       error
	)
	// 部门是否存在下级部门
	condition = core.DB.Where("parent_id = ?", deptId)
	if exist, err = dao.Dept.Exist(condition); err != nil || exist {
		core.Log.Error("部门存在下级部门")
		return response.CustomBusinessError(response.Failed, "部门存在下级部门")
	}
	// 部门存在用户
	condition = core.DB.Where("dept_id = ?", deptId)
	if exist, err = dao.User.Exist(condition); err != nil || exist {
		core.Log.Error("部门下存在用户")
		return response.CustomBusinessError(response.Failed, "部门下存在用户")
	}
	// 删除
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Dept.Delete(tx, deptId)
	}); err != nil {
		core.Log.Error("删除部门失败：%s", err.Error())
		return response.CustomBusinessError(response.Failed, "删除部门失败")
	}
	return nil
}
