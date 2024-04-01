package user

import (
	"errors"
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application"
)

var (
	// ErrNotFound is returned when a user is not found
	ErrNotFound            = errors.New("user not found")
	ErrMailExists          = errors.New("mail in use")
	ErrNoValidMail         = errors.New("a valid mail should be provided")
	ErrNoValidName         = errors.New("a valid name should be provided")
	ErrNoValidSurname      = errors.New("a valid surname should be provided")
	ErrAuthentication      = errors.New("wrong credentials")
	ErrUserDeleted         = errors.New("user has been deleted")
	ErrNoValidType         = errors.New("a valid userType should be provided")
	ErrPasswordLength      = errors.New("password should be of 8 characters long")
	ErrPasswordLowerCase   = errors.New("password should contain at least one lower case character")
	ErrPasswordUpperCase   = errors.New("password should contain at least one upper case character")
	ErrPasswordDigit       = errors.New("password should contain atleast one digit")
	ErrPasswordSpecialChar = errors.New("password should contain at least one special character")
)

type Service interface {
	User(id uuid.UUID) (User, error)
	Users(includeDel bool) ([]User, error)
	UsersByType(t Type, includeDel bool) ([]User, error)
	UserByType(id uuid.UUID, t Type, includeDel bool) (User, error)
	CreateUser(user application.UserCreateOptions) (User, string, error)
	UpdateUser(opts application.UserUpdateOptions, includeDel bool) (User, error)
	Authenticate(email string, password string) (User, string, error)
	DeleteUser(id uuid.UUID) error
	userByEmail(email string, includeDel bool) (User, bool)
	checkEmail(email string, id uuid.UUID, includeDel bool) error
	passwordValidation(password string) error
	userToken(user User) (string, error)
}

type Repository interface {
	CreateUser(user User) (User, error)
	User(id uuid.UUID) (User, error)
	Users(includeDel bool) ([]User, error)
	UpdateUser(u User) (User, error)
	DeleteUser(id uuid.UUID) error
}
