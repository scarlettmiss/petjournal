package pet

import (
	"errors"
	"github.com/google/uuid"
)

var (
	// ErrNotFound is returned when a pet is not found
	ErrNotFound = errors.New("pet not found")
)

type Service interface {
	Pet(id uuid.UUID) (Pet, error)
	Pets(userId uuid.UUID) map[uuid.UUID]Pet
	CreatePet(pet Pet) error
	UpdatePetInformation() error
	DeletePet(id uuid.UUID) error
}

type Repository interface {
	CreatePet(pet Pet) error
	Pet(id uuid.UUID) (Pet, error)
	Pets() map[uuid.UUID]Pet
	PetsByUser(userId uuid.UUID) map[uuid.UUID]Pet
	UpdatePet(pet Pet) error
	DeletePet(id uuid.UUID) error
}
