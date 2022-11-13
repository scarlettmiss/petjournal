package pet

import (
	"github.com/scarlettmiss/bestPal/application/domain/baseStruct"
	"github.com/scarlettmiss/bestPal/application/domain/user"
)

type Pet struct {
	baseStruct.BaseStruct
	name        string
	dateOfBirth string
	sex         string
	breed       string
	color       *string
	description *string
	pedigree    *string
	microchip   *string
	treatments  *[]string
	weight      *[]float64
	owner       user.User
	vet         *[]user.User
}
