package treatmentConverter

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/pet"
	"github.com/scarlettmiss/bestPal/application/domain/treatment"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	treatmentType "github.com/scarlettmiss/bestPal/cmd/server/types/treatment"
	"github.com/scarlettmiss/bestPal/converters/userConverter"
	"github.com/scarlettmiss/bestPal/utils"
	"time"
)

func TreatmentCreateRequestToTreatment(requestBody treatmentType.TreatmentCreateRequest, petId uuid.UUID, administeredBy uuid.UUID) (treatment.Treatment, error) {
	t := treatment.Nil

	typ, err := treatment.ParseType(requestBody.TreatmentType)
	if err != nil {
		return t, err
	}

	if utils.TextIsEmpty(requestBody.Name) {
		return t, treatment.ErrNotValidName
	}
	if requestBody.Date == 0 {
		return t, treatment.ErrNotValidDate
	}

	t.PetId = petId
	t.TreatmentType = typ
	t.Name = requestBody.Name
	t.Date = time.Unix(requestBody.Date/1000, (requestBody.Date%1000)*1000000)
	t.Lot = requestBody.Lot
	t.Result = requestBody.Result
	t.Description = requestBody.Description
	t.Notes = requestBody.Notes
	t.AdministeredBy = administeredBy
	if requestBody.NextDate != 0 {
		t.NextDate = time.Unix(requestBody.NextDate/1000, (requestBody.NextDate%1000)*1000000)
	}

	return t, nil
}

func TreatmentUpdateRequestToTreatment(requestBody treatmentType.TreatmentUpdateRequest, t treatment.Treatment) (treatment.Treatment, error) {
	typ, err := treatment.ParseType(requestBody.TreatmentType)
	if err != nil {
		return t, err
	}

	if utils.TextIsEmpty(requestBody.Name) {
		return t, treatment.ErrNotValidName
	}
	if requestBody.Date == 0 {
		return t, treatment.ErrNotValidDate
	}

	t.TreatmentType = typ
	t.Name = requestBody.Name
	t.Date = time.Unix(requestBody.Date/1000, (requestBody.Date%1000)*1000000)
	t.Lot = requestBody.Lot
	t.Result = requestBody.Result
	t.Description = requestBody.Description
	t.Notes = requestBody.Notes
	if requestBody.NextDate != 0 {
		t.NextDate = time.Unix(requestBody.NextDate/1000, (requestBody.NextDate%1000)*1000000)
	} else {
		t.NextDate = time.Time{}
	}
	return t, nil
}

func TreatmentToResponse(t treatment.Treatment, pet pet.Pet, administeredBy user.User, verifiedBy user.User) treatmentType.TreatmentResponse {
	resp := treatmentType.TreatmentResponse{}
	resp.Id = t.Id.String()
	resp.CreatedAt = t.CreatedAt
	resp.UpdatedAt = t.UpdatedAt
	resp.Deleted = t.Deleted
	resp.Pet = pet
	resp.TreatmentType = string(t.TreatmentType)
	resp.Name = t.Name
	resp.Date = t.Date
	resp.Lot = t.Lot
	resp.Result = t.Result
	resp.Description = t.Description
	resp.Notes = t.Notes
	resp.AdministeredBy = userConverter.UserToResponse(administeredBy)
	if verifiedBy != user.Nil {
		resp.VerifiedBy = userConverter.UserToResponse(verifiedBy)
	}
	if !t.NextDate.IsZero() {
		resp.NextDate = t.NextDate.String()
	}

	return resp
}
