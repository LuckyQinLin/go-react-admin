package goboot

import "net/http"

type Config struct {
}

// BootEngine 引擎
type BootEngine struct {
	http *http.Server
}

// Start 启动boot
func (b *BootEngine) Start() {

}
