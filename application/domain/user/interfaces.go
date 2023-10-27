package user

import (
	"errors"
	"github.com/google/uuid"
)

var (
	// ErrNotFound is returned when a user is not found
	ErrNotFound       = errors.New("user not found")
	ErrMailExists     = errors.New("mail in use")
	ErrNoValidMail    = errors.New("a valid mail should be provided")
	ErrNoValidName    = errors.New("a valid name should be provided")
	ErrNoValidSurname = errors.New("a valid surname should be provided")
	ErrAuthentication = errors.New("wrong credentials")
	ErrUserDeleted    = errors.New("user has been deleted")
)

type Service interface {
	User(id uuid.UUID) (User, error)
	Users(includeDel bool) ([]User, error)
	UserByEmail(email string, includeDel bool) (User, bool)
	UsersByType(t Type) ([]User, error)
	UserByType(id uuid.UUID, t Type) (User, error)
	CreateUser(user User) (User, error)
	Authenticate(email string, password string) (User, error)
	UpdateUser(u User) (User, error)
	DeleteUser(id uuid.UUID) error
}

type Repository interface {
	CreateUser(user User) (User, error)
	User(id uuid.UUID) (User, error)
	Users(includeDel bool) ([]User, error)
	UpdateUser(u User) (User, error)
	DeleteUser(id uuid.UUID) error
}
