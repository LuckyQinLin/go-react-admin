package service

import (
	"admin-api/app/dao"
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/core"
	"admin-api/internal/gorm"
	"fmt"
	"github.com/xuri/excelize/v2"
	"time"
)

var Role = new(RoleService)

type RoleService struct{}

// Page 分页查询
func (r *RoleService) Page(param *request.RolePageRequest) (*response.PageData, *response.BusinessError) {
	var (
		buildCondition = func(param *request.RolePageRequest) (condition *gorm.DB) {
			condition = core.DB.Where("del_flag = 1")
			if param.Status != 0 {
				condition.Where("status = ?", param.Status)
			}
			if param.Name != "" {
				condition.Where("(role_name like concat('%', ?, '%') or role_key like concat('%', ?, '%'))", param.Name, param.Name)
			}
			if !param.StartTime.Equal(time.Time{}) && !param.EndTime.Equal(time.Time{}) {
				condition.Where("create_time between ? and ?", param.StartTime, param.EndTime)
			}
			return
		}
		list  []response.RolePageResponse
		total int64
		err   error
	)
	if total, err = dao.Role.Total(buildCondition(param)); err != nil {
		core.Log.Error("统计角色数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取角色数据失败")
	}
	if total == 0 {
		list = make([]response.RolePageResponse, 0)
	} else {
		if list, err = dao.Role.Limit(buildCondition(param), param.Offset(), param.Size); err != nil {
			core.Log.Error("获取角色数据失败, 异常信息如下：%s", err.Error())
			return nil, response.CustomBusinessError(response.Failed, "获取角色数据失败")
		}
	}
	return &response.PageData{
		Total: total,
		Page:  param.Page,
		Size:  param.Size,
		Data:  list,
	}, nil
}

// RoleAll 获取所有角色
func (r *RoleService) RoleAll() (list []response.RoleKeyValueResponse, customErr *response.BusinessError) {
	if err := core.DB.Model(&entity.Role{}).Where("del_flag = 1").Find(&list).Error; err != nil {
		customErr = response.CustomBusinessError(response.Failed, "获取角色数据失败")
	}
	return
}

// Create 角色创建
func (r *RoleService) Create(param *request.RoleCreateRequest) *response.BusinessError {
	var (
		err       error
		isExist   bool
		condition *gorm.DB
		role      entity.Role
		maps      []*entity.RoleMenu
		now       time.Time
	)

	// 判断是否存在相同的角色名称或者权限字符
	condition = core.DB.Where("role_name = ? or role_key = ?", param.RoleName, param.RoleKey)
	if isExist, err = dao.Role.Exist(condition); err != nil || isExist {
		core.Log.Error("存在相同的角色名称或者权限字符")
		return response.CustomBusinessError(response.Failed, "存在相同的角色名称或者权限字符")
	}
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		now = time.Now()
		// 创建角色
		role = entity.Role{
			RoleName:   param.RoleName,
			RoleKey:    param.RoleKey,
			RoleSort:   param.RoleSort,
			Status:     param.Status,
			Remark:     param.Remark,
			DelFlag:    1,
			CreateBy:   param.UserName,
			CreateTime: &now,
		}
		if err = dao.Role.Create(tx, &role); err != nil {
			core.Log.Error("创建角色[%s]失败：%s", param.RoleName, err.Error())
			return err
		}
		// 创建角色菜单映射关系
		if len(param.MenuIds) > 0 {
			for _, id := range param.MenuIds {
				maps = append(maps, &entity.RoleMenu{RoleId: role.RoleId, MenuId: id})
			}
			if err = dao.Role.RoleMenuMapping(tx, maps); err != nil {
				core.Log.Error("创建角色[%s]映射关系失败：%s", param.RoleName, err.Error())
				return err
			}
		}
		return nil
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "创建角色失败")
	}
	core.Log.Info("创建角色[%d:%s]成功", role.RoleId, param.RoleName)
	return nil
}

