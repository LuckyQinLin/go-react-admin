package engine

import (
	"context"
	"net/http"
)

type HandlerFunc func(ctx IContext)

type HandlersChain []HandlerFunc

type Tree []IMethodTree

type IEngine interface {
	http.Handler
	StartServer(ctx context.Context, port int) // 启动服务
}

// IRoutes defines all router handle interface.
type IRoutes interface {
	Children(string, ...HandlerFunc) IRoutes // 添加下级路由

	Use(...HandlerFunc) IRoutes // 使用插件

	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes     // Any 任何请求
	Get(string, ...HandlerFunc) IRoutes     // GET请求
	Post(string, ...HandlerFunc) IRoutes    // POST请求
	Delete(string, ...HandlerFunc) IRoutes  // DELETE请求
	Patch(string, ...HandlerFunc) IRoutes   // PATCH请求
	Put(string, ...HandlerFunc) IRoutes     // PUT请求
	Options(string, ...HandlerFunc) IRoutes // OPTIONS请求
	Head(string, ...HandlerFunc) IRoutes    // HEAD请求
	Match([]string, string, ...HandlerFunc) IRoutes

	StaticFile(string, string) IRoutes
	StaticFileFS(string, string, http.FileSystem) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, http.FileSystem) IRoutes
}

// IContext 上下文
type IContext interface {
}

// IMethodTree 方法树
type IMethodTree interface {
}
