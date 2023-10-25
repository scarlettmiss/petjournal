package pet

type PetCreateRequest struct {
	Name        string   `json:"name"`
	DateOfBirth int64    `json:"dateOfBirth"`
	Gender      string   `json:"gender"`
	BreedName   string   `json:"breedName"`
	Colors      []string `json:"colors,omitempty"`
	Description string   `json:"description,omitempty"`
	Pedigree    string   `json:"pedigree,omitempty"`
	Microchip   string   `json:"microchip,omitempty"`
	VetId       string   `json:"vetId,omitempty"`
	Metas       []Meta   `json:"metas,omitempty"`
	Avatar      string   `json:"avatar,omitempty"`
}