func (r *RoleService) Update(param *request.RoleUpdateRequest) *response.BusinessError {
	var (
		err       error
		isNeed    bool
		role      entity.Role
		maps      []*entity.RoleMenu
		now       time.Time
		old       entity.Role
		customErr *response.BusinessError
		contrast  = func(old *entity.Role, param *request.RoleUpdateRequest) (isNeed bool, bErr *response.BusinessError) {
			var (
				condition *gorm.DB
				exist     bool
				err       error
			)
			isNeed = false
			if old.RoleName != param.RoleName {
				// 判断新的角色名称是否存在相同的角色信息
				condition = core.DB.Where("role_name = ? and del_flag = 1 and role_id != ?", param.RoleName, old.RoleId)
				if exist, err = dao.Role.Exist(condition); err != nil || exist {
					return false, response.CustomBusinessError(response.Failed, "存在相同的角色名称["+param.RoleName+"]")
				}
				old.RoleName = param.RoleName
				isNeed = true
			}
			if old.RoleKey != param.RoleKey {
				condition = core.DB.Where("role_key = ? and del_flag = 1 and role_id != ?", param.RoleKey, old.RoleId)
				if exist, err = dao.Role.Exist(condition); err != nil || exist {
					return false, response.CustomBusinessError(response.Failed, "存在相同的权限字符["+param.RoleKey+"]")
				}
				old.RoleKey = param.RoleKey
				isNeed = true
			}
			if old.Status != param.Status {
				old.Status = param.Status
				isNeed = true
			}
			if old.RoleSort != param.RoleSort {
				old.RoleSort = param.RoleSort
				isNeed = true
			}
			if old.Remark != param.Remark {
				old.Remark = param.Remark
				isNeed = true
			}
			return isNeed, nil
		} // 对比是否需要更新数据
	)
	// 获取修改数据
	if old, err = dao.Role.GetRoleById(param.RoleId); err != nil {
		core.Log.Error("当前角色[%d]不存在", param.RoleId)
		return response.CustomBusinessError(response.Failed, "当前角色不存在")
	}
	// 判断是否需要修改数据
	if isNeed, customErr = contrast(&old, param); customErr != nil || !isNeed {
		core.Log.Error("修改角色失败：%s", customErr.Error())
		return customErr
	}
	// 判断是否需要更新角色和菜单的授权信息
	if len(param.MenuIds) > 0 {
		for _, id := range param.MenuIds {
			maps = append(maps, &entity.RoleMenu{RoleId: role.RoleId, MenuId: id})
		}
	}
	// 执行更新
	if err = core.DB.Transaction(func(tx *gorm.DB) (err error) {
		if isNeed {
			now = time.Now()
			old.UpdateBy = param.UserName
			old.UpdateTime = &now
			if err = tx.Save(&old).Error; err != nil {
				core.Log.Error("更新角色数据失败:%s", err.Error())
				return
			}
		}
		// 判断是否需要更新映射关系
		if len(maps) > 0 {
			if err = tx.Where("role_id = ?", old.RoleId).Delete(&entity.RoleMenu{}).Error; err != nil {
				core.Log.Error("删除角色菜单旧的映射数据失败:%s", err.Error())
				return
			}
			if err = dao.Role.RoleMenuMapping(tx, maps); err != nil {
				core.Log.Error("创建角色[%s]映射关系失败：%s", param.RoleName, err.Error())
				return
			}
		}
		return
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "更新角色失败")
	}
	return nil
}

// Info 获取角色详情
func (r *RoleService) Info(roleId int64) (*response.RoleInfoResponse, *response.BusinessError) {
	var (
		role    entity.Role
		menus   []entity.RoleMenu
		menuIds []int64
		err     error
	)
	if role, err = dao.Role.GetRoleById(roleId); err != nil {
		return nil, response.CustomBusinessError(response.Failed, "当前角色不存在")
	}
	// 获取角色分配的菜单数据
	if menus, err = dao.Role.GetRoleMenu(roleId); err != nil {
		core.Log.Error("获取角色和菜单的关联数据失败：%s", err.Error())
	}
	if len(menus) > 0 {
		for _, item := range menus {
			menuIds = append(menuIds, item.MenuId)
		}
	}
	return &response.RoleInfoResponse{
		RoleId:   role.RoleId,
		RoleKey:  role.RoleKey,
		RoleName: role.RoleName,
		RoleSort: role.RoleSort,
		Status:   role.Status,
		Remark:   role.Remark,
		MenuIds:  menuIds,
	}, nil
}

