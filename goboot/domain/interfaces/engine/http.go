package engine

import (
	"context"
	"goboot/domain/model"
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
	Push(...IRouter) IRoutes                  // 添加下级路由
	Use(...FilterHandler) IRoutes             // 使用拦截器
	Match(path string, method string) IRouter // 匹配路由
	Filter() FilterChains                     // 获取全局拦截器
}

// IRouter 路由节点定义
type IRouter interface {
	Handler(method model.HttpMethod, ctx IContext) HandlerFunc

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
	Reset() IContext
	Copy() IContext
	ResponseClear(http.ResponseWriter) IContext // 重置
	UpdateRequest(*http.Request) IContext       // 设置请求
	GetRequest() *http.Request                  // 获取请求
}

type IParams interface {
	Get(string) (string, error)
}

type IResponse interface {
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	Status() int
	// Size returns the number of bytes already written into the response http body.
	// See Written()
	Size() int
	// WriteString writes the string into the response body.
	WriteString(string) (int, error)
	// Written returns true if the response body was already written.
	Written() bool
	// WriteHeaderNow forces to write the http header (status code + headers).
	WriteHeaderNow()
	// Pusher get the http.Pusher for server push
	Pusher() http.Pusher
	Clear(http.ResponseWriter)
}
