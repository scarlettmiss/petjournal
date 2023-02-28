package treatment

import "errors"

var (
	// ErrNotFound is returned when a treatment is not found
	ErrNotFound = errors.New("treatment not found")
)

type Service interface {
	Treatment(id string) (*Treatment, error)
	PetTreatments(petId string) map[string]*Treatment
	Treatments() map[string]*Treatment
	CreateTreatment(treatment Treatment) error
	UpdateTreatment(treatment Treatment) error
	DeleteTreatment(id string) error
}

type Repository interface {
	CreateTreatment(treatment Treatment) error
	Treatment(id string) (Treatment, error)
	Treatments() map[string]Treatment
	TreatmentsByPet(petId string) map[string]Treatment
	UpdateTreatment(treatment Treatment) error
	DeleteTreatment(id string) error
}
