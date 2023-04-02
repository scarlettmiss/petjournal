package pet

import "github.com/google/uuid"

type PetResponse struct {
	Name          string            `json:"name"`
	DateOfBirth   string            `json:"dateOfBirth"`
	Sex           string            `json:"sex"`
	BreedName     string            `json:"breedName"`
	Colors        []string          `json:"colors"`
	Description   string            `json:"description,omitempty"`
	Pedigree      string            `json:"pedigree,omitempty"`
	Microchip     string            `json:"microchip,omitempty"`
	WeightHistory []WeightEntry     `json:"weightHistory,omitempty"`
	OwnerId       uuid.UUID         `json:"ownerId"`
	VetIds        uuid.UUID         `json:"vetId,omitempty"`
	Metas         map[string]string `json:"metas"`
}
