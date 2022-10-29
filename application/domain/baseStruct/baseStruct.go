package baseStruct

import (
	"time"
)

type BaseStruct struct {
	id        string
	createdAt time.Time
	updatedAt time.Time
	deleted   bool
}
