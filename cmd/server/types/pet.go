package types

type Pet struct {
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
	OwnerId       string        `json:"ownerId"`
}

type WeightEntry struct {
	Date   uint64  `json:"date,omitempty"` //nano
	Weight float64 `json:"weight,omitempty"`
}
