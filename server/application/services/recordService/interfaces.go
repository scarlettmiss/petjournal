package service

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application/domain/record"
	"github.com/scarlettmiss/petJournal/application/services"
)

type Service interface {
	record(id uuid.UUID) (record.Record, error)
	PetsRecords(pIds []uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error)
	PetRecords(pId uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error)
	PetRecord(pId uuid.UUID, rId uuid.UUID, includeDel bool) (record.Record, error)
	CreateRecord(opts services.RecordCreateOptions) (record.Record, error)
	CreateRecords(opts services.RecordsCreateOptions) (map[uuid.UUID]record.Record, error)
	UpdateRecord(opts services.RecordUpdateOptions) (record.Record, error)
	DeleteRecord(id uuid.UUID) error
}

type Repository interface {
	CreateRecord(record record.Record) (record.Record, error)
	CreateRecords(records []record.Record) ([]record.Record, error)
	Record(id uuid.UUID) (record.Record, error)
	Records(includeDel bool) ([]record.Record, error)
	UpdateRecord(record record.Record) (record.Record, error)
	DeleteRecord(id uuid.UUID) error
}
