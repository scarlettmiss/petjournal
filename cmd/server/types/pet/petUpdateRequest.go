package pet

type PetUpdateRequest struct {
	Avatar      string   `json:"avatar,omitempty"`
	Name        string   `json:"name"`
	DateOfBirth int64    `json:"dateOfBirth"`
	Gender      string   `json:"gender"`
	BreedName   string   `json:"breedName,omitempty"`
	Colors      []string `json:"colors,omitempty"`
	Description string   `json:"description,omitempty"`
	Pedigree    string   `json:"pedigree,omitempty"`
	Microchip   string   `json:"microchip,omitempty"`
	VetId       string   `json:"vetId,omitempty"`
	Metas       []Meta   `json:"metas,omitempty"`
}
