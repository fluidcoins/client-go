package fluidcoins

import (
	"time"

	"github.com/google/uuid"
)

type Merchant struct {
	ID           uuid.UUID `json:"id"`
	BusinessName string    `json:"business_name"`
	Description  string    `json:"description"`
	IndustryID   uuid.UUID `json:"industry_id"`

	CountryID uuid.UUID `json:"country_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
