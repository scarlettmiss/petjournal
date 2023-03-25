package pet

type PetResponse struct {
	Name          string        `json:"name"`
	DateOfBirth   string        `json:"dateOfBirth"`
	Sex           string        `json:"sex"`
	BreedName     string        `json:"breedName"`
	Colors        string        `json:"colors"`
	Description   string        `json:"description"`
	Pedigree      string        `json:"pedigree"`
	Microchip     string        `json:"microchip"`
	Friendly      bool          `json:"friendly"`
	WeightHistory []WeightEntry `json:"weightHistory,omitempty"`
	OwnerIds      []string      `json:"OwnerIds"`
	VetIds        []string      `json:"vetIds"`
}
