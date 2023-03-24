package user

import (
	"errors"
	"github.com/google/uuid"
)

var (
	// ErrNotFound is returned when a user is not found
	ErrNotFound       = errors.New("user not found")
	ErrUserExists     = errors.New("user already exists")
	ErrAuthentication = errors.New("wrong credentials")
)

type Service interface {
	User(id uuid.UUID) (User, error)
	Users() map[uuid.UUID]User
	CreateUser(user User) error
	Authenticate(email string, password string) (User, error)
	UpdateUser(u User) error
	DeleteUser(id uuid.UUID) error
}

type Repository interface {
	CreateUser(user User) error
	User(id uuid.UUID) (User, error)
	Users() map[uuid.UUID]User
	UpdateUser(u User) error
	DeleteUser(id uuid.UUID) error
}
