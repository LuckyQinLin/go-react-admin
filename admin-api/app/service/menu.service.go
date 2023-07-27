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

var Menu = new(MenuService)

type MenuService struct{}

// Tree 菜单树
func (m *MenuService) Tree() ([]*response.MenuTree, *response.BusinessError) {
	var (
		recursionBuild func(pId int64, data []entity.Menu) []*response.MenuTree
		tree           []*response.MenuTree
		all            []entity.Menu
		err            error
	)
	// 递归构建树
	recursionBuild = func(pId int64, data []entity.Menu) []*response.MenuTree {
		var children []*response.MenuTree
		for _, item := range data {
			if pId == item.ParentId {
				child := &response.MenuTree{
					Key:   item.MenuId,
					Label: item.MenuName,
				}
				child.Children = recursionBuild(item.MenuId, data)
				children = append(children, child)
			}
		}
		return children
	}
	if all, err = dao.Menu.All(); err != nil {
		core.Log.Error("获取菜单数据发生异常：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取菜单数据失败")
	}
	if len(all) <= 0 {
		return make([]*response.MenuTree, 0), nil
	}
	tree = recursionBuild(0, all)
	return tree, nil
}

// Table 表格数据查询
func (m *MenuService) Table(param *request.MenuTableQueryRequest) ([]*response.MenuTableResponse, *response.BusinessError) {
	var (
		recursionBuild func(pId int64, data []entity.Menu) []*response.MenuTableResponse
		result         []*response.MenuTableResponse
		all            []entity.Menu
		err            error
	)
	// 递归构建树
	recursionBuild = func(pId int64, data []entity.Menu) []*response.MenuTableResponse {
		var children []*response.MenuTableResponse
		for _, item := range data {
			if pId == item.ParentId {
				child := &response.MenuTableResponse{
					MenuId:     item.MenuId,
					MenuName:   item.MenuName,
					ParentId:   item.ParentId,
					Status:     item.Status,
					Perms:      item.Perms,
					OrderNum:   item.OrderNum,
					Icon:       item.Icon,
					Path:       item.Path,
					CreateTime: item.CreateTime,
				}
				child.Children = recursionBuild(item.MenuId, data)
				children = append(children, child)
			}
		}
		return children
	}
	if all, err = dao.Menu.All(); err != nil {
		core.Log.Error("获取菜单数据发生异常：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取菜单数据失败")
	}
	if len(all) <= 0 {
		return nil, nil
	}
	result = recursionBuild(0, all)
	return result, nil
}

// Create 菜单创建
func (m *MenuService) Create(param *request.MenuCreateRequest) *response.BusinessError {
	var (
		menu      entity.Menu
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测菜单名称是否唯一
	condition = core.DB.Where("menu_code = ? or parent_id = ?", param.MenuName, param.ParentId)
	if exist, err = dao.Menu.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的资源名称")
		return response.CustomBusinessError(response.Failed, "存在相同的资源名称")
	}
	// 保存菜单数据
	now = time.Now()
	menu = entity.Menu{
		MenuName:   param.MenuName,
		ParentId:   param.ParentId,
		OrderNum:   param.MenuSort,
		Path:       param.Path,
		IsFrame:    param.IsLink,
		IsCache:    true,
		MenuType:   param.MenuType,
		Visible:    param.Show,
		Status:     param.Status,
		Icon:       param.Icon,
		CreateBy:   param.UserName,
		CreateTime: &now,
	}
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Menu.Create(tx, &menu)
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "创建菜单失败")
	}
	core.Log.Info("创建菜单[%d:%s]成功", menu.MenuId, param.MenuName)
	return nil
}

// Update 菜单修改
func (m *MenuService) Update(param *request.MenuUpdateRequest) *response.BusinessError {
	var (
		old       entity.Menu
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测菜单名称是否唯一
	condition = core.DB.Where("menu_code = ? or parent_id = ?", param.MenuName, param.ParentId)
	if exist, err = dao.Menu.Exist(condition); err != nil || !exist {
		core.Log.Error("存在相同的资源名称")
		return response.CustomBusinessError(response.Failed, "存在相同的资源名称")
	}
	// 获取修改的数据
	if old, err = dao.Menu.GetMenuById(param.MenuId); err != nil {
		core.Log.Error("当前菜单[%d]不存在", param.MenuId)
		return response.CustomBusinessError(response.Failed, "当前菜单不存在")
	}
	// 保存菜单数据
	now = time.Now()
	old.ParentId = param.ParentId
	old.MenuType = param.MenuType
	old.MenuName = param.MenuName
	old.OrderNum = param.MenuSort
	old.IsFrame = param.IsLink
	old.Path = param.Path
	old.Visible = param.Show
	old.Status = param.Status
	old.Icon = param.Icon
	old.UpdateBy = param.UserName
	old.UpdateTime = &now
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Menu.UpdateById(tx, &old)
	}); err != nil {
		core.Log.Error("更新菜单[%d]:%s", param.MenuId, err.Error())
		return response.CustomBusinessError(response.Failed, "更新菜单失败")
	}
	return nil
}

// Delete 菜单删除
func (m *MenuService) Delete(menuId int64, updateName string) *response.BusinessError {
	var (
		exist     bool
		condition *gorm.DB
		err       error
	)
	// 菜单是否存在子菜单
	condition = core.DB.Where("parent_id = ?", menuId)
	if exist, err = dao.Menu.Exist(condition); err != nil || exist {
		core.Log.Error("菜单是否存在子菜单")
		return response.CustomBusinessError(response.Failed, "菜单是否存在子菜单")
	}
	// 菜单是否已经分配角色
	condition = core.DB.Where("menu_id = ?", menuId)
	if exist, err = dao.RoleMenu.Exist(condition); err != nil || exist {
		core.Log.Error("菜单已经分配角色")
		return response.CustomBusinessError(response.Failed, "菜单已经分配角色")
	}
	// 删除
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Menu.Delete(tx, menuId)
	}); err != nil {
		core.Log.Error("删除菜单失败：%s", err.Error())
		return response.CustomBusinessError(response.Failed, "删除菜单失败")
	}
	return nil
}
