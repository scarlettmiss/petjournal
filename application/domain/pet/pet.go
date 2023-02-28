package pet

import (
	"github.com/scarlettmiss/bestPal/application/domain/base"
	"time"
)

type BehaviorType string

const (
	Aggressive BehaviorType = "aggressive"
	Friendly   BehaviorType = "friendly"
)

type Pet struct {
	base.Base
	Name          string
	DateOfBirth   string
	Sex           string
	BreedName     string
	Color         string
	Description   string
	Pedigree      string
	Microchip     string
	Behavior      BehaviorType
	WeightHistory map[time.Time]float64
	OwnerId       string
}

var Nil = Pet{}
