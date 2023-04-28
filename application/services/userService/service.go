package service

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	authService "github.com/scarlettmiss/bestPal/application/services/authService"
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

func (s *Service) Users() map[uuid.UUID]user.User {
	return s.repo.Users()
}

func (s *Service) UserByEmail(email string) (user.User, bool) {
	var u user.User
	var found bool

	for _, v := range s.Users() {
		if v.Email == email {
			u = v
			found = true
			break
		}
	}
	return u, found
}

func (s *Service) CreateUser(u user.User) (user.User, error) {
	return s.repo.CreateUser(u)
}

func (s *Service) UpdateUser(u user.User) (user.User, error) {
	return s.repo.UpdateUser(u)
}

func (s *Service) Authenticate(email string, password string) (user.User, error) {

	var u, ok = s.UserByEmail(email)
	if !ok {
		return user.User{}, user.ErrNotFound
	}

	if !authService.CheckPasswordHash(password, u.PasswordHash) {
		return user.User{}, user.ErrAuthentication
	}

	return u, nil
}

func (s *Service) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
