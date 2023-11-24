package api

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application"
	"github.com/scarlettmiss/petJournal/application/domain/pet"
	"github.com/scarlettmiss/petJournal/application/domain/record"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	"github.com/scarlettmiss/petJournal/utils"
	"time"
)

func PetCreateRequestToPetCreateOpts(requestBody PetCreateRequest, ownerId uuid.UUID) (application.PetCreateOptions, error) {
	vId := uuid.Nil
	if !utils.TextIsEmpty(requestBody.VetId) {
		var err error
		vId, err = uuid.Parse(requestBody.VetId)
		if err != nil {
			return application.PetCreateOptions{}, err
		}
	}

	p := application.PetCreateOptions{}
	p.Name = requestBody.Name
	p.DateOfBirth = time.Unix(requestBody.DateOfBirth/1000, (requestBody.DateOfBirth%1000)*1000000)
	p.Gender = requestBody.Gender
	p.BreedName = requestBody.BreedName
	p.Colors = requestBody.Colors
	p.Description = requestBody.Description
	p.Pedigree = requestBody.Pedigree
	p.Microchip = requestBody.Microchip
	p.OwnerId = ownerId
	p.VetId = vId
	p.Metas = requestBody.Metas
	p.Avatar = requestBody.Avatar

	return p, nil
}

func PetUpdateRequestToPetUpdateOpts(requestBody PetUpdateRequest, pId uuid.UUID, uId uuid.UUID) (application.PetUpdateOptions, error) {
	vId := uuid.Nil
	if !utils.TextIsEmpty(requestBody.VetId) {
		var err error
		vId, err = uuid.Parse(requestBody.VetId)
		if err != nil {
			return application.PetUpdateOptions{}, err
		}
	}

	opts := application.PetUpdateOptions{}
	opts.Id = pId
	opts.Name = requestBody.Name
	opts.DateOfBirth = time.Unix(requestBody.DateOfBirth/1000, (requestBody.DateOfBirth%1000)*1000000)
	opts.Gender = requestBody.Gender
	opts.BreedName = requestBody.BreedName
	opts.Colors = requestBody.Colors
	opts.Description = requestBody.Description
	opts.Pedigree = requestBody.Pedigree
	opts.Microchip = requestBody.Microchip
	opts.VetId = vId
	opts.OwnerId = uId
	opts.Metas = requestBody.Metas
	opts.Avatar = requestBody.Avatar

	return opts, nil
}

func PetToResponse(pet pet.Pet, owner user.User, vet user.User) PetResponse {
	p := PetResponse{}
	p.Id = pet.Id.String()
	p.CreatedAt = pet.CreatedAt.UnixMilli()
	p.UpdatedAt = pet.UpdatedAt.UnixMilli()
	p.Deleted = pet.Deleted
	p.Name = pet.Name
	p.DateOfBirth = pet.DateOfBirth.UnixMilli()
	p.Gender = string(pet.Gender)
	p.BreedName = pet.BreedName
	p.Colors = pet.Colors
	p.Description = pet.Description
	p.Pedigree = pet.Pedigree
	p.Microchip = pet.Microchip
	p.Owner = UserToResponse(owner)
	if vet != user.Nil {
		p.Vet = UserToResponse(vet)
	}
	p.Metas = pet.Metas
	p.Avatar = pet.Avatar

	return p
}

func PetToVerySimplifiedResponse(pet pet.Pet) PetResponse {
	p := PetResponse{}
	p.Id = pet.Id.String()
	p.Name = pet.Name
	p.DateOfBirth = pet.DateOfBirth.UnixMilli()
	p.Gender = string(pet.Gender)

	return p
}

func RecordCreateRequestToRecord(requestBody RecordCreateRequest, petId uuid.UUID, administeredBy user.User) application.RecordCreateOptions {
	opts := application.RecordCreateOptions{}
	verifierId := uuid.Nil
	if administeredBy.UserType == user.Vet {
		verifierId = administeredBy.Id
	}
	opts.PetId = petId
	opts.RecordType = requestBody.RecordType
	opts.Name = requestBody.Name
	opts.Date = time.Unix(requestBody.Date/1000, (requestBody.Date%1000)*1000000)
	opts.Lot = requestBody.Lot
	opts.Result = requestBody.Result
	opts.Description = requestBody.Description
	opts.Notes = requestBody.Notes
	opts.AdministeredBy = administeredBy.Id
	opts.VerifiedBy = verifierId
	opts.NextDate = time.Unix(requestBody.NextDate/1000, (requestBody.NextDate%1000)*1000000)

	return opts
}

