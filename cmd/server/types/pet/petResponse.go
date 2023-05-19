package pet

type PetResponse struct {
	Name          string            `json:"name"`
	DateOfBirth   int64             `json:"dateOfBirth"`
	Sex           string            `json:"sex"`
	BreedName     string            `json:"breedName"`
	Colors        []string          `json:"colors"`
	Description   string            `json:"description,omitempty"`
	Pedigree      string            `json:"pedigree,omitempty"`
	Microchip     string            `json:"microchip,omitempty"`
	WeightHistory []WeightEntry     `json:"weightHistory,omitempty"`
	OwnerId       string            `json:"ownerId"`
	VetId         string            `json:"vetId,omitempty"`
	Metas         map[string]string `json:"metas"`
}
