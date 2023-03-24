package base

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	Id        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}
