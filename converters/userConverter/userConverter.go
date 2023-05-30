package userConverter

import (
	"github.com/scarlettmiss/bestPal/application/domain/user"
	authService "github.com/scarlettmiss/bestPal/application/services/authService"
	user2 "github.com/scarlettmiss/bestPal/cmd/server/types/user"
)

func UserCreateRequestToUser(requestBody user2.UserCreateRequest) (user.User, error) {
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

func UserUpdateRequestToUser(requestBody user2.UserUpdateRequest, u user.User) user.User {
	if requestBody.Email != "" {
		u.Email = requestBody.Email
	}
	if requestBody.Name != "" {
		u.Name = requestBody.Name
	}
	if requestBody.Surname != "" {
		u.Surname = requestBody.Surname
	}
	if requestBody.Phone != "" {
		u.Phone = requestBody.Phone
	}
	if requestBody.Address != "" {
		u.Address = requestBody.Address
	}
	if requestBody.City != "" {
		u.City = requestBody.City
	}
	if requestBody.State != "" {
		u.State = requestBody.State
	}
	if requestBody.Country != "" {
		u.Country = requestBody.Country
	}
	if requestBody.Zip != "" {
		u.Zip = requestBody.Zip
	}

	return u
}

func UserToResponse(u user.User) user2.UserResponse {
	resp := user2.UserResponse{}
	resp.Id = u.Id.String()
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

func UserToSimplifiedResponse(u user.User) user2.UserResponseSimplified {
	resp := user2.UserResponseSimplified{}
	resp.Id = u.Id.String()
	resp.UserType = u.UserType
	resp.Email = u.Email
	resp.Name = u.Name
	resp.Surname = u.Surname
	resp.Phone = u.Phone
	return resp
}
