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

type AddressesService service

type Address struct {
	ID         uuid.UUID `json:"id"`
	Address    string    `json:"address"`
	Coin       Coin      `json:"coin"`
	Domain     string    `json:"domain"`
	MerchantID string    `json:"merchant_id"`
	// Refernece is the unique ID for this address. You can use this id to fetch the address again
	Reference string          `json:"reference"`
	Metadata  AddressMetadata `json:"metadata"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Coin struct {
	ID   uuid.UUID `json:"id"`
	Code string    `json:"code"`

	HumanReadableName string `json:"human_readable_name"`
	// This is a base64 encoded image not an actual url resource
	Image string `json:"image"`
	// IsAvailableOnWidget determines if this currency coin is a payment option on our widget
	IsAvailableOnWidget bool `json:"is_available_on_widget"`
	IsFiat              bool `json:"is_fiat"`
	IsStableCoin        bool `json:"is_stable_coin"`
	// Multiplier is what should be used to do the currency divisions
	Multiplier int64    `json:"multiplier"`
	Networks   []string `json:"networks"`
	HasTestNet bool     `json:"has_test_net"`

	Metadata CoinMetadata `json:"metadata"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddressMetadata struct {
	// XLM requires a memo
	Memo string `json:"memo"`
	// XRP addresses must have a destination tag
	DestinationTag string `json:"destination_tag"`
	// can be erc20, trc20 or bsc
	// if empty, then assume the default chain
	// in cases such as btc,ltc and what not
	Network string `json:"network"`
}

type CoinMetadata struct {
	MinDepositAmount int64 `json:"minimum_deposit_amount"`
	MinPayoutAmount  int64 `json:"minimum_payout_amount"`
	MinWidgetAmount  int64 `json:"minimum_widget_amount"`
	PayoutFee        int64 `json:"payout_fee"`
}

type AddressTransaction struct {
	ID               uuid.UUID                  `json:"id"`
	Amount           Amount                     `json:"amount"`
	Domain           Domain                     `json:"domain"`
	Coin             Coin                       `json:"coin"`
	Hash             string                     `json:"hash"`
	Metadata         AddressTransactionMetadata `json:"metadata"`
	PaymentAddressID string                     `json:"payment_address_id"`
	Reference        string                     `json:"reference"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AddressTransactionMetadata struct {
	BlockHash   string `json:"block_hash"`
	BlockHeight int64  `json:"block_height"`
	// The sender address. This field will not be available for all blockchains
	From string `json:"from"`
}

type AddressCreateOptions struct {
	// Code for the coin you want to generate an address for. e.g (XLM, USDC)
	Code string `json:"code"`
	// Optional: You can use a supported network like erc20,trc20,bep20 and others. If left empty, it will use the default network of the coin
	Network string `json:"network"`
}

type createAddressResponse struct {
	apiStatus
	Address Address `json:"address"`
}

type fetchAddressResponse struct {
	apiStatus
	Address Address `json:"address"`
}

type fetchAddressTransactionsResponse struct {
	apiStatusWithMeta
	Transactions []*AddressTransaction
}

type FetchAddressTransactionsOptions struct {
	// Number to items to return. Defaults to 20 items
	Perpage int64 `url:"per_page"`
	// Page to query data from. Defaults to 1
	Page int64 `url:"page"`
}

// Create creates a new address
func (a *AddressesService) Create(ctx context.Context, opts *AddressCreateOptions) (*Address, *Response, error) {
	req, err := a.c.NewRequest(http.MethodPost, "address", opts)

	if err != nil {
		return nil, nil, err
	}

	var addr = new(createAddressResponse)

	resp, err := a.c.Do(ctx, req, &addr)

	if err != nil {
		return nil, nil, err
	}

	return &addr.Address, resp, nil
}

// FetchByReference fetches an Address by the provided reference
func (a *AddressesService) FetchByReference(ctx context.Context, reference string) (*Address, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := a.c.NewRequest(http.MethodGet, fmt.Sprintf("address/%s", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var addr = new(fetchAddressResponse)

	resp, err := a.c.Do(ctx, req, addr)
	if err != nil {
		return nil, nil, err
	}

	return &addr.Address, resp, nil
}

// FetchTransactionsByAddressReference fetches transactions on an address (by its reference)
func (a *AddressesService) FetchTransactionsByAddressReference(ctx context.Context, reference string, opts *FetchAddressTransactionsOptions) ([]*AddressTransaction, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	if opts.Perpage <= 0 {
		opts.Perpage = 20
	}

	if opts.Page <= 0 {
		opts.Page = 1
	}

	var url string = fmt.Sprintf("address/%s/transactions", reference)

	s, err := addQueryString(url, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := a.c.NewRequest(http.MethodGet, s, nil)
	if err != nil {
		return nil, nil, err
	}

	var trans = new(fetchAddressTransactionsResponse)

	resp, err := a.c.Do(ctx, req, &trans)
	if err != nil {
		return nil, nil, err
	}

	return trans.Transactions, resp, nil
}
