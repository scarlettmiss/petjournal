package converters

import (
	"github.com/scarlettmiss/bestPal/application/domain/user"
	authService "github.com/scarlettmiss/bestPal/application/services/authService"
	"github.com/scarlettmiss/bestPal/cmd/server/types"
)

func UserCreateRequestToUser(requestBody types.UserCreateRequest) (user.User, error) {
	u := user.User{}
	typ, err := user.ParseType(requestBody.UserType)
	if err != nil {
		return user.Nil, err
	}
	u.UserType = typ
	u.Email = requestBody.Email
	hashed, err := authService.HashPassword(requestBody.Password)
	if err != nil {
		return user.Nil, err
	}
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

func UserUpdateRequestToUser(requestBody types.UserUpdateRequest, u user.User) user.User {
	u.Email = requestBody.Email
	u.Name = requestBody.Name
	u.Surname = requestBody.Surname
	u.Phone = requestBody.Phone
	u.Address = requestBody.Address
	u.City = requestBody.City
	u.State = requestBody.State
	u.Country = requestBody.Country
	u.Zip = requestBody.Zip
	return u
}

func UserToResponse(u user.User) types.UserResponse {
	resp := types.UserResponse{}
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
	resp.Deleted = u.Deleted
	return resp
}
