package user

import (
	"github.com/scarlettmiss/bestPal/application/domain/user"
)

type UserResponse struct {
	Id        string    `json:"id"`
	CreatedAt int64     `json:"createdAt,omitempty"`
	UpdatedAt int64     `json:"updatedAt,omitempty"`
	Deleted   bool      `json:"deleted,omitempty"`
	UserType  user.Type `json:"userType,omitempty"`
	Email     string    `json:"email,omitempty"`
	Name      string    `json:"name,omitempty"`
	Surname   string    `json:"surname,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Address   string    `json:"address,omitempty"`
	City      string    `json:"city,omitempty"`
	State     string    `json:"state,omitempty"`
	Country   string    `json:"country,omitempty"`
	Zip       string    `json:"zip,omitempty"`
}
