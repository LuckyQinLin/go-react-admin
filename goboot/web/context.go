package web

import (
	"goboot/domain/interfaces/engine"
	"net/http"
	"net/url"
	"sync"
)

type (
	Pool  sync.Pool
	Param struct {
		Key   string
		Value string
	}
	Params        []Param
	EngineContext struct {
		Request    *http.Request
		Response   engine.IResponse
		Keys       map[string]any
		Params     Params
		FullPath   string
		queryCache url.Values
		formCache  url.Values
		sameSite   http.SameSite
		lock       sync.RWMutex
		engine     engine.IEngine
	}
)

func (e *EngineContext) GetRequest() *http.Request {
	return e.Request
}

func (e *EngineContext) UpdateRequest(request *http.Request) engine.IContext {
	e.Request = request
	return e
}

func (e *EngineContext) ResponseClear(w http.ResponseWriter) engine.IContext {
	e.Response.Clear(w)
	return e
}

func (e *EngineContext) Reset() engine.IContext {
	//TODO implement me
	panic("implement me")
}

func (e *EngineContext) Copy() engine.IContext {
	//TODO implement me
	panic("implement me")
}

func NewPool(engine engine.IEngine) *sync.Pool {
	return &sync.Pool{New: func() any {
		return NewContext(engine)
	}}
}

func NewContext(engine engine.IEngine) engine.IContext {
	return &EngineContext{engine: engine}
}

// GetContext 获取context
func (p *Pool) GetContext() engine.IContext {
	return nil
}
