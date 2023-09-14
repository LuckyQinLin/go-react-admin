package web

import (
	"goboot/domain/interfaces/engine"
	"goboot/domain/model"
	"net/http"
)

type (
	// Routers 路由
	Routers struct {
		engine        *HttpEngine          // 引擎
		basePath      string               // 根路径
		globalHandler engine.HandlersChain // 全局处理器
		router        []engine.IRouter     // 路由
	}

	// Router 路由
	Router struct {
		path        string                                     // 路径
		handlers    engine.HandlersChain                       // 中间件
		controllers map[model.HttpMethod]engine.ControllerFunc // 控制器
		children    []engine.IRouter                           // 下级路由节点
	}
)

// =====================================================================================================================

// NewRouters 构建路由根节点
func NewRouters(basePath string, handler ...engine.HandlerFunc) engine.IRoutes {
	return &Routers{
		basePath:      basePath,
		globalHandler: handler,
		router:        make([]engine.IRouter, 0),
	}
}

func (r *Routers) Push(router ...engine.IRouter) engine.IRoutes {
	r.router = append(r.router, router...)
	return r
}

func (r *Routers) Use(handler ...engine.HandlerFunc) engine.IRoutes {
	r.globalHandler = append(r.globalHandler, handler...)
	return r
}

// =====================================================================================================================

// NewRouter 构建路由子节点
func NewRouter(path string) engine.IRouter {
	return &Router{
		path:        path,
		controllers: make(map[model.HttpMethod]engine.ControllerFunc),
		children:    make([]engine.IRouter, 0),
		handlers:    make(engine.HandlersChain, 0),
	}
}

// request 请求
func (r *Router) request(method model.HttpMethod, handler engine.ControllerFunc) engine.IRouter {
	r.controllers[method] = handler
	return r
}

// Push 添加下级节点
func (r *Router) Push(nodes ...engine.IRouter) engine.IRouter {
	r.children = append(r.children, nodes...)
	return r
}

// Handler 处理器
func (r *Router) Handler(handler ...engine.HandlerFunc) engine.IRouter {
	r.handlers = append(r.handlers, handler...)
	return r
}

// Get 请求
func (r *Router) Get(handler engine.ControllerFunc) engine.IRouter {
	return r.request(model.MethodGet, handler)
}

// Post 请求
func (r *Router) Post(handler engine.ControllerFunc) engine.IRouter {
	return r.request(model.MethodPost, handler)
}

// StaticFile 静态文件
func (r *Router) StaticFile(s string, s2 string) engine.IRouter {
	//TODO implement me
	panic("implement me")
}

func (r *Router) StaticFileFS(s string, s2 string, system http.FileSystem) engine.IRouter {
	//TODO implement me
	panic("implement me")
}

func (r *Router) Static(s string, s2 string) engine.IRouter {
	//TODO implement me
	panic("implement me")
}

func (r *Router) StaticFS(s string, system http.FileSystem) engine.IRouter {
	//TODO implement me
	panic("implement me")
}
