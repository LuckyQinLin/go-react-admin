package engine

import (
	"context"
	"net/http"
)

// FilterHandler 拦截器处理器
type FilterHandler interface {
	After(ctx IContext)  // 之后
	Before(ctx IContext) // 之前
	Order() int          // 排序
}

// HandlerFunc 业务处理器
type HandlerFunc func(ctx IContext)

// FilterChains 中间件
type FilterChains []FilterHandler

type IEngine interface {
	http.Handler
	StartServer(ctx context.Context, port int) // 启动服务
}

// IRoutes 根路由定义
type IRoutes interface {
	Push(...IRouter) IRoutes      // 添加下级路由
	Use(...FilterHandler) IRoutes // 使用拦截器
	Match(path string) IRouter    // 匹配路由
}

// IRouter 路由节点定义
type IRouter interface {
	Push(nodes ...IRouter) IRouter    // 添加下级
	Get(handler HandlerFunc) IRouter  // Get请求
	Post(handler HandlerFunc) IRouter // Post请求

	StaticFile(string, string) IRouter
	StaticFileFS(string, string, http.FileSystem) IRouter
	Static(string, string) IRouter
	StaticFS(string, http.FileSystem) IRouter
}

// IContext 上下文
type IContext interface {
}

type IParams interface {
	Get(string) (string, error)
}
