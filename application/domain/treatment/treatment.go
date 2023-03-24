package treatment

import (
	"github.com/scarlettmiss/bestPal/application/domain/base"
	"github.com/scarlettmiss/bestPal/application/domain/user"
	"time"
)

type TreatmentType string

const (
	Vaccine      TreatmentType = "vaccine"
	Surgery      TreatmentType = "surgery"
	Medicine     TreatmentType = "medicine"
	Endoparasite TreatmentType = "endoparasite"
	Ectoparasite TreatmentType = "ectoparasite"
	Examination  TreatmentType = "examination"
	Microchip    TreatmentType = "microchip"
	Diagnostic   TreatmentType = "diagnostic"
	Dental       TreatmentType = "dental"
	Other        TreatmentType = "other"
)

type Treatment struct {
	base.Base
	PetId         string
	TreatmentType TreatmentType
	Name          string
	Date          time.Time
	Lot           string
	Result        string
	Description   string
	Notes         string
	Administer    user.User
}

var Nil = Treatment{}
