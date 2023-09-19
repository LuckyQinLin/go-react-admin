package web

import (
	"goboot/domain/interfaces/engine"
	"testing"
)

type RoleRequest struct {
}

type User struct{}

func UserQuery(ctx engine.IContext) {
}

func UserList(ctx engine.IContext) {
}

func RoleCreate(ctx engine.IContext) {
}

func TestNewRouters(t *testing.T) {
	// type ControllerFunc func(param ...any) (any, error)
	// GET /api/user/list?id=1
	routes := NewRouters("api").
		Push(NewRouter("user").Get(UserQuery).
			Push(NewRouter("list").Get(UserList))).
		Push(NewRouter("role").Post(RoleCreate))

	t.Log(routes)
}
