package treatment

import (
	"errors"
	"github.com/google/uuid"
)

var (
	// ErrNotFound is returned when a treatment is not found
	ErrNotFound = errors.New("treatment not found")
)

type Service interface {
	Treatment(id uuid.UUID) (*Treatment, error)
	PetTreatments(petId uuid.UUID) map[uuid.UUID]Treatment
	Treatments() map[uuid.UUID]*Treatment
	CreateTreatment(treatment Treatment) error
	UpdateTreatment(treatment Treatment) error
	DeleteTreatment(id uuid.UUID) error
}

type Repository interface {
	CreateTreatment(treatment Treatment) error
	Treatment(id uuid.UUID) (Treatment, error)
	Treatments() map[uuid.UUID]Treatment
	TreatmentsByPet(petId uuid.UUID) map[uuid.UUID]Treatment
	UpdateTreatment(treatment Treatment) error
	DeleteTreatment(id uuid.UUID) error
}
