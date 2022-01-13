package fluidcoins

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fluidcoins/client-go/util"
	"github.com/google/uuid"
)

// TransactionService is an helper that communicates with the transactions API
type TransactionService service

// Transaction represents a transaction that occured via the payment widget
type Transaction struct {
	apiStatus
	Transaction struct {
		ID            uuid.UUID `json:"id"`
		Amount        Amount    `json:"amount"`
		Domain        Domain    `json:"domain"`
		CustomerID    uuid.UUID `json:"-"`
		Customer      *Customer `json:"customer"`
		MerchantID    uuid.UUID `json:"-"`
		PaymentLinkID uuid.UUID `json:"payment_link_id"`
		Reference     string    `json:"reference"`
		Status        string    `json:"status"`

		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"transaction"`
}

type Merchant struct {
	ID           uuid.UUID `json:"id"`
	BusinessName string    `json:"business_name"`
	Description  string    `json:"description"`
	IndustryID   uuid.UUID `json:"industry_id"`

	CountryID uuid.UUID `json:"country_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetByID fetches a transaction by the provided ID.
func (t *TransactionService) GetByID(ctx context.Context, id string) (*Transaction, *Response, error) {

	if util.IsStringEmpty(id) {
		return nil, nil, errors.New("please provide an ID")
	}

	req, err := t.c.NewRequest(http.MethodGet, fmt.Sprintf("transactions/%s", id), nil)
	if err != nil {
		return nil, nil, err
	}

	var trans = new(Transaction)

	resp, err := t.c.Do(ctx, req, trans)
	if err != nil {
		return nil, nil, err
	}

	return trans, resp, nil
}
