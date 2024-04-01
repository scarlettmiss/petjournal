package record

import (
	"errors"
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application"
)

var (
	// ErrNotFound is returned when a record is not found
	ErrNotFound         = errors.New("record not found")
	ErrNotValidName     = errors.New("record name not valid")
	ErrNotValidResult   = errors.New("record result not valid")
	ErrNotValidDate     = errors.New("record date not valid")
	ErrNotValidType     = errors.New("record type not valid")
	ErrNotValidVerifier = errors.New("record cannot be validated by this user")
)

type Service interface {
	record(id uuid.UUID) (Record, error)
	PetsRecords(pIds []uuid.UUID, includeDel bool) (map[uuid.UUID]Record, error)
	PetRecords(pId uuid.UUID, includeDel bool) (map[uuid.UUID]Record, error)
	PetRecord(pId uuid.UUID, rId uuid.UUID, includeDel bool) (Record, error)
	CreateRecord(opts application.UserCreateOptions) (Record, error)
	CreateRecords(opts application.UserCreateOptions) (map[uuid.UUID]Record, error)
	UpdateRecord(record Record) (Record, error)
	DeleteRecord(id uuid.UUID) error
}

type Repository interface {
	CreateRecord(record Record) (Record, error)
	CreateRecords(records []Record) ([]Record, error)
	Record(id uuid.UUID) (Record, error)
	Records(includeDel bool) ([]Record, error)
	UpdateRecord(record Record) (Record, error)
	DeleteRecord(id uuid.UUID) error
}
