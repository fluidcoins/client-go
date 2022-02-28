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

type TransactionStatus string

const (
	Successful TransactionStatus = "successful"
	Pending    TransactionStatus = "pending"
	Expired    TransactionStatus = "expired"
	Overpaid   TransactionStatus = "overpaid"
	Underpaid  TransactionStatus = "underpaid"
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
	PaymentLinkID uuid.UUID `json:"payment_link_id"`
	Reference     string    `json:"reference"`
	Status        string    `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type transactionResponse struct {
	apiStatus
	Transaction Transaction `json:"transaction"`
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

	var trans = new(transactionResponse)

	resp, err := t.c.Do(ctx, req, trans)
	if err != nil {
		return nil, nil, err
	}

	return &trans.Transaction, resp, nil
}

type TransactionListOptions struct {
	Status  TransactionStatus `url:"status"`
	Perpage int64             `url:"per_page"`
	Page    int64             `url:"page"`
}

type listTransactionResponse struct {
	apiStatus
	Transactions []*Transaction
}

// List fetches all transactions. It uses the TransactionListOptions to filter
// the results. We default to 20 items per page and the first page
func (t *TransactionService) List(ctx context.Context, opts *TransactionListOptions) ([]*Transaction, *Response, error) {

	if opts.Perpage <= 0 {
		opts.Perpage = 20
	}

	if opts.Page <= 0 {
		opts.Page = 1
	}

	s, err := addQueryString("transactions", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := t.c.NewRequest(http.MethodGet, s, nil)
	if err != nil {
		return nil, nil, err
	}

	var trans = new(listTransactionResponse)

	resp, err := t.c.Do(ctx, req, &trans)
	if err != nil {
		return nil, nil, err
	}

	return trans.Transactions, resp, nil
}
