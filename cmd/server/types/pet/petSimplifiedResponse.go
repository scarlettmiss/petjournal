package pet

import (
	"github.com/scarlettmiss/bestPal/cmd/server/types/user"
	"time"
)

type PetSimplifiedResponse struct {
	Id            string                      `json:"id"`
	Name          string                      `json:"name"`
	DateOfBirth   time.Time                   `json:"dateOfBirth"`
	Sex           string                      `json:"sex"`
	BreedName     string                      `json:"breedName"`
	Colors        []string                    `json:"colors"`
	Description   string                      `json:"description,omitempty"`
	Pedigree      string                      `json:"pedigree,omitempty"`
	Microchip     string                      `json:"microchip,omitempty"`
	WeightHistory []WeightEntry               `json:"weightHistory,omitempty"`
	Owner         user.UserResponseSimplified `json:"owner"`
	Vet           user.UserResponse           `json:"vet,omitempty"`
}
