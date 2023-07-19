package treatment

import (
	"errors"
	"github.com/google/uuid"
)

var (
	// ErrNotFound is returned when a treatment is not found
	ErrNotFound     = errors.New("treatment not found")
	ErrNotValidName = errors.New("treatment name not valid")
	ErrNotValidDate = errors.New("treatment date not valid")
)

type Service interface {
	Treatment(id uuid.UUID) (Treatment, error)
	Treatments() map[uuid.UUID]Treatment
	PetTreatment(petId uuid.UUID, id uuid.UUID) (Treatment, error)
	PetTreatments(petId uuid.UUID) (map[uuid.UUID]Treatment, error)
	CreateTreatment(treatment Treatment) (Treatment, error)
	UpdateTreatment(treatment Treatment) (Treatment, error)
	DeleteTreatment(id uuid.UUID) error
}

type Repository interface {
	CreateTreatment(treatment Treatment) (Treatment, error)
	Treatment(id uuid.UUID) (Treatment, error)
	Treatments() ([]Treatment, error)
	UpdateTreatment(treatment Treatment) (Treatment, error)
	DeleteTreatment(id uuid.UUID) error
}
