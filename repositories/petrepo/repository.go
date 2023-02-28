package petrepo

import (
	"github.com/samber/lo"
	"github.com/scarlettmiss/bestPal/application/domain/pet"
	"sync"
)

type Repository struct {
	mux  sync.Mutex
	pets map[string]pet.Pet
}

func New() *Repository {
	return &Repository{
		pets: map[string]pet.Pet{},
	}
}

func (r *Repository) CreatePet(p pet.Pet) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.pets[p.Id] = p

	return nil
}

func (r *Repository) Pet(id string) (pet.Pet, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	u, ok := r.pets[id]
	if !ok {
		return pet.Nil, pet.ErrNotFound
	}

	return u, nil
}

func (r *Repository) Pets() map[string]pet.Pet {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.pets
}

func (r *Repository) PetsByUser(uId string) map[string]pet.Pet {
	r.mux.Lock()
	defer r.mux.Unlock()
	pets := lo.PickBy[string, pet.Pet](r.pets, func(key string, value pet.Pet) bool {
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

func (r *Repository) DeletePet(id string) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	_, ok := r.pets[id]
	if !ok {
		return pet.ErrNotFound
	}
	delete(r.pets, id)

	return nil
}
