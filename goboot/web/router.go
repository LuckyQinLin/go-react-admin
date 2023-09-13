package web

import (
	"goboot/domain/interfaces/engine"
	"net/http"
)

// EngineRouter 路由
type EngineRouter struct {
	path    string               // 路径
	engine  *HttpEngine          // 引擎
	handler engine.HandlersChain // 中间件
}

func (e EngineRouter) Children(s string, handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Use(handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Handle(s string, s2 string, handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Any(s string, handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Get(s string, handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Post(s string, handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Delete(s string, handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Patch(s string, handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Put(s string, handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Options(s string, handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Head(s string, handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Match(strings []string, s string, handlerFunc ...engine.HandlerFunc) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) StaticFile(s string, s2 string) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) StaticFileFS(s string, s2 string, system http.FileSystem) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) Static(s string, s2 string) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}

func (e EngineRouter) StaticFS(s string, system http.FileSystem) engine.IRoutes {
	//TODO implement me
	panic("implement me")
}
