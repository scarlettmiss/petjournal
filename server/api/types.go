package api

import (
	"github.com/scarlettmiss/petJournal/application/domain/user"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateRequest struct {
	UserType string `json:"userType"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone,omitempty"`
	Address  string `json:"address,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	Country  string `json:"country,omitempty"`
	Zip      string `json:"zip,omitempty"`
}

type UserUpdateRequest struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
	Zip     string `json:"zip,omitempty"`
}

type UserResponse struct {
	Id        string    `json:"id,omitempty"`
	CreatedAt int64     `json:"createdAt,omitempty"`
	UpdatedAt int64     `json:"updatedAt,omitempty"`
	Deleted   bool      `json:"deleted,omitempty"`
	UserType  user.Type `json:"userType,omitempty"`
	Email     string    `json:"email,omitempty"`
	Name      string    `json:"name,omitempty"`
	Surname   string    `json:"surname,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Address   string    `json:"address,omitempty"`
	City      string    `json:"city,omitempty"`
	State     string    `json:"state,omitempty"`
	Country   string    `json:"country,omitempty"`
	Zip       string    `json:"zip,omitempty"`
}

type RecordCreateRequest struct {
	RecordType  string `json:"recordType"`
	Name        string `json:"name,omitempty"`
	Date        int64  `json:"date"`
	Lot         string `json:"lot,omitempty"`
	Result      string `json:"result,omitempty"`
	Description string `json:"description,omitempty"`
	Notes       string `json:"notes,omitempty"`
	NextDate    int64  `json:"nextDate,omitempty"`
}

type RecordUpdateRequest struct {
	RecordType  string `json:"recordType"`
	Name        string `json:"name,omitempty"`
	Date        int64  `json:"date"`
	Lot         string `json:"lot,omitempty"`
	Result      string `json:"result,omitempty"`
	Description string `json:"description,omitempty"`
	Notes       string `json:"notes,omitempty"`
	NextDate    int64  `json:"nextDate,omitempty"`
}

type RecordResponse struct {
	Id             string        `json:"id"`
	CreatedAt      int64         `json:"createdAt"`
	UpdatedAt      int64         `json:"updatedAt"`
	Deleted        bool          `json:"deleted"`
	Pet            PetResponse   `json:"pet"`
	RecordType     string        `json:"recordType"`
	Name           string        `json:"name"`
	Date           int64         `json:"date"`
	Lot            string        `json:"lot,omitempty"`
	Result         string        `json:"result,omitempty"`
	Description    string        `json:"description,omitempty"`
	Notes          string        `json:"notes,omitempty"`
	AdministeredBy *UserResponse `json:"administeredBy,omitempty"`
	VerifiedBy     *UserResponse `json:"verifiedBy,omitempty"`
	GroupId        string        `json:"groupId,omitempty"`
}

type PetCreateRequest struct {
	Name        string            `json:"name"`
	DateOfBirth int64             `json:"dateOfBirth"`
	Gender      string            `json:"gender"`
	BreedName   string            `json:"breedName"`
	Colors      []string          `json:"colors,omitempty"`
	Description string            `json:"description,omitempty"`
	Pedigree    string            `json:"pedigree,omitempty"`
	Microchip   string            `json:"microchip,omitempty"`
	VetId       string            `json:"vetId,omitempty"`
	Metas       map[string]string `json:"metas,omitempty"`
	Avatar      string            `json:"avatar,omitempty"`
}

type PetUpdateRequest struct {
	Avatar      string            `json:"avatar,omitempty"`
	Name        string            `json:"name"`
	DateOfBirth int64             `json:"dateOfBirth"`
	Gender      string            `json:"gender"`
	BreedName   string            `json:"breedName"`
	Colors      []string          `json:"colors,omitempty"`
	Description string            `json:"description,omitempty"`
	Pedigree    string            `json:"pedigree,omitempty"`
	Microchip   string            `json:"microchip,omitempty"`
	VetId       string            `json:"vetId,omitempty"`
	Metas       map[string]string `json:"metas,omitempty"`
}

type PetResponse struct {
	Id          string            `json:"id,omitempty"`
	CreatedAt   int64             `json:"createdAt,omitempty"`
	UpdatedAt   int64             `json:"updatedAt,omitempty"`
	Deleted     bool              `json:"deleted,omitempty"`
	Name        string            `json:"name,omitempty"`
	DateOfBirth int64             `json:"dateOfBirth,omitempty"`
	Gender      string            `json:"gender,omitempty"`
	BreedName   string            `json:"breedName,omitempty"`
	Colors      []string          `json:"colors,omitempty"`
	Description string            `json:"description,omitempty"`
	Pedigree    string            `json:"pedigree,omitempty"`
	Microchip   string            `json:"microchip,omitempty"`
	Owner       *UserResponse     `json:"owner,omitempty"`
	Vet         *UserResponse     `json:"vet,omitempty"`
	Metas       map[string]string `json:"metas,omitempty"`
	Avatar      string            `json:"avatar,omitempty"`
}
