package web

import (
	"context"
	"goboot/domain/interfaces/engine"
	"goboot/domain/interfaces/logger"
	"net"
	"net/http"
	"sync"
)

// HttpEngine httpEngine 引擎
type HttpEngine struct {
	trees              engine.IRoutes  // 用于优化路由匹配的数据结构
	logger             logger.ILogger  // 日志
	context            engine.IContext // 上下文
	pool               sync.Pool       // 临时对象的池
	maxParams          uint16          // 用于配置路由参数的最大数量
	maxSections        uint16          // 用于配置路由节的最大数量
	trustedProxies     []string
	trustedCIDRs       []*net.IPNet
	MaxMultipartMemory int64 // 上传文件最大内存
}

// ServeHTTP 处理http请求的handler
func (engine *HttpEngine) ServeHTTP(response http.ResponseWriter, request *http.Request) {

}

func (engine *HttpEngine) StartServer(ctx context.Context, port int) {

}
