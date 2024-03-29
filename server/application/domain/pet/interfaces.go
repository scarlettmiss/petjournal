package pet

import (
	"errors"
	"github.com/google/uuid"
)

var (
	// ErrNotFound is returned when a pet is not found
	ErrNotFound         = errors.New("pet not found")
	ErrNoValidName      = errors.New("a valid name should be provided")
	ErrNoValidBreedname = errors.New("a valid breed should be provided")
	ErrNoValidBirthDate = errors.New("a valid birthdate should be provided")
)

type Service interface {
	Pet(id uuid.UUID) (Pet, error)
	PetByUser(uid uuid.UUID, id uuid.UUID, includeDel bool) (Pet, error)
	PetByOwner(uid uuid.UUID, id uuid.UUID, includeDel bool) (Pet, error)
	Pets(includeDel bool) []Pet
	PetsByUser(userId uuid.UUID, includeDel bool) (map[uuid.UUID]Pet, error)
	PetsByOwner(userId uuid.UUID, includeDel bool) (map[uuid.UUID]Pet, error)
	CreatePet(opts Pet) error
	UpdatePet() error
	DeletePet(id uuid.UUID) error
	RemoveVet(id uuid.UUID) error
}

type Repository interface {
	CreatePet(pet Pet) (Pet, error)
	Pet(id uuid.UUID) (Pet, error)
	Pets(includeDel bool) ([]Pet, error)
	UpdatePet(pet Pet) (Pet, error)
	DeletePet(id uuid.UUID) error
}
