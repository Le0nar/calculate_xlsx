package portfolio

import (
	"os"

	"github.com/google/uuid"
)

type Portfolio struct {
	UserID  uuid.UUID `json:"userID" db:"user_id"`
	Capital float64   `json:"capital" db:"capital"`
}

type CaclulatePortfolioDto struct {
	UserID uuid.UUID `json:"userID" binding:"required" db:"user_id"`
	File   os.File   `json:"file" binding:"required" db:"file"`
}