func RecordUpdateRequestToRecord(requestBody RecordUpdateRequest, rId uuid.UUID, updatedBy user.User) application.RecordUpdateOptions {
	opts := application.RecordUpdateOptions{}
	verifierId := uuid.Nil
	if updatedBy.UserType == user.Vet {
		verifierId = updatedBy.Id
	}
	opts.Id = rId
	opts.VerifiedBy = verifierId
	opts.RecordType = requestBody.RecordType
	opts.Name = requestBody.Name
	opts.Date = time.Unix(requestBody.Date/1000, (requestBody.Date%1000)*1000000)
	opts.Lot = requestBody.Lot
	opts.Result = requestBody.Result
	opts.Description = requestBody.Description
	opts.Notes = requestBody.Notes
	opts.NextDate = time.Unix(requestBody.NextDate/1000, (requestBody.NextDate%1000)*1000000)
	return opts
}

func RecordToResponse(r record.Record, pet pet.Pet, administeredBy user.User, verifiedBy user.User) RecordResponse {
	resp := RecordResponse{}
	resp.Id = r.Id.String()
	resp.CreatedAt = r.CreatedAt.UnixMilli()
	resp.UpdatedAt = r.UpdatedAt.UnixMilli()
	resp.Deleted = r.Deleted
	resp.Pet = PetToVerySimplifiedResponse(pet)
	resp.RecordType = string(r.RecordType)
	resp.Name = r.Name
	resp.Date = r.Date.UnixMilli()
	resp.Lot = r.Lot
	resp.Result = r.Result
	resp.Description = r.Description
	resp.Notes = r.Notes
	resp.AdministeredBy = UserToResponse(administeredBy)
	if verifiedBy != user.Nil {
		resp.VerifiedBy = UserToResponse(verifiedBy)
	}
	if !r.NextDate.IsZero() {
		resp.NextDate = r.NextDate.UnixMilli()
	}

	return resp
}

func UserCreateRequestToUserCreateOptions(requestBody UserCreateRequest) application.UserCreateOptions {
	uOpts := application.UserCreateOptions{}
	uOpts.UserType = requestBody.UserType
	uOpts.Email = requestBody.Email
	uOpts.Password = requestBody.Password
	uOpts.Name = requestBody.Name
	uOpts.Surname = requestBody.Surname
	uOpts.Phone = requestBody.Phone
	uOpts.Address = requestBody.Address
	uOpts.City = requestBody.City
	uOpts.State = requestBody.State
	uOpts.Country = requestBody.Country
	uOpts.Zip = requestBody.Zip
	return uOpts
}

func UserUpdateRequestToUserOptions(requestBody UserUpdateRequest, uId uuid.UUID) application.UserUpdateOptions {
	uOpts := application.UserUpdateOptions{}
	uOpts.Id = uId
	uOpts.Email = requestBody.Email
	uOpts.Name = requestBody.Name
	uOpts.Surname = requestBody.Surname
	uOpts.Phone = requestBody.Phone
	uOpts.Address = requestBody.Address
	uOpts.City = requestBody.City
	uOpts.State = requestBody.State
	uOpts.Country = requestBody.Country
	uOpts.Zip = requestBody.Zip
	return uOpts
}

func UserToResponse(u user.User) UserResponse {
	resp := UserResponse{}
	resp.Id = u.Id.String()
	resp.CreatedAt = u.CreatedAt.UnixMilli()
	resp.UpdatedAt = u.UpdatedAt.UnixMilli()
	resp.Deleted = u.Deleted
	resp.UserType = u.UserType
	resp.Email = u.Email
	resp.Name = u.Name
	resp.Surname = u.Surname
	resp.Phone = u.Phone
	resp.Address = u.Address
	resp.City = u.City
	resp.State = u.State
	resp.Country = u.Country
	resp.Zip = u.Zip
	return resp
}
