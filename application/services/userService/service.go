package service

import (
	"github.com/scarlettmiss/bestPal/application/domain/user"
)

type Service struct {
	repo user.Repository
}

func New(repo user.Repository) (Service, error) {
	return Service{repo: repo}, nil
}

func (s Service) User(id string) (user.User, error) {
	u, err := s.repo.User(id)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s Service) Users() map[string]user.User {
	return s.repo.Users()
}

func (s Service) CreateUser(u user.User) error {
	return s.repo.CreateUser(u)
}

func (s Service) UpdateUser(u user.User) error {
	return s.repo.UpdateUser(u)
}

func (s Service) Authenticate(email string, password string) (user.User, error) {
	users := s.Users()

	var userByEmail user.User
	var found bool
	for _, v := range users {
		if v.Email == email {
			userByEmail = v
			found = true
			break
		}
	}

	if !found {
		return user.User{}, user.ErrAuthentication
	}

	if userByEmail.Password != password {
		return user.User{}, user.ErrAuthentication
	}

	return userByEmail, nil
}

func (s Service) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
