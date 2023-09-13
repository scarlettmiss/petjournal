package pet

type PetUpdateRequest struct {
	Avatar        string               `json:"avatar,omitempty"`
	Name          string               `json:"name,omitempty"`
	DateOfBirth   int64                `json:"dateOfBirth,omitempty"`
	Gender        string               `json:"gender,omitempty"`
	BreedName     string               `json:"breedName,omitempty"`
	Colors        []string             `json:"colors,omitempty"`
	Description   string               `json:"description,omitempty"`
	Pedigree      string               `json:"pedigree,omitempty"`
	Microchip     string               `json:"microchip,omitempty"`
	WeightMin     float64              `json:"weightMin,omitempty"`
	WeightMax     float64              `json:"weightMax,omitempty"`
	WeightHistory []WeightEntryRequest `json:"weightHistory,omitempty"`
	Metas         []Meta               `json:"metas,omitempty"`
}
