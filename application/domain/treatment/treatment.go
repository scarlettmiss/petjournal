package treatment

import (
	"github.com/scarlettmiss/bestPal/application/domain/baseStruct"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	"time"
)

type Type string

const (
	Vaccine      Type = "VACCINE"
	Surgery      Type = "SURGERY"
	Medicine     Type = "MEDICINE"
	Endoparasite Type = "ENDOPARASITE"
	Ectoparasite Type = "ECTOPARASITE"
	examination  Type = "EXAMINATION"
	microchip    Type = "MICROCHIP"
	Diagnostic   Type = "DIAGNOSTIC"
	Other        Type = "OTHER"
)

type Treatment struct {
	baseStruct.BaseStruct
	treatmentType Type
	name          string
	date          time.Time
	lot           *string
	result        *string
	description   string
	administer    user.User
}
