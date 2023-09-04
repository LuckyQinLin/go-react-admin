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

var Config = new(ConfigService)

type ConfigService struct{}

// Create 参数创建
func (d *ConfigService) Create(param *request.ConfigCreateRequest) *response.BusinessError {
	var (
		dict      entity.Setting
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测岗位名称是否唯一
	condition = core.DB.Where("config_name = ?", param.ConfigName)
	if exist, err = dao.Config.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的参数名称")
		return response.CustomBusinessError(response.Failed, "存在相同的参数名称")
	}
	// 检测岗位编码是否唯一
	condition = core.DB.Where("config_key = ?", param.ConfigKey)
	if exist, err = dao.Config.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的参数类型")
		return response.CustomBusinessError(response.Failed, "存在相同的参数类型")
	}
	// 保存菜单数据
	now = time.Now()
	dict = entity.Setting{
		ConfigName:  param.ConfigName,
		ConfigKey:   param.ConfigKey,
		ConfigValue: param.ConfigValue,
		ConfigType:  param.ConfigType,
		CreateBy:    param.UserName,
		Remark:      param.Remark,
		CreateTime:  &now,
	}
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Config.Create(tx, &dict)
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "创建参数失败")
	}
	core.Log.Info("创建参数[%d:%s]成功", dict.ConfigId, dict.ConfigName)
	return nil
}

// Update 参数更新
func (d *ConfigService) Update(param *request.ConfigUpdateRequest) *response.BusinessError {
	var (
		old       entity.Setting
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测岗位名称是否唯一
	condition = core.DB.Where("config_name = ? and config_id != ?", param.ConfigName, param.ConfigId)
	if exist, err = dao.Config.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的参数名称")
		return response.CustomBusinessError(response.Failed, "存在相同的参数名称")
	}
	// 检测岗位编码是否唯一
	condition = core.DB.Where("config_key = ? and config_id != ?", param.ConfigKey, param.ConfigId)
	if exist, err = dao.Config.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的参数类型")
		return response.CustomBusinessError(response.Failed, "存在相同的参数类型")
	}
	if old, err = dao.Config.GetById(param.ConfigId); err != nil {
		core.Log.Error("当前参数不存在：%s", err.Error())
		return response.CustomBusinessError(response.Failed, "当前参数不存在")
	}
	// 保存部门数据
	now = time.Now()
	old.ConfigName = param.ConfigName
	old.ConfigKey = param.ConfigKey
	old.ConfigValue = param.ConfigValue
	old.ConfigType = param.ConfigType
	old.Remark = param.Remark
	old.UpdateBy = param.UserName
	old.UpdateTime = &now
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Config.UpdateById(tx, &old)
	}); err != nil {
		core.Log.Error("更新参数[%d]:%s", param.ConfigId, err.Error())
		return response.CustomBusinessError(response.Failed, "更新参数失败")
	}
	return nil
}

// Info 参数详情
func (d *ConfigService) Info(configId int64) (*response.ConfigInfoResponse, *response.BusinessError) {
	var (
		config entity.Setting
		err    error
	)
	if config, err = dao.Config.GetById(configId); err != nil {
		return nil, response.CustomBusinessError(response.Failed, "当前参数不存在")
	}
	return &response.ConfigInfoResponse{
		ConfigId:    config.ConfigId,
		ConfigName:  config.ConfigName,
		ConfigKey:   config.ConfigKey,
		ConfigValue: config.ConfigValue,
		ConfigType:  config.ConfigType,
		Remark:      config.Remark,
	}, nil
}

// Page 参数分页
func (d *ConfigService) Page(param *request.ConfigPageRequest) (*response.PageData, *response.BusinessError) {
	var (
		buildCondition = func(param *request.ConfigPageRequest) func(db *gorm.DB) *gorm.DB {
			return func(db *gorm.DB) *gorm.DB {
				db.Model(&entity.Setting{}).Template(`
					1 = 1
					{% if param.ConfigType != -1 %}
						and config_type = {{param.ConfigType}}
					{% endif %}
					{% if param.ConfigName != "" %}
						and config_name like concat('%', {{param.ConfigName}}, '%')
					{% endif %}
					{% if param.ConfigKey != "" %}
						and config_key like concat('%', {{param.ConfigKey}}, '%')
					{% endif %}
				`, map[string]any{"param": param})
				//if param.ConfigType != -1 {
				//	db.Where("config_type = ?", param.ConfigType)
				//}
				//if param.ConfigName != "" {
				//	db.Where("config_name like concat('%', ?, '%')", param.ConfigName)
				//}
				//if param.ConfigKey != "" {
				//	db.Where("config_key like concat('%', ?, '%')", param.ConfigKey)
				//}
				return db
			}
		}
		list  []response.ConfigTableResponse
		total int64
		err   error
	)
	if err = core.DB.Scopes(buildCondition(param)).Debug().Count(&total).Error; err != nil {
		core.Log.Error("统计参数数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取参数数据失败")
	}
	if err = core.DB.Scopes(buildCondition(param)).
		Find(&list).
		Error; err != nil {
		core.Log.Error("查询参数数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取参数数据失败")
	}
	return &response.PageData{
		Total: total,
		Page:  param.Page,
		Size:  param.Size,
		Data:  list,
	}, nil
}

// Delete 参数删除
func (d *ConfigService) Delete(param *request.ConfigDeleteRequest, username string) *response.BusinessError {
	var (
		err error
	)
	if err = core.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err = dao.Config.Delete(tx, param.Ids...); err != nil {
			core.Log.Error("删除参数失败：%s", err.Error())
			return
		}
		return
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "删除参数失败")
	}
	return nil
}
