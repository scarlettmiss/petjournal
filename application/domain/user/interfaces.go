package user

import (
	"errors"
	"github.com/google/uuid"
)

var (
	// ErrNotFound is returned when a user is not found
	ErrNotFound       = errors.New("user not found")
	ErrMailExists     = errors.New("mail in use")
	ErrAuthentication = errors.New("wrong credentials")
)

type Service interface {
	User(id uuid.UUID) (User, error)
	Users() ([]User, error)
	CreateUser(user User) (User, error)
	Authenticate(email string, password string) (User, error)
	UpdateUser(u User) (User, error)
	DeleteUser(id uuid.UUID) error
	UserByEmail(email string) (User, bool)
}

type Repository interface {
	CreateUser(user User) (User, error)
	User(id uuid.UUID) (User, error)
	Users() ([]User, error)
	UpdateUser(u User) (User, error)
	DeleteUser(id uuid.UUID) error
}
