package fluidcoins

import (
	"time"

	"github.com/google/uuid"
)

// Coin is used to represent the structure of a currency
type Coin struct {
	ID                  uuid.UUID `json:"id"`
	HumanReadableName   string    `json:"human_readable_name"`
	Code                string    `json:"code"`
	Logo                string    `json:"logo"`
	IsStableCoin        bool      `json:"is_stable_coin"`
	IsAvailableOnWidget bool      `json:"is_available_on_widget"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
