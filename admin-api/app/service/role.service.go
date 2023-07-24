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
