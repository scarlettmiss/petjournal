package record

import (
	"errors"
	"github.com/google/uuid"
)

var (
	// ErrNotFound is returned when a Record is not found
	ErrNotFound     = errors.New("record not found")
	ErrNotValidName = errors.New("record name not valid")
	ErrNotValidDate = errors.New("record date not valid")
)

type Service interface {
	Record(id uuid.UUID) (Record, error)
	Records() map[uuid.UUID]Record
	PetRecord(petId uuid.UUID, id uuid.UUID) (Record, error)
	PetRecords(petId uuid.UUID) (map[uuid.UUID]Record, error)
	CreateRecord(record Record) (Record, error)
	UpdateRecord(record Record) (Record, error)
	DeleteRecord(id uuid.UUID) error
}

type Repository interface {
	CreateRecord(record Record) (Record, error)
	Record(id uuid.UUID) (Record, error)
	Records() ([]Record, error)
	UpdateRecord(record Record) (Record, error)
	DeleteRecord(id uuid.UUID) error
}
