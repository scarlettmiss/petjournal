package service

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application"
	"github.com/scarlettmiss/petJournal/application/domain/pet"
	textUtils "github.com/scarlettmiss/petJournal/utils/text"
)

type Service struct {
	repo pet.Repository
}

func New(repo pet.Repository) (Service, error) {
	return Service{repo: repo}, nil
}

func (s *Service) Pet(id uuid.UUID) (pet.Pet, error) {
	return s.repo.Pet(id)
}

func (s *Service) PetByUser(uId uuid.UUID, id uuid.UUID, includeDel bool) (pet.Pet, error) {
	pets, err := s.PetsByUser(uId, includeDel)
	if err != nil {
		return pet.Nil, err
	}

	p, ok := pets[id]
	if !ok {
		return pet.Nil, pet.ErrNotFound
	}

	return p, nil
}

func (s *Service) Pets(includeDel bool) ([]pet.Pet, error) {
	return s.repo.Pets(includeDel)
}

func (s *Service) PetsByUser(uId uuid.UUID, includeDel bool) (map[uuid.UUID]pet.Pet, error) {
	uPets := make(map[uuid.UUID]pet.Pet)

	pets, err := s.Pets(includeDel)
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

func (s *Service) CreatePet(opts application.PetCreateOptions) (pet.Pet, error) {
	if textUtils.TextIsEmpty(opts.Name) {
		return pet.Nil, pet.ErrNoValidName
	}

	if opts.DateOfBirth.IsZero() {
		return pet.Nil, pet.ErrNoValidBirthDate
	}

	gender, err := pet.ParseGender(opts.Gender)
	if err != nil {
		return pet.Nil, err
	}

	if textUtils.TextIsEmpty(opts.BreedName) {
		return pet.Nil, pet.ErrNoValidBreedname
	}

	p := pet.Pet{}
	p.Name = opts.Name
	p.DateOfBirth = opts.DateOfBirth
	p.Gender = gender
	p.BreedName = opts.BreedName
	p.Colors = opts.Colors
	p.Description = opts.Description
	p.Pedigree = opts.Pedigree
	p.Microchip = opts.Microchip
	p.OwnerId = opts.OwnerId
	p.VetId = opts.VetId
	p.Metas = opts.Metas
	p.Avatar = opts.Avatar
	return s.repo.CreatePet(p)
}

func (s *Service) UpdatePet(opts application.PetUpdateOptions) (pet.Pet, error) {
	p, err := s.PetByUser(opts.OwnerId, opts.Id, false)
	if err != nil {
		return pet.Nil, err
	}

	if textUtils.TextIsEmpty(opts.Name) {
		return pet.Nil, pet.ErrNoValidName
	}

	if textUtils.TextIsEmpty(opts.BreedName) {
		return pet.Nil, pet.ErrNoValidBreedname
	}

	if opts.DateOfBirth.IsZero() {
		return pet.Nil, pet.ErrNoValidBirthDate
	}

	gender, err := pet.ParseGender(opts.Gender)
	if err != nil {
		return pet.Nil, err
	}

	p.Name = opts.Name
	p.DateOfBirth = opts.DateOfBirth
	p.Gender = gender
	p.BreedName = opts.BreedName
	p.Colors = opts.Colors
	p.Description = opts.Description
	p.Pedigree = opts.Pedigree
	p.Microchip = opts.Microchip
	p.VetId = opts.VetId
	p.Metas = opts.Metas
	p.Avatar = opts.Avatar

	return s.repo.UpdatePet(p)
}

func (s *Service) DeletePet(uId uuid.UUID, id uuid.UUID) error {
	_, err := s.petByOwner(uId, id, true)

	if err != nil {
		// if it's not the owner then check if it's the pet vet
		// in this case we don't want to delete the pet but we want to remove the vet
		if err == pet.ErrNotFound {
			return s.removeVet(id)
		}
		return err
	}

	return s.repo.DeletePet(id)
}

func (s *Service) removeVet(id uuid.UUID) error {
	p, err := s.Pet(id)

	if err != nil {
		return err
	}
	p.VetId = uuid.Nil

	_, err = s.repo.UpdatePet(p)
	return err
}

func (s *Service) petsByOwner(uId uuid.UUID, includeDel bool) (map[uuid.UUID]pet.Pet, error) {
	uPets := make(map[uuid.UUID]pet.Pet)

	pets, err := s.Pets(includeDel)
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

func (s *Service) petByOwner(uId uuid.UUID, id uuid.UUID, includeDel bool) (pet.Pet, error) {
	pets, err := s.petsByOwner(uId, includeDel)
	if err != nil {
		return pet.Nil, err
	}

	p, ok := pets[id]
	if !ok {
		return pet.Nil, pet.ErrNotFound
	}

	return p, nil
}
