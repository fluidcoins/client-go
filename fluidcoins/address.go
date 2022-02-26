package fluidcoins

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// AddressService is an helper that communicates with the address API
type AddressService service

// Address is used to represent the structure of a created merchant payment address
type Address struct {
	ID            uuid.UUID `json:"id"`
	Reference     string    `json:"reference"`
	Address       string    `json:"address"`
	Amount        Amount    `json:"amount"`
	Domain        Domain    `json:"domain"`
	MerchantID    uuid.UUID `json:"merchant_id"`
	CoinID        uuid.UUID `json:"coin_id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	ProviderID    string    `json:"provider_id"`
	Provider      string    `json:"provider"`
	AddressIndex  int       `json:"address_index"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAddressOptions struct {
	Code    string `json:"code"`
	Network string `json:"network"`
}

type createAddressResponse struct {
	apiStatus
	Address Address `json:"address"`
}

// Create creates an address. It uses the createAddressOptions to specify
// the coin code and network.
func (s *AddressService) Create(ctx context.Context, opts *CreateAddressOptions) (*Address, *Response, error) {

	bt, err := json.Marshal(&opts)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.c.NewRequest(http.MethodPost, "address", bt)
	if err != nil {
		return nil, nil, err
	}

	var addr = new(createAddressResponse)

	res, err := s.c.Do(ctx, req, addr)
	if err != nil {
		return nil, nil, err
	}

	return &addr.Address, res, err
}

type getAddressResponse struct {
	apiStatus
	Address Address `json:"address"`
}

// GetByReference fetches a merchant address by the provided reference string. e.g ADDR_xy
func (s *AddressService) GetByReference(ctx context.Context, ref string) (*Address, *Response, error) {

	req, err := s.c.NewRequest(http.MethodGet, fmt.Sprintf("address/%s", ref), nil)
	if err != nil {
		return nil, nil, err
	}

	var addr = new(getAddressResponse)
	res, err := s.c.Do(ctx, req, &addr)
	if err != nil {
		return nil, nil, err
	}

	return &addr.Address, res, err

}

type ListAddressOptions struct {
	CoinID  uuid.UUID `url:"coin_id"`
	Perpage int64     `url:"per_page"`
	Page    int64     `url:"page"`
}

type listAddressResponse struct {
	apiStatus
	Addresses []*Address
}

// List fetches all addresses. It uses the AddressListOptions to filter
// the results.
func (s *AddressService) List(ctx context.Context, opts ListAddressOptions) ([]*Address, *Response, error) {

	str, err := addQueryString("address", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.c.NewRequest(http.MethodGet, str, nil)
	if err != nil {
		return nil, nil, err
	}

	var addrs = new(listAddressResponse)

	res, err := s.c.Do(ctx, req, &addrs)
	if err != nil {
		return nil, nil, err
	}

	return addrs.Addresses, res, nil

}

type AddressTransactionsOptions struct {
	Perpage int64 `url:"per_page"`
	Page    int64 `url:"page"`
}

type addressTransactionsResponse struct {
	apiStatus
	Transactions []*Transaction
}

// List fetches all transactions for a particulat address. It uses the AddressTransactionsOptions to filter
// the results. We default to 20 items per page and the first page
func (s *AddressService) ListTransactions(ctx context.Context, ref string, opts *AddressTransactionsOptions) ([]*Transaction, *Response, error) {

	if opts.Perpage <= 0 {
		opts.Perpage = 20
	}

	if opts.Page <= 0 {
		opts.Page = 1
	}

	str, err := addQueryString(fmt.Sprintf("address/%s/transactions", ref), opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.c.NewRequest(http.MethodGet, str, nil)
	if err != nil {
		return nil, nil, err
	}

	var trans = new(addressTransactionsResponse)

	res, err := s.c.Do(ctx, req, trans)
	if err != nil {
		return nil, nil, err
	}

	return trans.Transactions, res, nil

}
