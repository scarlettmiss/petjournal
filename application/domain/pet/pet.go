package pet

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Gender string

const (
	M Gender = "M"
	F Gender = "F"
)

var types = map[Gender]Gender{
	M: M,
	F: F,
}

func ParseGender(value string) (Gender, error) {
	value = strings.TrimSpace(strings.ToUpper(value))
	gender, ok := types[Gender(value)]
	if !ok {
		return "", errors.New("gender not valid")
	}
	return gender, nil
}

type Pet struct {
	Id            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Deleted       bool
	Name          string
	DateOfBirth   time.Time
	Gender        Gender
	BreedName     string
	Colors        []string
	Description   string
	Pedigree      string
	Microchip     string
	WeightMin     float64
	WeightMax     float64
	WeightHistory map[time.Time]float64
	OwnerId       uuid.UUID
	VetId         uuid.UUID
	Metas         map[string]string
}

var Nil = Pet{}
