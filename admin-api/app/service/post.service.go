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

var Post = new(PostService)

type PostService struct{}

// Page 分页查询
func (p *PostService) Page(param *request.PostPageRequest) (*response.PageData, *response.BusinessError) {
	var (
		buildCondition = func(param *request.PostPageRequest) func(db *gorm.DB) *gorm.DB {
			return func(db *gorm.DB) *gorm.DB {
				db.Model(&entity.Post{})
				if param.Status != 0 {
					db.Where("status = ?", param.Status)
				}
				if param.PostName != "" {
					db.Where("post_name like concat('%', ?, '%')", param.PostName)
				}
				if param.PostCode != "" {
					db.Where("post_code like concat('%', ?, '%')", param.PostCode)
				}
				return db
			}
		}
		list  []response.PostPageResponse
		total int64
		err   error
	)
	if err = core.DB.Scopes(buildCondition(param)).Count(&total).Error; err != nil {
		core.Log.Error("统计岗位数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取角色数据失败")
	}
	if err = core.DB.Scopes(buildCondition(param)).
		Select("post_id, post_name, post_code, post_sort, create_time, status").
		Find(&list).
		Error; err != nil {
		core.Log.Error("查询岗位数据失败, 异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取角色数据失败")
	}
	return &response.PageData{
		Total: total,
		Page:  param.Page,
		Size:  param.Size,
		Data:  list,
	}, nil
}

// Create 创建岗位
func (p *PostService) Create(param *request.PostCreateRequest) *response.BusinessError {
	var (
		post      entity.Post
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测岗位名称是否唯一
	condition = core.DB.Where("post_name = ?", param.PostName)
	if exist, err = dao.Post.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的岗位名称")
		return response.CustomBusinessError(response.Failed, "存在相同的岗位名称")
	}
	// 检测岗位编码是否唯一
	condition = core.DB.Where("post_code = ?", param.PostCode)
	if exist, err = dao.Post.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的岗位编码")
		return response.CustomBusinessError(response.Failed, "存在相同的岗位编码")
	}
	// 保存菜单数据
	now = time.Now()
	post = entity.Post{
		PostName:   param.PostName,
		PostCode:   param.PostCode,
		PostSort:   param.PostSort,
		CreateBy:   param.UserName,
		Status:     param.Status,
		Remark:     param.Remark,
		CreateTime: &now,
	}
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Post.Create(tx, &post)
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "创建岗位失败")
	}
	core.Log.Info("创建岗位[%d:%s]成功", post.PostId, post.PostName)
	return nil
}

// Delete 删除岗位
func (p *PostService) Delete(param *request.PostDeleteRequest, username string) *response.BusinessError {
	var (
		condition *gorm.DB
		err       error
		exist     bool
	)
	condition = core.DB.Where("post_id in ?", param.Ids)
	if exist, err = dao.UserPost.Exist(condition); err != nil || exist {
		core.Log.Error("存在已经分配用户的岗位")
		return response.CustomBusinessError(response.Failed, "存在已经分配用户的岗位")
	}
	if err := core.DB.Transaction(func(tx *gorm.DB) (err error) {
		if err = dao.Post.Delete(tx, param.Ids...); err != nil {
			core.Log.Error("删除岗位失败：%s", err.Error())
			return
		}
		return
	}); err != nil {
		return response.CustomBusinessError(response.Failed, "删除岗位失败")
	}
	return nil
}

// Info 岗位信息
func (p *PostService) Info(postId int64) (*response.PostInfoResponse, *response.BusinessError) {
	var (
		post entity.Post
		err  error
	)
	if post, err = dao.Post.GetById(postId); err != nil {
		return nil, response.CustomBusinessError(response.Failed, "当前岗位不存在")
	}
	return &response.PostInfoResponse{
		PostId:   post.PostId,
		PostName: post.PostName,
		PostCode: post.PostCode,
		PostSort: post.PostSort,
		Status:   post.Status,
		Remark:   post.Remark,
	}, nil
}

// Update 岗位更新
func (p *PostService) Update(param *request.PostUpdateRequest) *response.BusinessError {
	var (
		old       entity.Post
		now       time.Time
		condition *gorm.DB
		err       error
		exist     bool
	)
	// 检测岗位名称是否唯一
	condition = core.DB.Where("post_name = ? and post_id != ?", param.PostName, param.PostId)
	if exist, err = dao.Post.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的岗位名称")
		return response.CustomBusinessError(response.Failed, "存在相同的岗位名称")
	}
	// 检测岗位编码是否唯一
	condition = core.DB.Where("post_code = ? and post_id != ?", param.PostCode, param.PostId)
	if exist, err = dao.Post.Exist(condition); err != nil || exist {
		core.Log.Error("存在相同的岗位编码")
		return response.CustomBusinessError(response.Failed, "存在相同的岗位编码")
	}
	if old, err = dao.Post.GetById(param.PostId); err != nil {
		core.Log.Error("当前岗位不存在：%s", err.Error())
		return response.CustomBusinessError(response.Failed, "当前岗位不存在")
	}
	// 保存部门数据
	now = time.Now()
	old.PostName = param.PostName
	old.PostCode = param.PostCode
	old.PostSort = param.PostSort
	old.Status = param.Status
	old.Remark = param.Remark
	old.UpdateBy = param.UserName
	old.UpdateTime = &now
	if err = core.DB.Transaction(func(tx *gorm.DB) error {
		return dao.Post.UpdateById(tx, &old)
	}); err != nil {
		core.Log.Error("更新部门[%d]:%s", param.PostId, err.Error())
		return response.CustomBusinessError(response.Failed, "更新岗位失败")
	}
	return nil
}

// All 获取全部岗位
func (p *PostService) All() (list []*response.PostListResponse, customErr *response.BusinessError) {
	var (
		posts []entity.Post
		err   error
	)
	if posts, err = dao.Post.List(nil); err != nil {
		return nil, response.CustomBusinessError(response.Failed, "获取岗位失败")
	}
	for _, post := range posts {
		list = append(list, &response.PostListResponse{
			PostId:   post.PostId,
			PostName: post.PostName,
		})
	}
	return
}
