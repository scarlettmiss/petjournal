package treatmentrepo

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/treatment"
	"sync"
	"time"
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

func (r *Repository) CreateTreatment(t treatment.Treatment) (treatment.Treatment, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	id, err := uuid.NewUUID()
	if err != nil {
		return treatment.Nil, err
	}
	t.Id = id

	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now

	t.Deleted = false

	r.treatments[t.Id] = t

	return t, nil
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

func (r *Repository) UpdateTreatment(t treatment.Treatment) (treatment.Treatment, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	_, ok := r.treatments[t.Id]
	if !ok {
		return treatment.Nil, treatment.ErrNotFound
	}

	now := time.Now()
	t.UpdatedAt = now

	r.treatments[t.Id] = t

	return t, nil
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