// Delete 角色删除（支持批量）
func (r *RoleService) Delete(param *request.RoleDeleteRequest) *response.BusinessError {
	updateField := map[string]any{"update_by": param.UserName, "update_time": time.Now(), "del_flag": 0}
	if err := core.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err = dao.Role.UpdateById(tx, updateField, param.Ids...); err != nil {
			core.Log.Error("删除角色失败：%s", err.Error())
			return
		}
		if err = dao.RoleMenu.Delete(tx, param.Ids...); err != nil {
			core.Log.Error("删除角色关联菜单数据失败：%s", err.Error())
			return
		}
		return
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "删除角色失败")
	}
	return nil
}

// ChangeStatus 角色状态修改
func (r *RoleService) ChangeStatus(param *request.RoleStatusRequest) *response.BusinessError {
	updateField := map[string]any{"update_by": param.UserName, "update_time": time.Now(), "status": param.Status}
	if err := dao.Role.UpdateById(core.DB, updateField, param.RoleId); err != nil {
		core.Log.Error("修改角色状态失败：%s", err.Error())
		return response.CustomBusinessError(response.Failed, "修改角色状态失败")
	}
	return nil
}

func (r *RoleService) DataExport(ids []int64) (file *excelize.File, customErr *response.BusinessError) {
	var (
		sheetStr  string = "角色数据"
		condition *gorm.DB
		sheet     int
		list      []entity.Role
		err       error
	)
	// 创建一个新的Excel文件
	file = excelize.NewFile()
	// 创建一个新的工作表
	if sheet, err = file.NewSheet("角色数据"); err != nil {
		core.Log.Error("导出角色数据失败：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "导出角色数据失败")
	}
	// 设置表头
	_ = file.SetCellValue(sheetStr, "A1", "角色名称")
	_ = file.SetCellValue(sheetStr, "B1", "权限字符")
	_ = file.SetCellValue(sheetStr, "C1", "显示顺序")
	_ = file.SetCellValue(sheetStr, "D1", "角色状态")
	_ = file.SetCellValue(sheetStr, "E1", "创建用户")
	_ = file.SetCellValue(sheetStr, "F1", "备注")
	// 写数据
	condition = core.DB.Where("del_flag = 1 and role_id in ?", ids)
	if list, err = dao.Role.List(condition); err != nil {
		core.Log.Error("读取角色数据失败：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取角色数据失败")
	}
	for i, role := range list {
		_ = file.SetCellValue(sheetStr, fmt.Sprintf("A%d", i+2), role.RoleName)
		_ = file.SetCellValue(sheetStr, fmt.Sprintf("B%d", i+2), role.RoleKey)
		_ = file.SetCellValue(sheetStr, fmt.Sprintf("C%d", i+2), role.RoleSort)
		_ = file.SetCellValue(sheetStr, fmt.Sprintf("D%d", i+2), role.Status)
		_ = file.SetCellValue(sheetStr, fmt.Sprintf("E%d", i+2), role.CreateBy)
		_ = file.SetCellValue(sheetStr, fmt.Sprintf("F%d", i+2), role.Remark)
	}
	file.SetActiveSheet(sheet)
	return file, nil
}

// UserRole 获取用户拥有的角色
func (r *RoleService) UserRole(userId int64) ([]int64, *response.BusinessError) {
	var (
		ids []int64
		err error
	)
	if ids, err = dao.Role.UserRole(userId); err != nil {
		return nil, response.CustomBusinessError(response.Failed, "获取用户角色数据失败")
	}
	return ids, nil
}

// RoleUser 获取角色拥有的用户信息
func (r *RoleService) RoleUser(roleId int64) ([]int64, *response.BusinessError) {
	var (
		ids []int64
		err error
	)
	if ids, err = dao.Role.RoleUser(roleId); err != nil {
		return nil, response.CustomBusinessError(response.Failed, "获取角色用户数据失败")
	}
	return ids, nil
}

// RoleAllocateUser 角色分配用户
func (r *RoleService) RoleAllocateUser(param *request.RoleUserRequest) *response.BusinessError {
	if err := core.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.UserRole.DeleteByRoleId(tx, param.RoleId); err != nil {
			core.Log.Error("删除角色已有用户映射关系失败:%s", err.Error())
			return err
		}
		list := make([]*entity.UserRole, 0)
		for _, userId := range param.UserIds {
			list = append(list, &entity.UserRole{
				UserId: userId,
				RoleId: param.RoleId,
			})
		}
		if err := dao.UserRole.InsertBatch(tx, list); err != nil {
			core.Log.Error("保存角色已有用户映射关系失败:%s", err.Error())
			return err
		}
		return nil
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "给角色分配用户失败")
	}
	return nil
}
