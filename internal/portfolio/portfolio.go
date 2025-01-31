package portfolio

import (
	"github.com/google/uuid"
)

type Portfolio struct {
	UserID  uuid.UUID `json:"userID" db:"user_id"`
	Capital float64   `json:"capital" db:"capital"`
}
