package web

import (
	"bufio"
	"io"
	"net"
	"net/http"
)

const (
	noWritten     = -1
	defaultStatus = http.StatusOK
)

type EngineResponse struct {
	http.ResponseWriter
	size   int
	status int
}

func (e *EngineResponse) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if e.size < 0 {
		e.size = 0
	}
	return e.ResponseWriter.(http.Hijacker).Hijack()
}

func (e *EngineResponse) Flush() {
	e.WriteHeaderNow()
	e.ResponseWriter.(http.Flusher).Flush()
}

func (e *EngineResponse) Status() int {
	return e.status
}

func (e *EngineResponse) Size() int {
	return e.size
}

func (e *EngineResponse) Write(data []byte) (n int, err error) {
	e.WriteHeaderNow()
	n, err = e.ResponseWriter.Write(data)
	e.size += n
	return
}

func (e *EngineResponse) WriteString(s string) (n int, err error) {
	e.WriteHeaderNow()
	n, err = io.WriteString(e.ResponseWriter, s)
	e.size += n
	return
}

func (e *EngineResponse) Written() bool {
	return e.size != noWritten
}

func (e *EngineResponse) WriteHeaderNow() {
	if !e.Written() {
		e.size = 0
		e.ResponseWriter.WriteHeader(e.status)
	}
}

func (e *EngineResponse) Pusher() http.Pusher {
	if pusher, ok := e.ResponseWriter.(http.Pusher); ok {
		return pusher
	}
	return nil
}

func (e *EngineResponse) Clear(w http.ResponseWriter) {
	e.ResponseWriter = w
	e.size = noWritten
	e.status = defaultStatus
}
