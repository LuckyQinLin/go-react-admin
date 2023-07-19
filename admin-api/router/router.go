package router

import (
	"admin-api/app/controller"
	"admin-api/core"
	"admin-api/internal/gin"
	"context"
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
		user := root.Group("user")
		{
			user.GET("captchaImage", controller.User.CaptchaImage) // 获取验证码
			user.POST("login", controller.User.Login)              // 登陆
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
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
