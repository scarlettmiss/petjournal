package service

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/treatment"
)

type Service struct {
	repo treatment.Repository
}

func New(repo treatment.Repository) (Service, error) {
	return Service{repo: repo}, nil
}

func (s Service) Treatments() ([]treatment.Treatment, error) {
	return s.repo.Treatments()
}

func (s Service) Treatment(tId uuid.UUID) (treatment.Treatment, error) {
	return s.repo.Treatment(tId)
}

func (s Service) PetTreatments(pId uuid.UUID) (map[uuid.UUID]treatment.Treatment, error) {
	petTreatments := make(map[uuid.UUID]treatment.Treatment)

	treatments, err := s.repo.Treatments()
	if err != nil {
		return petTreatments, err
	}

	for _, t := range treatments {
		if t.PetId == pId {
			petTreatments[t.Id] = t
		}
	}

	return petTreatments, nil
}

func (s Service) PetTreatment(pId uuid.UUID, tId uuid.UUID) (treatment.Treatment, error) {
	petTreatments, err := s.PetTreatments(pId)
	if err != nil {
		return treatment.Nil, err
	}
	petTreatment, ok := petTreatments[tId]
	if !ok {
		return treatment.Nil, treatment.ErrNotFound
	}

	return petTreatment, nil
}

func (s Service) CreateTreatment(treatment treatment.Treatment) (treatment.Treatment, error) {
	return s.repo.CreateTreatment(treatment)
}

func (s Service) UpdateTreatment(treatment treatment.Treatment) (treatment.Treatment, error) {
	return s.repo.UpdateTreatment(treatment)
}

func (s Service) DeleteTreatment(id uuid.UUID) error {
	return s.repo.DeleteTreatment(id)
}
