package service

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/scarlettmiss/petJournal/application/domain/record"
)

type Service struct {
	repo record.Repository
}

func New(repo record.Repository) (Service, error) {
	return Service{repo: repo}, nil
}

func (s Service) Records(includeDel bool) ([]record.Record, error) {
	return s.repo.Records(includeDel)
}

func (s Service) Record(tId uuid.UUID) (record.Record, error) {
	return s.repo.Record(tId)
}

func (s Service) PetsRecords(pIds []uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error) {
	petRecords := make(map[uuid.UUID]record.Record)

	records, err := s.repo.Records(includeDel)
	if err != nil {
		return petRecords, err
	}

	for _, r := range records {
		lo.ForEach(pIds, func(pId uuid.UUID, _ int) {
			if r.PetId == pId {
				petRecords[r.Id] = r
			}
		})
	}

	return petRecords, nil
}

func (s Service) PetRecords(pId uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error) {
	petRecords := make(map[uuid.UUID]record.Record)

	records, err := s.repo.Records(includeDel)
	if err != nil {
		return petRecords, err
	}

	for _, r := range records {
		if r.PetId == pId {
			petRecords[r.Id] = r
		}
	}

	return petRecords, nil
}

func (s Service) PetRecord(pId uuid.UUID, rId uuid.UUID, includeDel bool) (record.Record, error) {
	petRecords, err := s.PetRecords(pId, includeDel)
	if err != nil {
		return record.Nil, err
	}
	petRecord, ok := petRecords[rId]
	if !ok {
		return record.Nil, record.ErrNotFound
	}

	return petRecord, nil
}

func (s Service) CreateRecord(record record.Record) (record.Record, error) {
	return s.repo.CreateRecord(record)
}

func (s Service) CreateRecords(records []record.Record) (map[uuid.UUID]record.Record, error) {
	recordsMap := make(map[uuid.UUID]record.Record)

	records, err := s.repo.CreateRecords(records)
	if err != nil {
		return recordsMap, err
	}

	for _, r := range records {
		recordsMap[r.Id] = r
	}
	return recordsMap, nil
}

func (s Service) UpdateRecord(record record.Record) (record.Record, error) {
	return s.repo.UpdateRecord(record)
}

func (s Service) DeleteRecord(id uuid.UUID) error {
	return s.repo.DeleteRecord(id)
}
