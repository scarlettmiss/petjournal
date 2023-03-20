package treatmentrepo

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/scarlettmiss/bestPal/application/domain/treatment"
	"sync"
)

type Repository struct {
	mux        sync.Mutex
	treatments map[uuid.UUID]treatment.Treatment
}

func New() *Repository {
	return &Repository{
		treatments: map[uuid.UUID]treatment.Treatment{},
	}
}

func (r *Repository) CreateTreatment(t treatment.Treatment) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.treatments[t.Id] = t

	return nil
}

func (r *Repository) Treatment(id uuid.UUID) (treatment.Treatment, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	t, ok := r.treatments[id]
	if !ok {
		return treatment.Nil, treatment.ErrNotFound
	}

	return t, nil
}

func (r *Repository) Treatments() map[uuid.UUID]treatment.Treatment {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.treatments
}

func (r *Repository) TreatmentsByPet(pId uuid.UUID) map[uuid.UUID]treatment.Treatment {
	r.mux.Lock()
	defer r.mux.Unlock()
	treatments := lo.PickBy[uuid.UUID, treatment.Treatment](r.treatments, func(key uuid.UUID, value treatment.Treatment) bool {
		return value.Id == pId
	})

	return treatments
}

func (r *Repository) UpdateTreatment(t treatment.Treatment) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	_, ok := r.treatments[t.Id]
	if !ok {
		return treatment.ErrNotFound
	}

	r.treatments[t.Id] = t

	return nil
}

func (r *Repository) DeleteTreatment(id uuid.UUID) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	_, ok := r.treatments[id]
	if !ok {
		return treatment.ErrNotFound
	}
	delete(r.treatments, id)

	return nil
}
