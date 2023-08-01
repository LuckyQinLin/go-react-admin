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

var Dict = new(DictService)

type DictService struct{}

// Create 字典创建
func (d *DictService) Create(param *request.DictCreateRequest) *response.BusinessError {
	var (
		dict      entity.DictType
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测岗位名称是否唯一
	condition = core.DB.Where("dict_name = ?", param.DictName)
	if exist, err = dao.Dict.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的字典名称")
		return response.CustomBusinessError(response.Failed, "存在相同的字典名称")
	}
	// 检测岗位编码是否唯一
	condition = core.DB.Where("dict_type = ?", param.DictType)
	if exist, err = dao.Dict.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的字典类型")
		return response.CustomBusinessError(response.Failed, "存在相同的字典类型")
	}
	// 保存菜单数据
	now = time.Now()
	dict = entity.DictType{
		DictName:   param.DictName,
		DictType:   param.DictType,
		CreateBy:   param.UserName,
		Status:     param.Status,
		Remark:     param.Remark,
		CreateTime: &now,
	}
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Dict.Create(tx, &dict)
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "创建字典失败")
	}
	core.Log.Info("创建字典[%d:%s]成功", dict.DictId, dict.DictName)
	return nil
}

// Update 字典更新
func (d *DictService) Update(param *request.DictUpdateRequest) *response.BusinessError {
	var (
		old       entity.DictType
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测岗位名称是否唯一
	condition = core.DB.Where("dict_name = ? and dict_id != ?", param.DictName, param.DictId)
	if exist, err = dao.Dict.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的字典名称")
		return response.CustomBusinessError(response.Failed, "存在相同的字典名称")
	}
	// 检测岗位编码是否唯一
	condition = core.DB.Where("dict_type = ? and dict_id != ?", param.DictType, param.DictId)
	if exist, err = dao.Dict.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的字典类型")
		return response.CustomBusinessError(response.Failed, "存在相同的字典类型")
	}
	if old, err = dao.Dict.GetDictById(param.DictId); err != nil {
		core.Log.Error("当前字典不存在：%s", err.Error())
		return response.CustomBusinessError(response.Failed, "当前字典不存在")
	}
	// 保存部门数据
	now = time.Now()
	old.DictName = param.DictName
	old.DictType = param.DictType
	old.Status = param.Status
	old.Remark = param.Remark
	old.UpdateBy = param.UserName
	old.UpdateTime = &now
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Dict.UpdateById(tx, &old)
	}); err != nil {
		core.Log.Error("更新字典[%d]:%s", param.DictId, err.Error())
		return response.CustomBusinessError(response.Failed, "更新字典失败")
	}
	return nil
}

// Info 字典详情
func (d *DictService) Info(dictId int64) (*response.DictInfoResponse, *response.BusinessError) {
	var (
		dict entity.DictType
		err  error
	)
	if dict, err = dao.Dict.GetDictById(dictId); err != nil {
		return nil, response.CustomBusinessError(response.Failed, "当前字典不存在")
	}
	return &response.DictInfoResponse{
		DictId:   dict.DictId,
		DictName: dict.DictName,
		DictType: dict.DictType,
		Status:   dict.Status,
		Remark:   dict.Remark,
	}, nil
}

// Page 字典分页
func (d *DictService) Page(param *request.DictPageRequest) (*response.PageData, *response.BusinessError) {
	var (
		buildCondition = func(param *request.DictPageRequest) func(db *gorm.DB) *gorm.DB {
			return func(db *gorm.DB) *gorm.DB {
				db.Model(&entity.DictType{})
				if param.Status != 0 {
					db.Where("status = ?", param.Status)
				}
				if param.DictName != "" {
					db.Where("dict_name like concat('%', ?, '%')", param.DictName)
				}
				if param.DictType != "" {
					db.Where("dict_type like concat('%', ?, '%')", param.DictType)
				}
				return db
			}
		}
		list  []response.DictTableResponse
		total int64
		err   error
	)
	if err = core.DB.Scopes(buildCondition(param)).Count(&total).Error; err != nil {
		core.Log.Error("统计字典数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取字典数据失败")
	}
	if err = core.DB.Scopes(buildCondition(param)).
		Select("dict_id, dict_name, dict_type, status, create_time, remark").
		Find(&list).
		Error; err != nil {
		core.Log.Error("查询字典数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取字典数据失败")
	}
	return &response.PageData{
		Total: total,
		Page:  param.Page,
		Size:  param.Size,
		Data:  list,
	}, nil
}

// Delete 字典删除
func (d *DictService) Delete(param *request.DictDeleteRequest, username string) *response.BusinessError {
	var (
		err error
	)
	if err = core.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err = dao.Dict.Delete(tx, param.Ids...); err != nil {
			core.Log.Error("删除字典失败：%s", err.Error())
			return
		}
		return
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "删除字典失败")
	}
	return nil
}
