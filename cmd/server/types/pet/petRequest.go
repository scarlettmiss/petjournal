package pet

import "github.com/google/uuid"

type PetRequest struct {
	Name          string            `json:"name"`
	DateOfBirth   string            `json:"dateOfBirth"`
	Sex           string            `json:"sex"`
	BreedName     string            `json:"breedName"`
	Colors        []string          `json:"colors"`
	Description   string            `json:"description,omitempty"`
	Pedigree      string            `json:"pedigree,omitempty"`
	Microchip     string            `json:"microchip,omitempty"`
	WeightHistory []WeightEntry     `json:"weightHistory,omitempty"`
	VetId         uuid.UUID         `json:"vetIds,omitempty"`
	Metas         map[string]string `json:"metas"`
}
