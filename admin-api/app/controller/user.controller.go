package controller

import (
	"admin-api/app/models/entity"
	"admin-api/app/models/request"
	"admin-api/app/models/response"
	"admin-api/app/models/vo"
	"admin-api/app/service"
	"admin-api/utils"
	"gitee.com/molonglove/goboot/gin"
	"net/http"
)

var User = new(UserController)

type UserController struct {
	BaseController
}

// CaptchaImage 获取验证码
func (u *UserController) CaptchaImage(c *gin.Context) {
	var (
		result *response.CaptchaImageResponse
		err    *response.BusinessError
	)
	if result, err = service.User.CaptchaImage(); err != nil {
		c.JSON(http.StatusOK, response.ResultCustom(err))
		return
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// Login 登陆
func (u *UserController) Login(c *gin.Context) {
	var (
		err       error
		customErr *response.BusinessError
		param     request.UserLoginRequest
		result    *response.UserLoginResponse
	)
	if err = c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.User.UserLogin(&param, c); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.OkMsg(response.LoginSuccess, result))
}

// GetUserInfo 获取用户信息
func (u *UserController) GetUserInfo(c *gin.Context) {
	var (
		userId    int64
		customErr *response.BusinessError
		result    *response.UserInfoResponse
		err       error
	)
	if userId, err = c.QueryInt64("userId"); err != nil {
		c.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.User.GetUserInfo(userId); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// UserLoginInfo 获取用户登录信息
func (u *UserController) UserLoginInfo(c *gin.Context) {
	var (
		claims    *vo.UserClaims
		customErr *response.BusinessError
		result    *response.UserLoginInfoResponse
	)
	claims = u.GetCurrentUser(c)
	if result, customErr = service.User.UserLoginInfo(claims.UserId); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// Page 分页
func (u *UserController) Page(c *gin.Context) {
	var (
		param     request.UserPageRequest
		result    *response.PageData
		customErr *response.BusinessError
		err       error
	)
	if err = c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusOK, response.Fail(response.RequestParamError))
		return
	}
	if result, customErr = service.User.Page(&param); customErr != nil {
		c.JSON(http.StatusOK, response.ResultCustom(customErr))
		return
	}
	c.JSON(http.StatusOK, response.Ok(result))
}

// UserCreate 用户创建
func (u *UserController) UserCreate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.UserCreateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = u.Parse(ctx, "用户创建", vo.Add, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		u.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.CreateName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.User.Create(&param); customErr != nil {
		u.Failed(ctx, operate, response.ResultCustom(customErr))
		return
	}
	u.Success(ctx, operate, response.Ok("用户创建成功"))
}

// UserUpdate 用户更新
func (u *UserController) UserUpdate(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.UserUpdateRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = u.Parse(ctx, "用户修改", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		u.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UpdateName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.User.Update(&param); customErr != nil {
		u.Failed(ctx, operate, response.ResultCustom(customErr))
	}
	u.Success(ctx, operate, response.Ok("用户修改成功"))
}

// ResetPassword 重置密码
func (u *UserController) ResetPassword(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.UserPasswordRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = u.Parse(ctx, "修改用户密码", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		u.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UpdateName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.User.ResetPassword(&param); customErr != nil {
		u.Failed(ctx, operate, response.ResultCustom(customErr))
	}
	u.Success(ctx, operate, response.Ok("修改用户密码成功"))
}

// ChangeStatus 修改状态
func (u *UserController) ChangeStatus(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.UserStatusRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = u.Parse(ctx, "修改用户状态", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		u.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.UpdateName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.User.ChangeStatus(&param); customErr != nil {
		u.Failed(ctx, operate, response.ResultCustom(customErr))
	}
	u.Success(ctx, operate, response.Ok("修改用户状态成功"))
}

// UserDelete 用户删除
func (u *UserController) UserDelete(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.UserDeleteRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = u.Parse(ctx, "删除用户", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		u.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	if utils.Include[int64](param.Ids, claims.UserId) {
		u.Failed(ctx, operate, response.Fail(response.UserNotAllowDelete))
		return
	}
	param.UpdateName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.User.DeleteUser(&param); customErr != nil {
		u.Failed(ctx, operate, response.ResultCustom(customErr))
	}
	u.Success(ctx, operate, response.Ok("删除用户成功"))
}

// UserRole 用户分配角色
func (u *UserController) UserRole(ctx *gin.Context) {
	var (
		claims    *vo.UserClaims
		operate   *entity.Operate
		param     request.UserRoleRequest
		customErr *response.BusinessError
		err       error
	)
	claims, operate = u.Parse(ctx, "用户分配角色", vo.Update, nil)
	if err = ctx.ShouldBind(&param); err != nil {
		u.Failed(ctx, operate, response.Fail(response.RequestParamError))
		return
	}
	param.CreateName = claims.Username
	operate.ParamToJson(param)
	if customErr = service.User.UserRole(&param); customErr != nil {
		u.Failed(ctx, operate, response.ResultCustom(customErr))
	}
	u.Success(ctx, operate, response.Ok("用户分配角色成功"))
}

// UserRoutes 获取用户菜单
func (u *UserController) UserRoutes(ctx *gin.Context) {

}
