package controller

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/app/service"
	"gitee.com/molonglove/goboot/gin"
	"net/http"
)

var Post = new(PostController)

type PostController struct{ BaseController }

// PostList 获取全部岗位信息
func (p *PostController) PostList(ctx *gin.Context) {
	var (
		list      []*response.PostListResponse
		customErr *response.BusinessError
	)
	if list, customErr = service.Post.All(); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(list))
}

// PostCreate 岗位创建
func (p *PostController) PostCreate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.PostCreateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = p.Parse(ctx, "岗位创建", vo.Add, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		p.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Post.Create(&param); customErr != nil {
		p.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	p.Success(ctx, operate, response.Ok("岗位创建成功"))
}

// PostUpdate 岗位更新
func (p *PostController) PostUpdate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.PostUpdateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = p.Parse(ctx, "岗位修改", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		p.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UserName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.Post.Update(&param); customErr != nil {
		p.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	p.Success(ctx, operate, response.Ok("岗位修改成功"))
}

// PostDelete 岗位删除
func (p *PostController) PostDelete(ctx *gin.Context) {
	var (
		err       error
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.PostDeleteRequest
		customErr *response.BusinessError
	)
	claims, operate = p.Parse(ctx, "岗位删除", vo.Delete, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		p.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	operate.ParamToJson(param)
	if customErr = service.Post.Delete(&param, claims.Username); customErr != nil {
		p.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	p.Success(ctx, operate, response.Ok("删除岗位数据成功"))
}

// PostPage 岗位分页
func (p *PostController) PostPage(ctx *gin.Context) {
	var (
		param     request.PostPageRequest
		result    *response.PageData
		customErr *response.BusinessError
		err       error
	)
	if err = ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.Post.Page(&param); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(result))
}

// PostInfo 岗位详情
func (p *PostController) PostInfo(ctx *gin.Context) {
	var (
		err       error
		postId    int64
		info      *response.PostInfoResponse
		customErr *response.BusinessError
	)
	if postId, err = ctx.QueryInt64("postId"); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if info, customErr = service.Post.Info(postId); customErr != nil {
		ctx.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(info))
}

// PostExport 岗位导出
func (p *PostController) PostExport(ctx *gin.Context) {

}
