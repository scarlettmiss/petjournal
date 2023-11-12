package recordConverter

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application/domain/pet"
	"github.com/scarlettmiss/petJournal/application/domain/record"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	recordType "github.com/scarlettmiss/petJournal/cmd/server/types/record"
	"github.com/scarlettmiss/petJournal/converters/petConverter"
	"github.com/scarlettmiss/petJournal/converters/userConverter"
	"time"
)

func RecordCreateRequestToRecord(requestBody recordType.RecordCreateRequest, petId uuid.UUID, administeredBy uuid.UUID) (record.Record, error) {
	r := record.Nil

	typ, err := record.ParseType(requestBody.RecordType)
	if err != nil {
		return r, err
	}

	if requestBody.Date == 0 {
		return r, record.ErrNotValidDate
	}

	r.PetId = petId
	r.RecordType = typ
	r.Name = requestBody.Name
	r.Date = time.Unix(requestBody.Date/1000, (requestBody.Date%1000)*1000000)
	r.Lot = requestBody.Lot
	r.Result = requestBody.Result
	r.Description = requestBody.Description
	r.Notes = requestBody.Notes
	r.AdministeredBy = administeredBy
	if requestBody.NextDate != 0 {
		r.NextDate = time.Unix(requestBody.NextDate/1000, (requestBody.NextDate%1000)*1000000)
	}

	return r, nil
}

func RecordUpdateRequestToRecord(requestBody recordType.RecordUpdateRequest, r record.Record) (record.Record, error) {
	typ, err := record.ParseType(requestBody.RecordType)
	if err != nil {
		return r, err
	}

	if requestBody.Date == 0 {
		return r, record.ErrNotValidDate
	}

	r.RecordType = typ
	r.Name = requestBody.Name
	r.Date = time.Unix(requestBody.Date/1000, (requestBody.Date%1000)*1000000)
	r.Lot = requestBody.Lot
	r.Result = requestBody.Result
	r.Description = requestBody.Description
	r.Notes = requestBody.Notes
	if requestBody.NextDate != 0 {
		r.NextDate = time.Unix(requestBody.NextDate/1000, (requestBody.NextDate%1000)*1000000)
	} else {
		r.NextDate = time.Time{}
	}
	return r, nil
}

func RecordToResponse(r record.Record, pet pet.Pet, administeredBy user.User, verifiedBy user.User) recordType.RecordResponse {
	resp := recordType.RecordResponse{}
	resp.Id = r.Id.String()
	resp.CreatedAt = r.CreatedAt.UnixMilli()
	resp.UpdatedAt = r.UpdatedAt.UnixMilli()
	resp.Deleted = r.Deleted
	resp.Pet = petConverter.PetToVerySimplifiedResponse(pet)
	resp.RecordType = string(r.RecordType)
	resp.Name = r.Name
	resp.Date = r.Date.UnixMilli()
	resp.Lot = r.Lot
	resp.Result = r.Result
	resp.Description = r.Description
	resp.Notes = r.Notes
	resp.AdministeredBy = userConverter.UserToResponse(administeredBy)
	if verifiedBy != user.Nil {
		resp.VerifiedBy = userConverter.UserToResponse(verifiedBy)
	}
	if !r.NextDate.IsZero() {
		resp.NextDate = r.NextDate.UnixMilli()
	}

	return resp
}
