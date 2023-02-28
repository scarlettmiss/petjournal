package user

import "errors"

var (
	// ErrNotFound is returned when a user is not found
	ErrNotFound       = errors.New("user not found")
	ErrUserExists     = errors.New("user already exists")
	ErrAuthentication = errors.New("wrong credentials")
)

type Service interface {
	User(id string) (User, error)
	Users() map[string]User
	CreateUser(user User) error
	Authenticate(email string, password string) (User, error)
	UpdateUser(u User) error
	DeleteUser(id string) error
}

type Repository interface {
	CreateUser(user User) error
	User(id string) (User, error)
	Users() map[string]User
	UpdateUser(u User) error
	DeleteUser(id string) error
}
