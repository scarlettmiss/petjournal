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

func (s *Service) Pets() ([]pet.Pet, error) {
	return s.repo.Pets()
}

func (s *Service) PetsByUser(uId uuid.UUID) (map[uuid.UUID]pet.Pet, error) {
	uPets := make(map[uuid.UUID]pet.Pet)

	pets, err := s.Pets()
	if err != nil {
		return uPets, err
	}

	for _, p := range pets {
		if p.OwnerId == uId || p.VetId == uId {
			uPets[p.Id] = p
		}
	}

	return uPets, nil
}

func (s *Service) PetsByOwner(uId uuid.UUID) (map[uuid.UUID]pet.Pet, error) {
	uPets := make(map[uuid.UUID]pet.Pet)

	pets, err := s.Pets()
	if err != nil {
		return uPets, err
	}

	for _, p := range pets {
		if p.OwnerId == uId {
			uPets[p.Id] = p
		}
	}

	return uPets, nil
}

func (s *Service) PetByUser(uId uuid.UUID, id uuid.UUID) (pet.Pet, error) {
	pets, err := s.PetsByUser(uId)
	if err != nil {
		return pet.Nil, err
	}

	p, ok := pets[id]
	if !ok {
		return pet.Nil, pet.ErrNotFound
	}

	return p, nil
}

func (s *Service) PetByOwner(uId uuid.UUID, id uuid.UUID) (pet.Pet, error) {
	pets, err := s.PetsByOwner(uId)
	if err != nil {
		return pet.Nil, err
	}

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

func (s *Service) RemoveVet(id uuid.UUID) error {
	p, err := s.Pet(id)

	if err != nil {
		return err
	}
	p.VetId = uuid.Nil

	_, err = s.repo.UpdatePet(p)
	return err
}
