package userConverter

import (
	"github.com/scarlettmiss/petJournal/application/domain/user"
	authService "github.com/scarlettmiss/petJournal/application/services/authService"
	user2 "github.com/scarlettmiss/petJournal/cmd/server/types/user"
	"github.com/scarlettmiss/petJournal/utils"
)

func UserCreateRequestToUser(requestBody user2.UserCreateRequest) (user.User, error) {
	u := user.Nil
	typ, err := user.ParseType(requestBody.UserType)
	if err != nil {
		return u, err
	}

	if !utils.IsEmailValid(requestBody.Email) {
		return u, user.ErrNoValidMail
	}

	err = utils.IsPasswordValid(requestBody.Password)
	if err != nil {
		return u, err
	}
	hashed, err := authService.HashPassword(requestBody.Password)
	if err != nil {
		return u, err
	}

	if utils.TextIsEmpty(requestBody.Name) {
		return u, user.ErrNoValidName
	}
	if utils.TextIsEmpty(requestBody.Surname) {
		return u, user.ErrNoValidSurname
	}

	u.UserType = typ
	u.Email = requestBody.Email
	u.PasswordHash = hashed
	u.Name = requestBody.Name
	u.Surname = requestBody.Surname
	u.Phone = requestBody.Phone
	u.Address = requestBody.Address
	u.City = requestBody.City
	u.State = requestBody.State
	u.Country = requestBody.Country
	u.Zip = requestBody.Zip
	return u, nil
}

func UserUpdateRequestToUser(requestBody user2.UserUpdateRequest, u user.User) (user.User, error) {
	if utils.IsEmailValid(requestBody.Email) {
		return u, user.ErrNoValidMail
	}
	if utils.TextIsEmpty(requestBody.Name) {
		return u, user.ErrNoValidName
	}
	if utils.TextIsEmpty(requestBody.Surname) {
		return u, user.ErrNoValidSurname
	}

	u.Email = requestBody.Email
	u.Name = requestBody.Name
	u.Surname = requestBody.Surname
	u.Phone = requestBody.Phone
	u.Address = requestBody.Address
	u.City = requestBody.City
	u.State = requestBody.State
	u.Country = requestBody.Country
	u.Zip = requestBody.Zip
	return u, nil
}

func UserToResponse(u user.User) user2.UserResponse {
	resp := user2.UserResponse{}
	resp.Id = u.Id.String()
	resp.CreatedAt = u.CreatedAt.UnixMilli()
	resp.UpdatedAt = u.UpdatedAt.UnixMilli()
	resp.Deleted = u.Deleted
	resp.UserType = u.UserType
	resp.Email = u.Email
	resp.Name = u.Name
	resp.Surname = u.Surname
	resp.Phone = u.Phone
	resp.Address = u.Address
	resp.City = u.City
	resp.State = u.State
	resp.Country = u.Country
	resp.Zip = u.Zip
	return resp
}

func UserToSimplifiedResponse(u user.User) user2.UserResponse {
	resp := user2.UserResponse{}
	resp.Id = u.Id.String()
	resp.Deleted = u.Deleted
	resp.UserType = u.UserType
	resp.Email = u.Email
	resp.Name = u.Name
	resp.Surname = u.Surname
	resp.Phone = u.Phone
	return resp
}
