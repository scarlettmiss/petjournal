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
	p.Name = requestBody.Name
	p.DateOfBirth = time.Unix(requestBody.DateOfBirth/1000, (requestBody.DateOfBirth%1000)*1000000)
	p.Sex = requestBody.Sex
	p.BreedName = requestBody.BreedName
	p.Colors = requestBody.Colors
	p.Description = requestBody.Description
	p.Pedigree = requestBody.Pedigree
	p.Microchip = requestBody.Microchip
	p.WeightHistory = weightEntriesToMap(requestBody.WeightHistory)
	vetId, err := getVetId(requestBody.VetId)
	if err != nil {
		return pet.Nil, err
	}
	p.VetId = vetId
	p.Metas = metaToMap(requestBody.Metas)

	return p, nil
}

func metaToMap(metas []pet2.Meta) map[string]string {
	metaMap := make(map[string]string)

	for _, meta := range metas {
		metaMap[meta.Key] = meta.Value
	}
	return metaMap
}

func weightEntriesToMap(weightEntries []pet2.WeightEntry) map[int64]float64 {
	weightMap := make(map[int64]float64)

	for _, entry := range weightEntries {
		weightMap[entry.Date] = entry.Weight
	}
	return weightMap
}

func weightMapToEntries(weightEntries map[int64]float64) []pet2.WeightEntry {
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
	p.DateOfBirth = pet.DateOfBirth.Unix()
	p.Sex = pet.Sex
	p.BreedName = pet.BreedName
	p.Colors = pet.Colors
	p.Description = pet.Description
	p.Pedigree = pet.Pedigree
	p.Microchip = pet.Microchip
	p.WeightHistory = weightMapToEntries(pet.WeightHistory)
	p.OwnerId = pet.OwnerId.String()
	if pet.VetId == uuid.Nil {
		p.VetId = ""
	} else {
		p.VetId = pet.VetId.String()
	}
	p.Metas = pet.Metas

	return p
}
