package types

import (
	"github.com/scarlettmiss/bestPal/application/domain/user"
)

type Account struct {
	UserType user.Type `json:"userType,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Name     string    `json:"name,omitempty"`
	Surname  string    `json:"surname,omitempty"`
	Phone    string    `json:"phone,omitempty"`
	Address  string    `json:"address,omitempty"`
	Country  string    `json:"country,omitempty"`
}
