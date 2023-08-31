package treatment

import (
	"github.com/scarlettmiss/bestPal/cmd/server/types/user"
	"time"
)

type TreatmentResponse struct {
	Id             string            `json:"id"`
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt"`
	Deleted        bool              `json:"deleted"`
	PetId          string            `json:"petId"`
	TreatmentType  string            `json:"treatmentType"`
	Name           string            `json:"name"`
	Date           time.Time         `json:"date"`
	Lot            string            `json:"lot,omitempty"`
	Result         string            `json:"result,omitempty"`
	Description    string            `json:"description,omitempty"`
	Notes          string            `json:"notes,omitempty"`
	AdministeredBy user.UserResponse `json:"administeredBy"`
	VerifiedBy     user.UserResponse `json:"verifiedBy,omitempty"`
	RecurringRule  string            `json:"recurringRule,omitempty"`
}
