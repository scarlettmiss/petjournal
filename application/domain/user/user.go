package user

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Type string

const (
	Vet   Type = "vet"
	Owner Type = "owner"
)

var types = map[Type]Type{
	Vet:   Vet,
	Owner: Owner,
}

func ParseType(value string) (Type, error) {
	value = strings.TrimSpace(strings.ToLower(value))
	typ, ok := types[Type(value)]
	if !ok {
		return Owner, errors.New("type not found")
	}
	return typ, nil
}

type User struct {
	Id           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Deleted      bool
	UserType     Type
	Email        string
	PasswordHash string
	Name         string
	Surname      string
	Phone        string
	Address      string
	City         string
	State        string
	Country      string
	Zip          string
	VetId        uuid.UUID
}

var Nil = User{}
