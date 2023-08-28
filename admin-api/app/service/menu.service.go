package service

import (
	"admin-api/app/dao"
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/core"
	"admin-api/internal/gorm"
	"admin-api/utils"
	"sort"
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
	condition = core.DB.Where("menu_code = ? and parent_id = ?", param.MenuName, param.ParentId)
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
		MenuType:   param.MenuType,
		CreateBy:   param.UserName,
		CreateTime: &now,
	}
	if param.MenuType == "M" {
		// 目录
		menu.Icon = param.Icon
		menu.Path = param.Path
		menu.IsFrame = param.IsLink
		menu.Visible = param.IsShow
		menu.Status = param.Status
	} else if param.MenuType == "C" {
		// 菜单
		menu.Perms = param.Perms
		menu.Icon = param.Icon
		menu.Path = param.Path
		menu.Component = param.Component
		menu.IsFrame = param.IsLink
		menu.IsCache = param.IsCache
		menu.Visible = param.IsShow
		menu.Status = param.Status
	} else if param.MenuType == "F" {
		// 按钮
		menu.Perms = param.Perms
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
	old.UpdateBy = param.UserName
	old.UpdateTime = &now
	if param.MenuType == "M" {
		// 目录
		old.Icon = param.Icon
		old.Path = param.Path
		old.IsFrame = param.IsLink
		old.Visible = param.IsShow
		old.Status = param.Status
	} else if param.MenuType == "C" {
		// 菜单
		old.Perms = param.Perms
		old.Icon = param.Icon
		old.Path = param.Path
		old.Component = param.Component
		old.IsFrame = param.IsLink
		old.IsCache = param.IsCache
		old.Visible = param.IsShow
		old.Status = param.Status
	} else if param.MenuType == "F" {
		// 按钮
		old.Perms = param.Perms
	}
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

// Info 菜单详情
func (m *MenuService) Info(menuId int64) (*response.MenuInfoResponse, *response.BusinessError) {
	var (
		menu entity.Menu
		err  error
	)
	if menu, err = dao.Menu.GetMenuById(menuId); err != nil {
		core.Log.Info("获取菜单失败：%s", menuId)
		return nil, response.CustomBusinessError(response.Failed, "获取菜单失败")
	}
	return &response.MenuInfoResponse{
		MenuId:   menu.MenuId,
		ParentId: menu.ParentId,
		MenuType: menu.MenuType,
		Icon:     menu.Icon,
		MenuName: menu.MenuName,
		MenuSort: menu.OrderNum,
		IsLink:   menu.IsFrame,
		Path:     menu.Path,
		IsShow:   menu.Visible,
		Status:   menu.Status,
	}, nil
}

// UserRouter 用户路由
func (m *MenuService) UserRouter(userId int64, roleId int64) ([]response.UserRouterResponse, *response.BusinessError) {
	var (
		buildTree func(data []response.UserRouterResponse, parentId int64) []response.UserRouterResponse
		result    []response.UserRouterResponse
		roleIds   []int64
		err       error
	)
	// 构建Tree
	buildTree = func(data []response.UserRouterResponse, parentId int64) []response.UserRouterResponse {
		var children []response.UserRouterResponse
		for _, item := range data {
			if parentId == item.ParentId {
				item.Children = buildTree(data, item.MenuId)
				children = append(children, item)
			}
		}
		// 排序
		sort.Slice(children, func(i, j int) bool {
			return children[i].MenuSort > children[j].MenuSort
		})
		return children
	}
	// 构建查询条件
	condition := core.DB.Model(&entity.Menu{}).
		//Debug().
		Alias("sm").
		Select("sm.menu_id,sm.menu_code as menu_name,sm.parent_id,sm.order_num as menu_sort,sm.path,sm.component,sm.perms,sm.icon").
		Where("(sm.menu_type = 'M' or sm.menu_type = 'C')")
	// 获取当前用户的角色信息
	if roleIds, err = dao.User.UserRoleId(userId); err != nil {
		core.Log.Error("当前用户不存在角色信息")
		return nil, response.CustomBusinessError(response.Failed, "获取用户路由失败")
	}
	if len(roleIds) == 0 {
		if userId != vo.SUPER_USER_ID {
			core.Log.Error("当前用户不存在角色信息")
			return nil, response.CustomBusinessError(response.Failed, "获取用户路由失败")
		}
	} else {
		if roleId != 0 {
			if utils.Contain(roleId, roleIds) {
				condition.Where("exists(select 1 from sys_role_menu srm where srm.role_id = ? and srm.menu_id = sm.menu_id)", roleId)
			} else {
				core.Log.Error("当前用户不存在该角色信息")
				return nil, response.CustomBusinessError(response.Failed, "获取用户路由失败")
			}
		} else {
			condition.Where("exists(select 1 from sys_role_menu srm where srm.role_id in ? and srm.menu_id = sm.menu_id)", roleIds)
		}
	}
	// 查询数据
	if err = condition.Find(&result).Error; err != nil {
		core.Log.Error("查询用户的路由信息发生异常：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取用户路由失败")
	}
	return buildTree(result, 0), nil

}
