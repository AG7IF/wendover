package auth

import (
	"github.com/ag7if/wendover/pkg/org"
)

type UserRole struct {
	activity org.Activity
	role     Role
}

func NewUserRole(
	activity org.Activity,
	role Role,
) UserRole {
	return UserRole{
		activity: activity,
		role:     role,
	}
}

func (ur UserRole) Activity() org.Activity {
	return ur.activity
}

func (ur UserRole) Role() Role {
	return ur.role
}
