package fluidcoins

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// TransactionService is an helper that communicates with the transactions API
type TransactionService service

// Transaction represents a transaction that occured via the payment widget
type Transaction struct {
	ID            uuid.UUID `json:"id"`
	Amount        Amount    `json:"amount"`
	Domain        Domain    `json:"domain"`
	CustomerID    uuid.UUID `json:"-"`
	Customer      *Customer `json:"customer"`
	MerchantID    uuid.UUID `json:"-"`
	Merchant      *Merchant `json:"-" pg:"rel:has-one"`
	PaymentLinkID uuid.UUID `json:"payment_link_id" `
	Reference     string    `json:"reference"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Merchant struct {
	ID           uuid.UUID `json:"id"`
	BusinessName string    `json:"business_name"`
	Description  string    `json:"description"`
	IndustryID   uuid.UUID `json:"industry_id"`

	CountryID uuid.UUID `json:"country_id"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}

// GetByID fetches a transaction.
func (t *TransactionService) GetByID(ctx context.Context, id string) {
}
