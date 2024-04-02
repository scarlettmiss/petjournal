package services

import (
	"github.com/google/uuid"
	"time"
)

type PetCreateOptions struct {
	OwnerId     uuid.UUID
	VetId       uuid.UUID
	Avatar      string
	Name        string
	DateOfBirth time.Time
	Gender      string
	BreedName   string
	Colors      []string
	Description string
	Pedigree    string
	Microchip   string
	Metas       map[string]string
}

type PetUpdateOptions struct {
	Id          uuid.UUID
	Avatar      string
	Name        string
	DateOfBirth time.Time
	Gender      string
	BreedName   string
	Colors      []string
	Description string
	Pedigree    string
	Microchip   string
	OwnerId     uuid.UUID
	VetId       uuid.UUID
	Metas       map[string]string
}

type RecordCreateOptions struct {
	PetId          uuid.UUID
	RecordType     string
	Name           string
	Date           time.Time
	Lot            string
	Result         string
	Description    string
	Notes          string
	AdministeredBy uuid.UUID
	VerifiedBy     uuid.UUID
}

type RecordsCreateOptions struct {
	PetId          uuid.UUID
	RecordType     string
	Name           string
	Date           time.Time
	Lot            string
	Result         string
	Description    string
	Notes          string
	AdministeredBy uuid.UUID
	VerifiedBy     uuid.UUID
	NextDate       time.Time
}

type RecordUpdateOptions struct {
	Id             uuid.UUID
	RecordType     string
	Name           string
	Date           time.Time
	Lot            string
	Result         string
	Description    string
	Notes          string
	NextDate       time.Time
	VerifiedBy     uuid.UUID
	AdministeredBy uuid.UUID
}

type LoginOptions struct {
	Email    string
	Password string
}

type UserCreateOptions struct {
	UserType string
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

type UserUpdateOptions struct {
	Id      uuid.UUID
	Email   string
	Name    string
	Surname string
	Phone   string
	Address string
	City    string
	State   string
	Country string
	Zip     string
}
