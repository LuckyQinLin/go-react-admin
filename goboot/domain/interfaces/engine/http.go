package engine

import (
	"context"
	"net/http"
)

type HandlerFunc func(ctx IContext)

// HandlersChain 中间件
type HandlersChain []HandlerFunc

// ControllerFunc 控制器
type ControllerFunc any

type Tree []IMethodTree

type IEngine interface {
	http.Handler
	StartServer(ctx context.Context, port int) // 启动服务
}

// IRoutes 根路由定义
type IRoutes interface {
	Push(...IRouter) IRoutes    // 添加下级路由
	Use(...HandlerFunc) IRoutes // 使用插件
}

// IRouter 路由节点定义
type IRouter interface {
	Push(nodes ...IRouter) IRouter          // 添加下级
	Get(handler ControllerFunc) IRouter     // Get请求
	Post(handler ControllerFunc) IRouter    // Post请求
	Handler(handler ...HandlerFunc) IRouter // 添加连接

	StaticFile(string, string) IRouter
	StaticFileFS(string, string, http.FileSystem) IRouter
	Static(string, string) IRouter
	StaticFS(string, http.FileSystem) IRouter
}

// IContext 上下文
type IContext interface {
}

// IMethodTree 方法树
type IMethodTree interface {
}
