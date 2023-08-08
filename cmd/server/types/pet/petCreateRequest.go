package pet

type PetCreateRequest struct {
	Name          string               `json:"name"`
	DateOfBirth   int64                `json:"dateOfBirth"`
	Gender        string               `json:"gender"`
	BreedName     string               `json:"breedName"`
	Colors        []string             `json:"colors,omitempty"`
	Description   string               `json:"description,omitempty"`
	Pedigree      string               `json:"pedigree,omitempty"`
	Microchip     string               `json:"microchip,omitempty"`
	WeightMin     float64              `json:"weightMin,omitempty"`
	WeightMax     float64              `json:"weightMax,omitempty"`
	WeightHistory []WeightEntryRequest `json:"weightHistory,omitempty"`
	VetId         string               `json:"vetId,omitempty"`
	Metas         []Meta               `json:"metas,omitempty"`
	Avatar        string               `json:"avatar,omitempty"`
}
