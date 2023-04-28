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
	Colors        []string
	Description   string
	Pedigree      string
	Microchip     string
	WeightHistory []WeightEntry
	OwnerId       uuid.UUID
	VetId         uuid.UUID
	Metas         map[string]string
}

type WeightEntry struct {
	Date   uint64
	Weight float64
}

var Nil = Pet{}
