package service

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	authUtils "github.com/scarlettmiss/petJournal/utils/authorization"
)

type Service struct {
	repo user.Repository
}

func New(repo user.Repository) (Service, error) {
	return Service{repo: repo}, nil
}

func (s *Service) User(id uuid.UUID) (user.User, error) {
	u, err := s.repo.User(id)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *Service) Users(includeDel bool) ([]user.User, error) {
	return s.repo.Users(includeDel)
}

func (s *Service) UserByEmail(email string, includeDel bool) (user.User, bool) {
	var u user.User
	var found bool

	users, err := s.Users(includeDel)
	if err != nil {
		return user.Nil, false
	}
	for _, v := range users {
		if v.Email == email {
			u = v
			found = true
			break
		}
	}
	return u, found
}

func (s *Service) UsersByType(t user.Type, includeDel bool) ([]user.User, error) {
	var users []user.User

	allUsers, err := s.Users(includeDel)
	if err != nil {
		return users, err
	}

	for _, u := range allUsers {
		if u.UserType == t {
			users = append(users, u)
		}
	}

	return users, nil
}

func (s *Service) UserByType(id uuid.UUID, t user.Type, includeDel bool) (user.User, error) {
	var u user.User

	users, err := s.UsersByType(t, includeDel)
	if err != nil {
		return user.Nil, err
	}
	err = user.ErrNotFound
	for _, v := range users {
		if v.Id == id {
			u = v
			err = nil
			break
		}
	}
	return u, err
}

func (s *Service) CreateUser(u user.User) (user.User, error) {
	return s.repo.CreateUser(u)
}

func (s *Service) UpdateUser(u user.User) (user.User, error) {
	return s.repo.UpdateUser(u)
}

func (s *Service) Authenticate(email string, password string) (user.User, error) {
	var u, ok = s.UserByEmail(email, true)
	if !ok {
		return user.User{}, user.ErrNotFound
	}

	if u.Deleted {
		return user.User{}, user.ErrUserDeleted
	}

	if !authUtils.CheckPasswordHash(password, u.PasswordHash) {
		return user.User{}, user.ErrAuthentication
	}

	return u, nil
}

func (s *Service) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
