package pet

type PetRequest struct {
	Name          string               `json:"name,omitempty"`
	DateOfBirth   int64                `json:"dateOfBirth,omitempty"`
	Sex           string               `json:"sex,omitempty"`
	BreedName     string               `json:"breedName,omitempty"`
	Colors        []string             `json:"colors,omitempty"`
	Description   string               `json:"description,omitempty"`
	Pedigree      string               `json:"pedigree,omitempty"`
	Microchip     string               `json:"microchip,omitempty"`
	WeightHistory []WeightEntryRequest `json:"weightHistory,omitempty"`
	VetId         string               `json:"vetId,omitempty"`
	Metas         []Meta               `json:"metas"`
}

type Meta struct {
	Key   string `json:"key,omitempty"` //nano
	Value string `json:"value,omitempty"`
}
