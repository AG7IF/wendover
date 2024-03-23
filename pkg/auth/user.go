package auth

import (
	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	username  string
	email     string
	userRoles map[string]UserRole
}

func NewUser(
	id uuid.UUID,
	username string,
	email string,
	userRoles map[string]UserRole,
) User {
	return User{
		id:        id,
		username:  username,
		email:     email,
		userRoles: userRoles,
	}
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Username() string {
	return u.username
}

func (u User) Email() string {
	return u.email
}

func (u User) UserRoles() map[string]UserRole {
	return u.userRoles
}
