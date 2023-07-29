package router

import (
	"admin-api/app/controller"
	"admin-api/core"
	"admin-api/internal/gin"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var Engine *gin.Engine

func initRouter() {
	core.InitCore()
	core.Log.Info("初始化路由配置")
	Engine = gin.Default(core.Log)
	Engine.Use(CorsMiddle()) // 跨域
	Engine.Use(JwtMiddle())  // jwt
	root := Engine.Group(core.Config.Web.ContextPath)
	{
		user := root.Group("user") // 用户模块
		{
			user.GET("captchaImage", controller.User.CaptchaImage) // 获取验证码
			user.POST("login", controller.User.Login)              // 登陆
			user.GET("info", controller.User.GetUserInfo)          // 获取用户信息
			user.POST("userAllot", controller.User.AllotRole)      // 用户分配角色
		}
		role := root.Group("role") // 角色模块
		{
			role.GET("all", controller.Role.RoleAll)            // 所有角色
			role.POST("page", controller.Role.Page)             // 角色分页
			role.POST("create", controller.Role.RoleCreate)     // 角色创建
			role.POST("update", controller.Role.RoleUpdate)     // 角色修改
			role.GET("info", controller.Role.RoleInfo)          // 角色详情
			role.POST("status", controller.Role.RoleStatus)     // 角色状态修改
			role.POST("delete", controller.Role.RoleDelete)     // 角色删除
			role.POST("dataAuth", controller.Role.RoleDataAuth) // 角色数据权限分配
			role.GET("export", controller.Role.RoleExport)      // 角色导出
		}
		menu := root.Group("menu") // 菜单模块
		{
			menu.GET("tree", controller.Menu.MenuTree)      // 菜单树
			menu.POST("table", controller.Menu.MenuTable)   // 菜单表格
			menu.POST("create", controller.Menu.MenuCreate) // 菜单创建
			menu.POST("update", controller.Menu.MenuUpdate) // 菜单修改
			menu.GET("delete", controller.Menu.MenuDelete)  // 菜单删除
			menu.GET("info", controller.Menu.MenuInfo)      // 菜单详情
		}
		dept := root.Group("dept") // 部门模块
		{
			dept.GET("tree", controller.Dept.DeptTree)        // 部门树
			dept.POST("table", controller.Dept.DeptTableTree) // 部门表格
			dept.POST("create", controller.Dept.DeptCreate)   // 部门创建
			dept.POST("update", controller.Dept.DeptUpdate)   // 部门修改
			dept.GET("delete", controller.Dept.DeptDelete)    // 部门删除
			dept.GET("info", controller.Dept.DeptInfo)        // 部门详情
		}
		post := root.Group("post") // 岗位模块
		{
			post.POST("page", controller.Post.PostPage)     // 岗位分页
			post.POST("create", controller.Post.PostCreate) // 岗位创建
			post.POST("update", controller.Post.PostUpdate) // 岗位修改
			post.GET("delete", controller.Post.PostDelete)  // 岗位删除
			post.GET("info", controller.Post.PostInfo)      // 岗位详情
			post.GET("export", controller.Post.PostExport)  // 岗位导出
		}
	}
}

func Run(port int64) {
	initRouter()
	gin.SetMode(core.Config.Web.RunModel)
	if port == 0 {
		port = core.Config.Web.Port
	}
	core.Log.Info("服务器启动在端口[%d]", port)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        Engine,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: core.Config.Web.MaxHeaderBytes * 1024 * 1024,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			core.Log.Error("启动服务失败：%s\n", err.Error())
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	core.Log.Info("服务正在停止")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		core.Log.Error("服务停止失败：%s", err.Error())
	}
	core.Log.Info("服务停止")
}
