package fluidcoins

import (
	"time"

	"github.com/google/uuid"
)

type model struct {
	ID uuid.UUID `json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
