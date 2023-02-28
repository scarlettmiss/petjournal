package pet

import (
	"errors"
)

var (
	// ErrNotFound is returned when a pet is not found
	ErrNotFound = errors.New("pet not found")
)

type Service interface {
	Pet(id string) (Pet, error)
	Pets(userId string) map[string]Pet
	CreatePet(pet Pet) error
	UpdatePetInformation() error
	DeletePet(id string) error
}

type Repository interface {
	CreatePet(pet Pet) error
	Pet(id string) (Pet, error)
	Pets() map[string]Pet
	PetsByUser(userId string) map[string]Pet
	UpdatePet(pet Pet) error
	DeletePet(id string) error
}
