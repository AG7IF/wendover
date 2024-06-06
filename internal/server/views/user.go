package views

import (
	"github.com/google/uuid"

	"github.com/ag7if/wendover/pkg/auth"
)

type RoleActivityKeyView struct {
	ActivityKey string    `json:"activity"`
	Role        auth.Role `json:"role"`
}

type RoleActivityNameView struct {
	ActivityName string    `json:"activity_name"`
	Role         auth.Role `json:"role"`
}

func NewRoleView(role auth.UserRole) RoleActivityNameView {
	roleActivity := role.Activity()

	return RoleActivityNameView{
		ActivityName: roleActivity.Name(),
		Role:         role.Role(),
	}
}

type UserView struct {
	ID       uuid.UUID                       `json:"id"`
	Username string                          `json:"username"`
	Roles    map[string]RoleActivityNameView `json:"roles,omitempty"`
}

func NewUserView(user auth.User) UserView {
	rvs := make(map[string]RoleActivityNameView)
	for k, v := range user.UserRoles() {
		rvs[k] = NewRoleView(v)
	}

	return UserView{
		ID:       user.ID(),
		Username: user.Username(),
		Roles:    rvs,
	}
}

func (uv UserView) ToDomainObject() auth.User {
	return auth.NewUser(
		uv.ID,
		uv.Username,
		nil, // Roles are not added through this view, so we just simply ignore whatever comes in here.
	)
}
