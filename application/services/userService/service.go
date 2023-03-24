package service

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/base"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	"github.com/scarlettmiss/bestPal/cmd/server/types"
	"time"
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

func (s *Service) CreateUser(a types.Account) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	t := time.Time{}
	base := base.Base{
		Id:        id,
		CreatedAt: t,
		UpdatedAt: t,
		Deleted:   false,
	}
	u := user.User{
		Base:     base,
		UserType: a.UserType,
		Email:    a.Email,
		Password: a.Password,
		Name:     a.Name,
		Surname:  a.Surname,
		Phone:    a.Phone,
		Address:  a.Address,
		City:     a.City,
		State:    a.State,
		Country:  a.Country,
		Zip:      a.Zip,
	}

	return s.repo.CreateUser(u)
}

func (s *Service) UpdateUser(u user.User) error {
	return s.repo.UpdateUser(u)
}

func (s *Service) Authenticate(email string, password string) (user.User, error) {
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

func (s *Service) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}
