package web

import (
	"context"
	"goboot/domain/interfaces/engine"
	"goboot/domain/interfaces/logger"
	"goboot/domain/model"
	"net"
	"net/http"
	"sync"
)

// HttpEngine httpEngine 引擎
type HttpEngine struct {
	trees              engine.IRoutes  // 用于优化路由匹配的数据结构
	logger             logger.ILogger  // 日志
	context            engine.IContext // 上下文
	pool               *sync.Pool      // 临时对象的池
	maxParams          uint16          // 用于配置路由参数的最大数量
	maxSections        uint16          // 用于配置路由节的最大数量
	trustedProxies     []string
	trustedCIDRs       []*net.IPNet
	MaxMultipartMemory int64 // 上传文件最大内存
}

func Default() engine.IEngine {
	h := new(HttpEngine)
	h.pool = NewPool(h)
	return h
}

// ServeHTTP 处理http请求的handler
func (h *HttpEngine) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	ctx := h.getCtx(). // 获取上下文
				ResponseClear(response). // 设置返回
				UpdateRequest(request).  // 设置请求头
				Reset()                  // 初始化上下文
	h.reallyHandler(ctx) // 执行请求
	h.pool.Put(ctx)      // 将请求后的数据放入对象池中
}

func (h *HttpEngine) getCtx() engine.IContext {
	return h.pool.Get().(engine.IContext)
}

// reallyHandler 真正处理请求的handler
func (h *HttpEngine) reallyHandler(ctx engine.IContext) {
	method, path := ctx.GetRequest().Method, ctx.GetRequest().URL.Path
	router := h.trees.Match(path, method) // 匹配路由
	filter := h.trees.Filter()            // 获取全局拦截器
	for _, handler := range filter {      // 执行前置拦截器
		handler.Before(ctx)
	}
	_ = router.Handler(model.MethodGet, ctx) // 执行业务处理器
	for _, handler := range filter {         // 执行后置拦截器
		handler.After(ctx)
	}
}

func (h *HttpEngine) StartServer(ctx context.Context, port int) {

}
