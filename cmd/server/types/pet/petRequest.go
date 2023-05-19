package pet

type PetRequest struct {
	Name          string            `json:"name"`
	DateOfBirth   int64             `json:"dateOfBirth"`
	Sex           string            `json:"sex"`
	BreedName     string            `json:"breedName"`
	Colors        []string          `json:"colors"`
	Description   string            `json:"description,omitempty"`
	Pedigree      string            `json:"pedigree,omitempty"`
	Microchip     string            `json:"microchip,omitempty"`
	WeightHistory []WeightEntry     `json:"weightHistory,omitempty"`
	VetId         string            `json:"vetIds,omitempty"`
	Metas         map[string]string `json:"metas"`
}
