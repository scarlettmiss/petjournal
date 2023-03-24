package user

import (
	"github.com/scarlettmiss/bestPal/application/domain/base"
)

type Type string

const (
	Vet   Type = "vet"
	Owner Type = "owner"
)

type User struct {
	base.Base
	UserType Type
	Email    string
	Password string
	Name     string
	Surname  string
	Phone    string
	Address  string
	City     string
	State    string
	Country  string
	Zip      string
}

var Nil = User{}
