package engine

import (
	"context"
	"net/http"
)

type IEngine interface {
	http.Handler
	StartServer(ctx context.Context, port int) // 启动服务
}

// IContext 上下文
type IContext interface {
}

// IMethodTree 方法树
type IMethodTree interface {
}
