package pet

import (
	"errors"
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application"
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
	Pets(includeDel bool) []Pet
	PetsByUser(userId uuid.UUID, includeDel bool) (map[uuid.UUID]Pet, error)
	CreatePet(opts application.PetCreateOptions) (Pet, error)
	UpdatePet(opts application.PetUpdateOptions) (Pet, error)
	DeletePet(uId uuid.UUID, id uuid.UUID) error
	removeVet(id uuid.UUID) error
	PetsByOwner(userId uuid.UUID, includeDel bool) (map[uuid.UUID]Pet, error)
	petByOwner(uid uuid.UUID, id uuid.UUID, includeDel bool) (Pet, error)
}

type Repository interface {
	CreatePet(pet Pet) (Pet, error)
	Pet(id uuid.UUID) (Pet, error)
	Pets(includeDel bool) ([]Pet, error)
	UpdatePet(pet Pet) (Pet, error)
	DeletePet(id uuid.UUID) error
}
