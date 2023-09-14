package web

import (
	"net/http"
	"testing"
)

type User struct{}

func UserQuery(id int) (int, any) {
	return http.StatusOK, nil
}

func TestNewRouters(t *testing.T) {
	// type ControllerFunc func(param ...any) (any, error)
	NewRouters("api").Push(NewRouter("user").Get(UserQuery).Handler())
}
