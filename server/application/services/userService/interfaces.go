package service

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	"github.com/scarlettmiss/petJournal/application/services"
)

type Service interface {
	User(id uuid.UUID) (user.User, error)
	Users(includeDel bool) ([]user.User, error)
	UsersByType(t user.Type, includeDel bool) ([]user.User, error)
	UserByType(id uuid.UUID, t user.Type, includeDel bool) (user.User, error)
	CreateUser(user services.UserCreateOptions) (user.User, string, error)
	UpdateUser(opts services.UserUpdateOptions, includeDel bool) (user.User, error)
	Authenticate(email string, password string) (user.User, string, error)
	DeleteUser(id uuid.UUID) error
	userByEmail(email string, includeDel bool) (user.User, bool)
	checkEmail(email string, id uuid.UUID, includeDel bool) error
}

type Repository interface {
	CreateUser(user user.User) (user.User, error)
	User(id uuid.UUID) (user.User, error)
	Users(includeDel bool) ([]user.User, error)
	UpdateUser(u user.User) (user.User, error)
	DeleteUser(id uuid.UUID) error
}
