package petrepo

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/scarlettmiss/bestPal/application/domain/pet"
	"sync"
)

type Repository struct {
	mux  sync.Mutex
	pets map[uuid.UUID]pet.Pet
}

func New() *Repository {
	return &Repository{
		pets: map[uuid.UUID]pet.Pet{},
	}
}

func (r *Repository) CreatePet(p pet.Pet) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.pets[p.Id] = p

	return nil
}

func (r *Repository) Pet(id uuid.UUID) (pet.Pet, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	u, ok := r.pets[id]
	if !ok {
		return pet.Nil, pet.ErrNotFound
	}

	return u, nil
}

func (r *Repository) Pets() map[uuid.UUID]pet.Pet {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.pets
}

func (r *Repository) PetsByUser(uId uuid.UUID) map[uuid.UUID]pet.Pet {
	r.mux.Lock()
	defer r.mux.Unlock()
	pets := lo.PickBy[uuid.UUID, pet.Pet](r.pets, func(key uuid.UUID, value pet.Pet) bool {
		return value.Id == uId
	})

	return pets
}

func (r *Repository) UpdatePet(p pet.Pet) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	_, ok := r.pets[p.Id]
	if !ok {
		return pet.ErrNotFound
	}

	r.pets[p.Id] = p

	return nil
}

func (r *Repository) DeletePet(id uuid.UUID) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	_, ok := r.pets[id]
	if !ok {
		return pet.ErrNotFound
	}
	delete(r.pets, id)

	return nil
}
