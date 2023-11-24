package record

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Type string

const (
	Vaccine      Type = "vaccine"
	Weight       Type = "weight"
	Temperature  Type = "temperature"
	Nutering     Type = "nutering"
	Surgery      Type = "surgery"
	Medicine     Type = "medicine"
	Endoparasite Type = "endoparasite"
	Ectoparasite Type = "ectoparasite"
	Examination  Type = "examination"
	Microchip    Type = "microchip"
	Diagnostic   Type = "diagnostic"
	Dental       Type = "dental"
	Other        Type = "other"
)

var recordTypes = map[Type]Type{
	Vaccine:      Vaccine,
	Weight:       Weight,
	Temperature:  Temperature,
	Nutering:     Nutering,
	Surgery:      Surgery,
	Medicine:     Medicine,
	Endoparasite: Endoparasite,
	Ectoparasite: Ectoparasite,
	Examination:  Examination,
	Microchip:    Microchip,
	Diagnostic:   Diagnostic,
	Dental:       Dental,
	Other:        Other,
}

func ParseType(value string) (Type, error) {
	value = strings.TrimSpace(strings.ToLower(value))
	typ, ok := recordTypes[Type(value)]
	if !ok {
		return Other, errors.New("type not found")
	}
	return typ, nil
}

type Record struct {
	Id             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Deleted        bool
	PetId          uuid.UUID
	RecordType     Type
	Name           string
	Date           time.Time
	Lot            string
	Result         string
	Description    string
	Notes          string
	AdministeredBy uuid.UUID
	VerifiedBy     uuid.UUID
	NextDate       time.Time
}

var Nil = Record{}
