package base

import (
	"time"
)

type Base struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}
