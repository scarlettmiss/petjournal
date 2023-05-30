package petConverter

import (
	"github.com/google/uuid"
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
	p.WeightHistory = weightEntriesToMap(requestBody.WeightHistory)
	p.OwnerId = ownerId
	vetId, err := getVetId(requestBody.VetId)
	if err != nil {
		return pet.Nil, err
	}
	p.VetId = vetId
	p.Metas = metaToMap(requestBody.Metas)

	return p, nil
}

func PetUpdateRequestToPet(requestBody pet2.PetRequest, p pet.Pet) (pet.Pet, error) {
	if requestBody.Name != "" {
		p.Name = requestBody.Name
	}
	if requestBody.DateOfBirth != 0 {
		p.DateOfBirth = time.Unix(requestBody.DateOfBirth/1000, (requestBody.DateOfBirth%1000)*1000000)
	}
	if requestBody.Sex != "" {
		p.Sex = requestBody.Sex
	}
	if requestBody.BreedName != "" {
		p.BreedName = requestBody.BreedName
	}
	if len(requestBody.Colors) > 0 {
		p.Colors = requestBody.Colors
	}
	if requestBody.Description != "" {
		p.Description = requestBody.Description
	}
	if requestBody.Pedigree != "" {
		p.Pedigree = requestBody.Pedigree
	}
	if requestBody.Microchip != "" {
		p.Microchip = requestBody.Microchip
	}
	if len(requestBody.WeightHistory) > 0 {
		p.WeightHistory = weightEntriesToMap(requestBody.WeightHistory)
	}
	if requestBody.VetId != "" {
		vetId, err := getVetId(requestBody.VetId)
		if err != nil {
			return pet.Nil, err
		}
		p.VetId = vetId
	}
	if len(requestBody.Metas) > 0 {
		p.Metas = metaToMap(requestBody.Metas)
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

func PetToResponse(pet pet.Pet) pet2.PetResponse {
	p := pet2.PetResponse{}
	p.Id = pet.Id.String()
	p.Name = pet.Name
	p.DateOfBirth = pet.DateOfBirth
	p.Sex = pet.Sex
	p.BreedName = pet.BreedName
	p.Colors = pet.Colors
	p.Description = pet.Description
	p.Pedigree = pet.Pedigree
	p.Microchip = pet.Microchip
	p.WeightHistory = weightMapToEntries(pet.WeightHistory)
	p.OwnerId = pet.OwnerId.String()
	if pet.VetId != uuid.Nil {
		p.VetId = pet.VetId.String()
	}
	p.Metas = pet.Metas

	return p
}
