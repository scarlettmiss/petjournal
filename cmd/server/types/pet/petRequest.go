package pet

type PetRequest struct {
	Name          string        `json:"name"`
	DateOfBirth   int64         `json:"dateOfBirth"`
	Sex           string        `json:"sex"`
	BreedName     string        `json:"breedName"`
	Colors        []string      `json:"colors"`
	Description   string        `json:"description,omitempty"`
	Pedigree      string        `json:"pedigree,omitempty"`
	Microchip     string        `json:"microchip,omitempty"`
	WeightHistory []WeightEntry `json:"weightHistory,omitempty"`
	VetId         string        `json:"vetId,omitempty"`
	Metas         []Meta        `json:"metas"`
}

type Meta struct {
	Key   string `json:"key,omitempty"` //nano
	Value string `json:"value,omitempty"`
}
