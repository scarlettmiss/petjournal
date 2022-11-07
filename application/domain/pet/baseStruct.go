package pet

import "github.com/scarlettmiss/bestPal/application/domain/baseStruct"

type Pet struct {
	baseStruct.BaseStruct
	name       string
	treatments []string
}
