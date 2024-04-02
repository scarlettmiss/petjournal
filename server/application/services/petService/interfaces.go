package service

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application/domain/pet"
	"github.com/scarlettmiss/petJournal/application/services"
)

type Service interface {
	Pet(id uuid.UUID) (pet.Pet, error)
	PetByUser(uid uuid.UUID, id uuid.UUID, includeDel bool) (pet.Pet, error)
	Pets(includeDel bool) ([]pet.Pet, error)
	PetsByUser(userId uuid.UUID, includeDel bool) (map[uuid.UUID]pet.Pet, error)
	CreatePet(opts services.PetCreateOptions) (pet.Pet, error)
	UpdatePet(opts services.PetUpdateOptions) (pet.Pet, error)
	DeletePet(uId uuid.UUID, id uuid.UUID) error
	removeVet(id uuid.UUID) error
	petsByOwner(userId uuid.UUID, includeDel bool) (map[uuid.UUID]pet.Pet, error)
	petByOwner(uid uuid.UUID, id uuid.UUID, includeDel bool) (pet.Pet, error)
}

type Repository interface {
	CreatePet(pet pet.Pet) (pet.Pet, error)
	Pet(id uuid.UUID) (pet.Pet, error)
	Pets(includeDel bool) ([]pet.Pet, error)
	UpdatePet(pet pet.Pet) (pet.Pet, error)
	DeletePet(id uuid.UUID) error
}
