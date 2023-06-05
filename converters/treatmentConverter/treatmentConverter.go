package treatmentConverter

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/treatment"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	treatmentType "github.com/scarlettmiss/bestPal/cmd/server/types/treatment"
	"github.com/scarlettmiss/bestPal/converters/userConverter"
	"github.com/scarlettmiss/bestPal/utils"
	"time"
)

func TreatmentCreateRequestToTreatment(requestBody treatmentType.TreatmentCreateRequest, petId uuid.UUID, administeredBy uuid.UUID) (treatment.Treatment, error) {
	t := treatment.Treatment{}
	t.PetId = petId

	typ, err := treatment.ParseType(requestBody.TreatmentType)
	if err != nil {
		return treatment.Nil, err
	}
	t.TreatmentType = typ

	if utils.TextIsEmpty(requestBody.Name) {
		return treatment.Nil, treatment.ErrNotValidName
	}
	t.Name = requestBody.Name

	if requestBody.Date == 0 {
		return treatment.Nil, treatment.ErrNotValidDate
	}
	t.Date = time.Unix(requestBody.Date/1000, (requestBody.Date%1000)*1000000)

	t.Lot = requestBody.Lot
	t.Result = requestBody.Result
	t.Description = requestBody.Description
	t.Notes = requestBody.Notes
	t.AdministeredBy = administeredBy
	if err != nil {
		return treatment.Nil, err
	}
	t.RecurringRule = requestBody.RecurringRule

	return t, nil
}

func TreatmentUpdateRequestToTreatment(requestBody treatmentType.TreatmentUpdateRequest, t treatment.Treatment) (treatment.Treatment, error) {
	if !utils.TextIsEmpty(requestBody.TreatmentType) {
		typ, err := treatment.ParseType(requestBody.TreatmentType)
		if err != nil {
			return treatment.Nil, err
		}
		t.TreatmentType = typ
	}
	if !utils.TextIsEmpty(requestBody.Name) {
		t.Name = requestBody.Name
	}
	if requestBody.Date != 0 {
		t.Date = time.Unix(requestBody.Date/1000, (requestBody.Date%1000)*1000000)
	}
	if !utils.TextIsEmpty(requestBody.Lot) {
		t.Lot = requestBody.Lot
	}
	if !utils.TextIsEmpty(requestBody.Result) {
		t.Result = requestBody.Result
	}
	if !utils.TextIsEmpty(requestBody.Description) {
		t.Description = requestBody.Description
	}
	if !utils.TextIsEmpty(requestBody.Notes) {
		t.Notes = requestBody.Notes
	}
	if requestBody.RecurringRule != "" {
		t.RecurringRule = requestBody.RecurringRule
	}

	return t, nil
}

func TreatmentToResponse(t treatment.Treatment, administeredBy user.User, verifiedBy user.User) treatmentType.TreatmentResponse {
	resp := treatmentType.TreatmentResponse{}
	resp.Id = t.Id.String()
	resp.PetId = t.PetId.String()
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
	resp.RecurringRule = t.RecurringRule

	return resp
}
