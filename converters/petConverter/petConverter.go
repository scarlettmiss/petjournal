package petConverter

import (
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/scarlettmiss/bestPal/application/domain/pet"
	pet2 "github.com/scarlettmiss/bestPal/cmd/server/types/pet"
	"time"
)

func getVetId(id string) (uuid.UUID, error) {
	if id == "" {
		return uuid.Nil, nil
	}

	vetId, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, err
	}

	return vetId, nil
}

func PetCreateRequestToPet(requestBody pet2.PetRequest, ownerId uuid.UUID) (pet.Pet, error) {
	p := pet.Pet{}
	p.Name = requestBody.Name
	p.DateOfBirth = time.Unix(requestBody.DateOfBirth/1000, (requestBody.DateOfBirth%1000)*1000000)
	p.Sex = requestBody.Sex
	p.BreedName = requestBody.BreedName
	p.Colors = requestBody.Colors
	p.Description = requestBody.Description
	p.Pedigree = requestBody.Pedigree
	p.Microchip = requestBody.Microchip
	p.WeightHistory = weightEntryToPet(requestBody.WeightHistory)
	p.OwnerId = ownerId
	vetId, err := getVetId(requestBody.VetId)
	if err != nil {
		return pet.Nil, err
	}
	p.VetId = vetId
	p.Metas = requestBody.Metas

	return p, nil
}

func PetUpdateRequestToPet(requestBody pet2.PetRequest, p pet.Pet) (pet.Pet, error) {
	p.Name = requestBody.Name
	p.DateOfBirth = time.Unix(requestBody.DateOfBirth/1000, (requestBody.DateOfBirth%1000)*1000000)
	p.Sex = requestBody.Sex
	p.BreedName = requestBody.BreedName
	p.Colors = requestBody.Colors
	p.Description = requestBody.Description
	p.Pedigree = requestBody.Pedigree
	p.Microchip = requestBody.Microchip
	p.WeightHistory = weightEntryToPet(requestBody.WeightHistory)
	vetId, err := getVetId(requestBody.VetId)
	if err != nil {
		return pet.Nil, err
	}
	p.VetId = vetId
	p.Metas = requestBody.Metas

	return p, nil
}

func weightEntryToPet(weightEntries []pet2.WeightEntry) []pet.WeightEntry {
	return lo.Map(weightEntries, func(entry pet2.WeightEntry, _ int) pet.WeightEntry {
		weightEntry := pet.WeightEntry{}
		weightEntry.Weight = entry.Weight
		weightEntry.Date = entry.Date
		return weightEntry
	})
}

func weightEntryToResponse(weightEntries []pet.WeightEntry) []pet2.WeightEntry {
	return lo.Map(weightEntries, func(entry pet.WeightEntry, _ int) pet2.WeightEntry {
		weightEntry := pet2.WeightEntry{}
		weightEntry.Weight = entry.Weight
		weightEntry.Date = entry.Date
		return weightEntry
	})
}

func PetToResponse(pet pet.Pet) pet2.PetResponse {
	p := pet2.PetResponse{}
	p.Name = pet.Name
	p.DateOfBirth = pet.DateOfBirth.Unix()
	p.Sex = pet.Sex
	p.BreedName = pet.BreedName
	p.Colors = pet.Colors
	p.Description = pet.Description
	p.Pedigree = pet.Pedigree
	p.Microchip = pet.Microchip
	p.WeightHistory = weightEntryToResponse(pet.WeightHistory)
	p.OwnerId = pet.OwnerId.String()
	if pet.VetId == uuid.Nil {
		p.VetId = ""
	}
	p.VetId = pet.VetId.String()
	p.Metas = pet.Metas

	return p
}
