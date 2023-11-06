package Record

import (
	"github.com/scarlettmiss/bestPal/cmd/server/types/pet"
	"github.com/scarlettmiss/bestPal/cmd/server/types/user"
)

type RecordResponse struct {
	Id             string            `json:"id"`
	CreatedAt      int64             `json:"createdAt"`
	UpdatedAt      int64             `json:"updatedAt"`
	Deleted        bool              `json:"deleted"`
	Pet            pet.PetResponse   `json:"pet"`
	RecordType     string            `json:"recordType"`
	Name           string            `json:"name"`
	Date           int64             `json:"date"`
	Lot            string            `json:"lot,omitempty"`
	Result         string            `json:"result,omitempty"`
	Description    string            `json:"description,omitempty"`
	Notes          string            `json:"notes,omitempty"`
	AdministeredBy user.UserResponse `json:"administeredBy"`
	VerifiedBy     user.UserResponse `json:"verifiedBy,omitempty"`
	NextDate       int64             `json:"nextDate,omitempty"`
}
