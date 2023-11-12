package petConverter

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/petJournal/application/domain/pet"
	"github.com/scarlettmiss/petJournal/application/domain/user"
	pet2 "github.com/scarlettmiss/petJournal/cmd/server/types/pet"
	"github.com/scarlettmiss/petJournal/converters/userConverter"
	"github.com/scarlettmiss/petJournal/utils"
	"time"
)

func PetCreateRequestToPet(requestBody pet2.PetCreateRequest, ownerId uuid.UUID, vetId uuid.UUID) (pet.Pet, error) {
	p := pet.Pet{}
	if utils.TextIsEmpty(requestBody.Name) {
		return pet.Nil, pet.ErrNoValidName
	}
	p.Name = requestBody.Name
	if requestBody.DateOfBirth == 0 {
		return pet.Nil, pet.ErrNoValidBirthDate
	}
	p.DateOfBirth = time.Unix(requestBody.DateOfBirth/1000, (requestBody.DateOfBirth%1000)*1000000)
	if requestBody.DateOfBirth == 0 {
		return pet.Nil, pet.ErrNoValidBirthDate
	}
	gender, err := pet.ParseGender(requestBody.Gender)
	if err != nil {
		return pet.Nil, err
	}
	p.Gender = gender
	if utils.TextIsEmpty(requestBody.BreedName) {
		return pet.Nil, pet.ErrNoValidBreedname
	}
	p.BreedName = requestBody.BreedName
	p.Colors = requestBody.Colors
	p.Description = requestBody.Description
	p.Pedigree = requestBody.Pedigree
	p.Microchip = requestBody.Microchip
	p.OwnerId = ownerId
	p.VetId = vetId
	p.Metas = requestBody.Metas
	p.Avatar = requestBody.Avatar

	return p, nil
}

func PetUpdateRequestToPet(requestBody pet2.PetUpdateRequest, p pet.Pet, vetId uuid.UUID) (pet.Pet, error) {
	if utils.TextIsEmpty(requestBody.Name) {
		return p, pet.ErrNoValidName
	}
	p.Name = requestBody.Name
	if requestBody.DateOfBirth == 0 {
		return p, pet.ErrNoValidBirthDate

	}
	p.DateOfBirth = time.Unix(requestBody.DateOfBirth/1000, (requestBody.DateOfBirth%1000)*1000000)
	gender, err := pet.ParseGender(requestBody.Gender)
	if err != nil {
		return p, err
	}
	p.Gender = gender
	p.BreedName = requestBody.BreedName
	p.Colors = requestBody.Colors
	p.Description = requestBody.Description
	p.Pedigree = requestBody.Pedigree
	p.Microchip = requestBody.Microchip
	p.VetId = vetId
	p.Metas = requestBody.Metas
	p.Avatar = requestBody.Avatar

	return p, nil
}

func PetToResponse(pet pet.Pet, owner user.User, vet user.User) pet2.PetResponse {
	p := pet2.PetResponse{}
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
	p.Owner = userConverter.UserToResponse(owner)
	if vet != user.Nil {
		p.Vet = userConverter.UserToResponse(vet)
	}
	p.Metas = pet.Metas
	p.Avatar = pet.Avatar

	return p
}

func PetToSimplifiedResponse(pet pet.Pet, owner user.User, vet user.User) pet2.PetResponse {
	p := pet2.PetResponse{}
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
	p.Microchip = pet.Microchip
	p.Owner = userConverter.UserToSimplifiedResponse(owner)
	p.Avatar = pet.Avatar
	if vet != user.Nil {
		p.Vet = userConverter.UserToResponse(vet)
	}

	return p
}

func PetToVerySimplifiedResponse(pet pet.Pet) pet2.PetResponse {
	p := pet2.PetResponse{}
	p.Id = pet.Id.String()
	p.Name = pet.Name
	p.DateOfBirth = pet.DateOfBirth.UnixMilli()
	p.Gender = string(pet.Gender)

	return p
}
