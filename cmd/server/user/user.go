package user

import (
	"time"

	"github.com/google/uuid"
)

// User represents the user entity.
type User struct {
	ID         uuid.UUID
	Username   string
	Email      string
	CreateTime time.Time
	UpdateTime time.Time
}
