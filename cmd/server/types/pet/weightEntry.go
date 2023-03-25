package pet

type WeightEntry struct {
	Date   uint64  `json:"date,omitempty"` //nano
	Weight float64 `json:"weight,omitempty"`
}
