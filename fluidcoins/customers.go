package fluidcoins

import (
	"time"

	"github.com/google/uuid"
)

// Customer is used to represent the structure of a created customer
type Customer struct {
	ID            uuid.UUID `json:"id"`
	MerchantID    uuid.UUID `json:"merchant_id"`
	FullName      string    `json:"full_name"`
	Reference     string    `json:"reference"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phone_number"`
	IsBlacklisted bool      `json:"is_blacklisted"`
	Domain        Domain    `json:"domain"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
