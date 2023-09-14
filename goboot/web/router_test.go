package web

import (
	"net/http"
	"testing"
)

type RoleRequest struct {
}

type User struct{}

func UserQuery(id int) (int, any) {
	return http.StatusOK, nil
}

func UserList() ([]User, error) {
	return nil, nil
}

func RoleCreate(param RoleRequest) error {
	return nil
}

func TestNewRouters(t *testing.T) {
	// type ControllerFunc func(param ...any) (any, error)
	// GET /api/user/list?id=1
	routes := NewRouters("api").
		Push(NewRouter("user").Get(UserQuery).
			Push(NewRouter("list").Get(UserList)),
		).Push(NewRouter("role").Post(RoleCreate))

	t.Log(routes)
}
