package user

import "github.com/scarlettmiss/bestPal/application/domain/user"

type UserResponseSimplified struct {
	Id       string    `json:"id,omitempty"`
	UserType user.Type `json:"userType,omitempty"`
	Email    string    `json:"email,omitempty"`
	Name     string    `json:"name,omitempty"`
	Surname  string    `json:"surname,omitempty"`
	Phone    string    `json:"phone,omitempty"`
}
