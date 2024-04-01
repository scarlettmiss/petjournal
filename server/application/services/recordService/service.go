package service

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/scarlettmiss/petJournal/application"
	"github.com/scarlettmiss/petJournal/application/domain/record"
	textUtils "github.com/scarlettmiss/petJournal/utils/text"
	"time"
)

type Service struct {
	repo record.Repository
}

func New(repo record.Repository) (Service, error) {
	return Service{repo: repo}, nil
}

func (s *Service) record(tId uuid.UUID) (record.Record, error) {
	return s.repo.Record(tId)
}

func (s *Service) PetsRecords(pIds []uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error) {
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

func (s *Service) PetRecords(pId uuid.UUID, includeDel bool) (map[uuid.UUID]record.Record, error) {
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

func (s *Service) PetRecord(pId uuid.UUID, rId uuid.UUID, includeDel bool) (record.Record, error) {
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

func (s *Service) CreateRecord(opts application.RecordCreateOptions) (record.Record, error) {
	r := record.Nil

	typ, err := record.ParseType(opts.RecordType)
	if err != nil {
		return r, record.ErrNotValidType
	}

	if typ == record.Weight || typ == record.Temperature {
		if textUtils.TextIsEmpty(opts.Result) {
			return record.Nil, record.ErrNotValidResult
		}
		if opts.Date.After(time.Now()) {
			return record.Nil, record.ErrNotValidDate
		}
	} else {
		if textUtils.TextIsEmpty(opts.Name) {
			return record.Nil, record.ErrNotValidName
		}
	}

	if opts.Date.IsZero() {
		return r, record.ErrNotValidDate
	}

	r.PetId = opts.PetId
	r.RecordType = typ
	r.Name = opts.Name
	r.Date = opts.Date
	r.Lot = opts.Lot
	r.Result = opts.Result
	r.Description = opts.Description
	r.Notes = opts.Notes

	if !opts.Date.After(time.Now()) {
		r.AdministeredBy = opts.AdministeredBy
		r.VerifiedBy = opts.VerifiedBy
	}

	return s.repo.CreateRecord(r)
}

func (s *Service) CreateRecords(opts application.RecordsCreateOptions) (map[uuid.UUID]record.Record, error) {
	r := record.Nil

	typ, err := record.ParseType(opts.RecordType)
	if err != nil {
		return nil, record.ErrNotValidType
	}

	if typ == record.Weight || typ == record.Temperature {
		if textUtils.TextIsEmpty(opts.Result) {
			return nil, record.ErrNotValidResult
		}
		if opts.Date.After(time.Now()) {
			return nil, record.ErrNotValidDate
		}
	} else {
		if textUtils.TextIsEmpty(opts.Name) {
			return nil, record.ErrNotValidName
		}
	}

	if opts.Date.IsZero() {
		return nil, record.ErrNotValidDate
	}

	if opts.NextDate.IsZero() {
		return nil, record.ErrNotValidDate
	}

	recs := make([]record.Record, 2)
	r.PetId = opts.PetId
	r.RecordType = typ
	r.Name = opts.Name
	r.Date = opts.Date
	r.Lot = opts.Lot
	r.Result = opts.Result
	r.Description = opts.Description
	r.Notes = opts.Notes

	if !opts.Date.After(time.Now()) {
		r.AdministeredBy = opts.AdministeredBy
		r.VerifiedBy = opts.VerifiedBy
	}
	recs[0] = r

	nextRecord := record.Record{}
	nextRecord.PetId = opts.PetId
	nextRecord.RecordType = typ
	nextRecord.Name = opts.Name
	nextRecord.Date = opts.NextDate
	recs[1] = nextRecord

	recordsMap := make(map[uuid.UUID]record.Record)

	records, err := s.repo.CreateRecords(recs)
	if err != nil {
		return recordsMap, err
	}

	for _, rec := range records {
		recordsMap[rec.Id] = rec
	}
	return recordsMap, nil
}

func (s *Service) UpdateRecord(opts application.RecordUpdateOptions) (record.Record, error) {
	typ, err := record.ParseType(opts.RecordType)
	if err != nil {
		return record.Nil, record.ErrNotValidType
	}

	if typ == record.Weight || typ == record.Temperature {
		if textUtils.TextIsEmpty(opts.Result) {
			return record.Nil, record.ErrNotValidResult
		}
		if opts.Date.After(time.Now()) {
			return record.Nil, record.ErrNotValidDate
		}
	} else {
		if textUtils.TextIsEmpty(opts.Name) {
			return record.Nil, record.ErrNotValidName
		}
	}

	if opts.Date.IsZero() {
		return record.Nil, record.ErrNotValidDate
	}

	r, err := s.record(opts.Id)
	if err != nil {
		return record.Nil, err
	}

	r.RecordType = typ
	r.Name = opts.Name
	r.Date = opts.Date
	r.Lot = opts.Lot
	r.Result = opts.Result
	r.Description = opts.Description
	r.Notes = opts.Notes
	r.VerifiedBy = opts.VerifiedBy
	if r.AdministeredBy == uuid.Nil {
		r.AdministeredBy = opts.AdministeredBy
	}
	if opts.Date.After(time.Now()) {
		r.AdministeredBy = uuid.Nil
		r.VerifiedBy = uuid.Nil
	}

	return s.repo.UpdateRecord(r)
}

func (s *Service) DeleteRecord(id uuid.UUID) error {
	return s.repo.DeleteRecord(id)
}
