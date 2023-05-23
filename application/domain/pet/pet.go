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
	DateOfBirth   time.Time
	Sex           string
	BreedName     string
	Colors        []string
	Description   string
	Pedigree      string
	Microchip     string
	WeightHistory map[int64]float64
	OwnerId       uuid.UUID
	VetId         uuid.UUID
	Metas         map[string]string
}

var Nil = Pet{}
