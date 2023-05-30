package pet

import "time"

type WeightEntry struct {
	Date   time.Time `json:"date"`
	Weight float64   `json:"weight"`
}

type WeightEntryRequest struct {
	Date   int64   `json:"date"`
	Weight float64 `json:"weight"`
}
