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
	PetByUser(uid uuid.UUID, id uuid.UUID) (Pet, error)
	Pets(userId uuid.UUID) map[uuid.UUID]Pet
	CreatePet(pet Pet) error
	UpdatePet() error
	PetsByUser(userId uuid.UUID) map[uuid.UUID]Pet
	DeletePet(id uuid.UUID) error
}

type Repository interface {
	CreatePet(pet Pet) (Pet, error)
	Pet(id uuid.UUID) (Pet, error)
	Pets() map[uuid.UUID]Pet
	UpdatePet(pet Pet) (Pet, error)
	DeletePet(id uuid.UUID) error
}
