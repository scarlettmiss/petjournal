package pet

import (
	"github.com/google/uuid"
	"time"
)

type Pet struct {
	Id            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Deleted       bool
	Name          string
	DateOfBirth   string
	Sex           string
	BreedName     string
	Color         string
	Description   string
	Pedigree      string
	Microchip     string
	Friendly      bool
	WeightHistory map[time.Time]float64
	OwnerId       string
}

var Nil = Pet{}
