package petConverter

import (
	"github.com/samber/lo"
	"github.com/scarlettmiss/bestPal/application/domain/pet"
	pet2 "github.com/scarlettmiss/bestPal/cmd/server/types/pet"
)

func PetRequestToPet(requestBody pet2.PetRequest) (pet.Pet, error) {
	p := pet.Pet{}
	p.Name = requestBody.Name
	p.DateOfBirth = requestBody.DateOfBirth
	p.Sex = requestBody.Sex
	p.BreedName = requestBody.BreedName
	p.Colors = requestBody.Colors
	p.Description = requestBody.Description
	p.Pedigree = requestBody.Pedigree
	p.Microchip = requestBody.Microchip
	p.Friendly = requestBody.Friendly
	p.WeightHistory = weightEntryToPet(requestBody.WeightHistory)
	p.OwnerIds = requestBody.OwnerIds
	p.VetIds = requestBody.VetIds

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
	p.DateOfBirth = pet.DateOfBirth
	p.Sex = pet.Sex
	p.BreedName = pet.BreedName
	p.Colors = pet.Colors
	p.Description = pet.Description
	p.Pedigree = pet.Pedigree
	p.Microchip = pet.Microchip
	p.Friendly = pet.Friendly
	p.WeightHistory = weightEntryToResponse(pet.WeightHistory)
	p.OwnerIds = pet.OwnerIds
	p.VetIds = pet.VetIds

	return p
}
