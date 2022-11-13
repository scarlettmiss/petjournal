package user

import "github.com/scarlettmiss/bestPal/application/domain/baseStruct"

type Type string

const (
	Vet   Type = Type("VET")
	Owner Type = Type("OWNER")
)

type User struct {
	baseStruct.BaseStruct
	userType Type
	email    string
	password string
	name     string
	surname  string
	phone    string
	address  string
	country  string
}
