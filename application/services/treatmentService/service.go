package service

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/scarlettmiss/bestPal/application/domain/treatment"
)

type Service struct {
	repo treatment.Repository
}

func New(repo treatment.Repository) (Service, error) {
	return Service{repo: repo}, nil
}

func (s Service) Treatments() map[uuid.UUID]treatment.Treatment {
	return s.repo.Treatments()
}

func (s Service) Treatment(tId uuid.UUID) (treatment.Treatment, error) {
	treatments := s.repo.Treatments()
	t, ok := treatments[tId]
	if !ok {
		return treatment.Nil, treatment.ErrNotFound
	}
	return t, nil
}

func (s Service) PetTreatments(pId uuid.UUID) map[uuid.UUID]treatment.Treatment {
	treatments := s.repo.Treatments()
	petTreatments := lo.PickBy[uuid.UUID, treatment.Treatment](treatments, func(key uuid.UUID, value treatment.Treatment) bool {
		return value.Id == pId
	})

	return petTreatments
}

func (s Service) PetTreatment(pId uuid.UUID, tId uuid.UUID) (treatment.Treatment, error) {
	treatments := s.repo.Treatments()
	petTreatments := lo.PickBy[uuid.UUID, treatment.Treatment](treatments, func(key uuid.UUID, value treatment.Treatment) bool {
		return value.Id == pId
	})

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
