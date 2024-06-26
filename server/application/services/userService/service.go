package service

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	"github.com/scarlettmiss/petJournal/application/services"
	"github.com/scarlettmiss/petJournal/repositories/userrepo"
	authUtils "github.com/scarlettmiss/petJournal/utils/authorization"
	jwtUtils "github.com/scarlettmiss/petJournal/utils/jwt"
	textUtils "github.com/scarlettmiss/petJournal/utils/text"
	"regexp"
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

type service struct {
	repo userrepo.Repository
}

func New(repo userrepo.Repository) (Service, error) {
	return service{repo: repo}, nil
}

func (s service) User(id uuid.UUID) (user.User, error) {
	return s.repo.User(id)
}

func (s service) Users(includeDel bool) ([]user.User, error) {
	return s.repo.Users(includeDel)
}

func (s service) UsersByType(t user.Type, includeDel bool) ([]user.User, error) {
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

func (s service) UserByType(id uuid.UUID, t user.Type, includeDel bool) (user.User, error) {
	u := user.Nil

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

func (s service) CreateUser(opts services.UserCreateOptions) (user.User, string, error) {
	u := user.Nil

	typ, err := user.ParseType(opts.UserType)
	if err != nil {
		return u, "", user.ErrNoValidType
	}

	err = s.checkEmail(opts.Email, u.Id, true)
	if err != nil {
		return u, "", err
	}

	err = passwordValidation(opts.Password)
	if err != nil {
		return u, "", err
	}

	hashed, err := authUtils.HashPassword(opts.Password)
	if err != nil {
		return u, "", err
	}

	if textUtils.TextIsEmpty(opts.Name) {
		return u, "", user.ErrNoValidName
	}

	if textUtils.TextIsEmpty(opts.Surname) {
		return u, "", user.ErrNoValidSurname
	}

	u.UserType = typ
	u.Email = opts.Email
	u.PasswordHash = hashed
	u.Name = opts.Name
	u.Surname = opts.Surname
	u.Phone = opts.Phone
	u.Address = opts.Address
	u.City = opts.City
	u.State = opts.State
	u.Country = opts.Country
	u.Zip = opts.Zip

	u, err = s.repo.CreateUser(u)
	if err != nil {
		return u, "", err
	}

	token, err := userToken(u)
	if err != nil {
		return u, token, err
	}

	return u, token, nil
}

func (s service) UpdateUser(opts services.UserUpdateOptions, includeDel bool) (user.User, error) {
	u, err := s.User(opts.Id)
	if err != nil {
		return u, user.ErrNotFound
	}

	err = s.checkEmail(opts.Email, u.Id, includeDel)
	if err != nil {
		return u, err
	}

	if textUtils.TextIsEmpty(opts.Name) {
		return u, user.ErrNoValidName
	}

	if textUtils.TextIsEmpty(opts.Surname) {
		return u, user.ErrNoValidSurname
	}

	u.Email = opts.Email
	u.Name = opts.Name
	u.Surname = opts.Surname
	u.Phone = opts.Phone
	u.Address = opts.Address
	u.City = opts.City
	u.State = opts.State
	u.Country = opts.Country
	u.Zip = opts.Zip

	return s.repo.UpdateUser(u)
}

func (s service) Authenticate(email string, password string) (user.User, string, error) {
	var u, ok = s.userByEmail(email, true)
	if !ok {
		return u, "", user.ErrNotFound
	}

	if u.Deleted {
		return u, "", user.ErrUserDeleted
	}

	if !authUtils.CheckPasswordHash(password, u.PasswordHash) {
		return u, "", user.ErrAuthentication
	}

	token, err := userToken(u)
	if err != nil {
		return u, token, err
	}

	return u, token, nil
}

func (s service) DeleteUser(id uuid.UUID) error {
	return s.repo.DeleteUser(id)
}

func (s service) userByEmail(email string, includeDel bool) (user.User, bool) {
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

func (s service) checkEmail(email string, id uuid.UUID, includeDel bool) error {
	if !textUtils.IsEmailValid(email) {
		return user.ErrNoValidMail
	}

	u, ok := s.userByEmail(email, includeDel)

	if !ok {
		return nil
	}

	if u.Id == id {
		return nil
	}

	return user.ErrMailExists
}

// IsPasswordValid
// Password should be of 8 characters long
// Password should contain at least one lower case character
// Password should contain at least one upper case character
// Password should contain at least one digit
// Password should contain at least one special character
func passwordValidation(password string) error {
	if len(password) < 8 {
		return user.ErrPasswordLength
	}
	done, err := regexp.MatchString("([a-z])+", password)
	if err != nil {
		return err
	}
	if !done {
		return user.ErrPasswordLowerCase
	}
	done, err = regexp.MatchString("([A-Z])+", password)
	if err != nil {
		return err
	}
	if !done {
		return user.ErrPasswordUpperCase
	}
	done, err = regexp.MatchString("([0-9])+", password)
	if err != nil {
		return err
	}
	if !done {
		return user.ErrPasswordDigit
	}

	done, err = regexp.MatchString("([!@#$%^&*.?-])+", password)
	if err != nil {
		return err
	}
	if !done {
		return user.ErrPasswordSpecialChar
	}

	return nil
}

func userToken(u user.User) (string, error) {
	if u.Deleted {
		return "", user.ErrUserDeleted
	}
	return jwtUtils.GenerateJWT(u.Id, u.UserType)
}
