package service

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/pet"
)

type Service struct {
	repo pet.Repository
}

func New(repo pet.Repository) (Service, error) {
	return Service{repo: repo}, nil
}

func (s *Service) Pet(id uuid.UUID) (pet.Pet, error) {
	p, err := s.repo.Pet(id)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (s *Service) PetsByUser(uId uuid.UUID) map[uuid.UUID]pet.Pet {
	pets := s.repo.Pets()
	uPets := make(map[uuid.UUID]pet.Pet)

	for petId, p := range pets {
		if p.OwnerId == uId || p.VetId == uId {
			uPets[petId] = p
		}
	}

	return uPets
}

func (s *Service) PetByUser(uId uuid.UUID, id uuid.UUID) (pet.Pet, error) {
	pets := s.PetsByUser(uId)
	p, ok := pets[id]
	if !ok {
		return pet.Nil, pet.ErrNotFound
	}

	return p, nil
}

func (s *Service) CreatePet(p pet.Pet) (pet.Pet, error) {
	return s.repo.CreatePet(p)
}

func (s *Service) UpdatePet(p pet.Pet) (pet.Pet, error) {
	return s.repo.UpdatePet(p)
}

func (s *Service) DeletePet(id uuid.UUID) error {
	return s.repo.DeletePet(id)
}
