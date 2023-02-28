package types

import (
	"github.com/scarlettmiss/bestPal/application/domain/pet"
)

type PetDto struct {
	Name          string           `json:"name"`
	DateOfBirth   string           `json:"dateOfBirth"`
	Sex           string           `json:"sex"`
	BreedName     string           `json:"breedName"`
	Color         string           `json:"color"`
	Description   string           `json:"description"`
	Pedigree      string           `json:"pedigree"`
	Microchip     string           `json:"microchip"`
	Behavior      pet.BehaviorType `json:"behavior"`
	WeightHistory []WeightEntry    `json:"weightHistory"`
	OwnerId       string           `json:"ownerId"`
}

type WeightEntry struct {
	Date   uint64  `json:"date,omitempty"` //nano
	Weight float64 `json:"weight,omitempty"`
}
