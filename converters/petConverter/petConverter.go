package petConverter

import (
	"github.com/google/uuid"
	"github.com/scarlettmiss/bestPal/application/domain/pet"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	pet2 "github.com/scarlettmiss/bestPal/cmd/server/types/pet"
	"github.com/scarlettmiss/bestPal/converters/userConverter"
	"github.com/scarlettmiss/bestPal/utils"
	"time"
)

func PetCreateRequestToPet(requestBody pet2.PetCreateRequest, ownerId uuid.UUID) (pet.Pet, error) {
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
	p.WeightMin = requestBody.WeightMin
	p.WeightMax = requestBody.WeightMax
	p.WeightHistory = weightEntriesToMap(requestBody.WeightHistory)
	p.OwnerId = ownerId
	p.Metas = metaToMap(requestBody.Metas)
	p.Avatar = requestBody.Avatar

	return p, nil
}

func PetUpdateRequestToPet(requestBody pet2.PetUpdateRequest, p pet.Pet) (pet.Pet, error) {
	if !utils.TextIsEmpty(requestBody.Name) {
		p.Name = requestBody.Name
	}
	if requestBody.DateOfBirth != 0 {
		p.DateOfBirth = time.Unix(requestBody.DateOfBirth/1000, (requestBody.DateOfBirth%1000)*1000000)
	}
	if !utils.TextIsEmpty(requestBody.Gender) {
		gender, err := pet.ParseGender(requestBody.Gender)
		if err != nil {
			return pet.Nil, err
		}
		p.Gender = gender
	}
	if !utils.TextIsEmpty(requestBody.BreedName) {
		p.BreedName = requestBody.BreedName
	}
	if len(requestBody.Colors) > 0 {
		p.Colors = requestBody.Colors
	}
	if !utils.TextIsEmpty(requestBody.Description) {
		p.Description = requestBody.Description
	}
	if !utils.TextIsEmpty(requestBody.Pedigree) {
		p.Pedigree = requestBody.Pedigree
	}
	if !utils.TextIsEmpty(requestBody.Microchip) {
		p.Microchip = requestBody.Microchip
	}
	if requestBody.WeightMin != 0 {
		p.WeightMin = requestBody.WeightMin
	}
	if requestBody.WeightMax != 0 {
		p.WeightMax = requestBody.WeightMax
	}
	if len(requestBody.WeightHistory) > 0 {
		p.WeightHistory = weightEntriesToMap(requestBody.WeightHistory)
	}
	if len(requestBody.Metas) > 0 {
		p.Metas = metaToMap(requestBody.Metas)
	}
	if !utils.TextIsEmpty(requestBody.Avatar) {
		p.Avatar = requestBody.Avatar
	}

	return p, nil
}

func metaToMap(metas []pet2.Meta) map[string]string {
	metaMap := make(map[string]string)

	for _, meta := range metas {
		metaMap[meta.Key] = meta.Value
	}
	return metaMap
}

func weightEntriesToMap(weightEntries []pet2.WeightEntryRequest) map[time.Time]float64 {
	weightMap := make(map[time.Time]float64)

	for _, entry := range weightEntries {
		date := time.Unix(entry.Date/1000, (entry.Date%1000)*1000000)
		weightMap[date] = entry.Weight
	}
	return weightMap
}

func weightMapToEntries(weightEntries map[time.Time]float64) []pet2.WeightEntry {
	weights := make([]pet2.WeightEntry, 0, len(weightEntries))

	for key, value := range weightEntries {
		meta := pet2.WeightEntry{
			Date:   key,
			Weight: value,
		}
		weights = append(weights, meta)
	}

	return weights
}

func PetToResponse(pet pet.Pet, owner user.User, vet user.User) pet2.PetResponse {
	p := pet2.PetResponse{}
	p.Id = pet.Id.String()
	p.CreatedAt = pet.CreatedAt
	p.UpdatedAt = pet.UpdatedAt
	p.Deleted = pet.Deleted
	p.Name = pet.Name
	p.DateOfBirth = pet.DateOfBirth
	p.Gender = string(pet.Gender)
	p.BreedName = pet.BreedName
	p.Colors = pet.Colors
	p.Description = pet.Description
	p.Pedigree = pet.Pedigree
	p.Microchip = pet.Microchip
	p.WeightMin = pet.WeightMin
	p.WeightMax = pet.WeightMax
	p.WeightHistory = weightMapToEntries(pet.WeightHistory)
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
	p.CreatedAt = pet.CreatedAt
	p.UpdatedAt = pet.UpdatedAt
	p.Deleted = pet.Deleted
	p.Name = pet.Name
	p.DateOfBirth = pet.DateOfBirth
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
