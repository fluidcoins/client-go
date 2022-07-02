package fluidcoins

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/fluidcoins/client-go/util"
	"github.com/google/uuid"
)

type PayoutService service

type PayoutAccount struct {
	model
	MerchantID uuid.UUID    `json:"merchant_id"`
	CoinID     string       `json:"coin_id"`
	Domain     string       `json:"domain"`
	Bank       PayoutBank   `json:"bank"`
	Crypto     PayoutCrypto `json:"crypto"`

	PayoutType string `json:"payout_type"`
	Reference  string `json:"reference"`
}

type Payout struct {
	model
	Amount     int64          `json:"amount"`
	MerchantID uuid.UUID      `json:"merchant_id"`
	CoinID     string         `json:"coin_id"`
	Coin       PayoutCoin     `json:"coin"`
	Domain     string         `json:"domain"`
	Metadata   PayoutMetadata `json:"metadata"`
	Reference  string         `json:"reference"`

	// This will be automatically updated based on payout attempts
	Status int64 `json:"status"`
}

type PayoutBank struct {
	AccountName   string      `json:"account_name"`
	AccountNumber string      `json:"account_number"`
	BankCode      string      `json:"bank_code"`
	BankName      string      `json:"bank_name"`
	Currency      string      `json:"currency"`
	Metadata      interface{} `json:"metadata"`
}

type PayoutCrypto struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Label   string `json:"label"`
	Network string `json:"network"`
}

type PayoutCoin struct {
	model
	Code string `json:"code"`
	// This is a url link ( hosted by us ) to the logo of the coin.
	Image      string `json:"image"`
	HasTestNet bool   `json:"has_test_net"`

	HumanReadableName string `json:"human_readable_name"`
	// IsAvailableOnWidget determines if this currency coin is a payment option on our widget
	IsAvailableOnWidget bool `json:"is_available_on_widget"`

	IsFiat       bool `json:"is_fiat"`
	IsStableCoin bool `json:"is_stable_coin"`
}

type PayoutMetadata struct {
	Fee    PayoutFee    `json:"fee"`
	Crypto PayoutCrypto `json:"crypto"`
	Bank   PayoutBank   `json:"bank"`
}

type PayoutFee struct {
	Amount int64 `json:"amount"`
	Coin   int64 `json:"coin"`
}

type PayoutCreateBankOptions struct {
	// Bank account number
	AccountNumber string `json:"account_number" url:"account"`
	// Sort code of the bank
	BankCode string `json:"bank_code" url:"bank_code"`
}

type RequestPayoutOptions struct {
	// Amount to be sent to the recipient. Please note that this is should
	// be in the currency's lowest denomination.
	// 1,000 Naira would be 100000 while 1 BTC would be 100 million satoshis
	// The currency is automatically retrieved from the payout account
	Amount int64 `json:"amount"`
	// The reference of the payout account, PAY_ACCT_XYZ
	Recipient string `json:"recipient"`
}

type ResolveBankInfoOptions PayoutCreateBankOptions

type ResolveBankInfo struct {
	Name string `json:"name"`
}

type ListBankOptions struct {
	// Optional. Defaults to NG
	Country string `url:"country"`
}

type Bank struct {
	Code    string `json:"code"`
	Country string `json:"country"`
	Name    string `json:"name"`
}

type PayoutCreateCryptoOptions struct {
	Address string `json:"address"`
	Label   string `json:"label"`
	Network string `json:"network"`
}

type PayoutAccountCreateOptions struct {
	Currency string                    `json:"currency"`
	Bank     PayoutCreateBankOptions   `json:"bank"`
	Crypto   PayoutCreateCryptoOptions `json:"crypto"`
}

type createPayoutAccountResponse struct {
	apiStatus
	Payout PayoutAccount `json:"payout"`
}

type requestPayoutResponse struct {
	apiStatus
	Payout Payout `json:"payout"`
}

type fetchPayoutResponse requestPayoutResponse

type resolveBankInfoResponse struct {
	apiStatus
	Account ResolveBankInfo `json:"account"`
}

type listBanksResponse struct {
	apiStatus
	Banks []*Bank `json:"banks"`
}

// Create a new payout account
func (p *PayoutService) CreatePayoutAccount(ctx context.Context, opts *PayoutAccountCreateOptions) (*PayoutAccount, *Response, error) {
	req, err := p.c.NewRequest(http.MethodPost, "payouts/accounts", opts)

	if err != nil {
		return nil, nil, err
	}

	var payout = new(createPayoutAccountResponse)

	resp, err := p.c.Do(ctx, req, &payout)

	if err != nil {
		return nil, nil, err
	}

	return &payout.Payout, resp, nil
}

// Request for a payout
func (p *PayoutService) RequestPayout(ctx context.Context, opts *RequestPayoutOptions) (*Payout, *Response, error) {
	req, err := p.c.NewRequest(http.MethodPost, "payouts", opts)

	if err != nil {
		return nil, nil, err
	}

	var payout = new(requestPayoutResponse)

	resp, err := p.c.Do(ctx, req, &payout)

	if err != nil {
		return nil, nil, err
	}

	return &payout.Payout, resp, nil
}

// Fetch the details of a payout
func (p *PayoutService) PayoutDetailsByReference(ctx context.Context, reference string) (*Payout, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := p.c.NewRequest(http.MethodGet, fmt.Sprintf("payouts/%s", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var payout = new(fetchPayoutResponse)

	resp, err := p.c.Do(ctx, req, payout)
	if err != nil {
		return nil, nil, err
	}

	return &payout.Payout, resp, nil
}

// Cancel a payout
func (p *PayoutService) Cancel(ctx context.Context, reference string) (*Payout, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := p.c.NewRequest(http.MethodDelete, fmt.Sprintf("payouts/%s", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var payout = new(fetchPayoutResponse)

	resp, err := p.c.Do(ctx, req, payout)
	if err != nil {
		return nil, nil, err
	}

	return &payout.Payout, resp, nil
}

// ResolveBankAccount bank account
func (p *PayoutService) ResolveBankAccount(ctx context.Context, opts *ResolveBankInfoOptions) (*ResolveBankInfo, *Response, error) {
	s, err := addQueryString("payouts/accounts/banks/resolve", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := p.c.NewRequest(http.MethodGet, s, opts)
	if err != nil {
		return nil, nil, err
	}

	var resolve = new(resolveBankInfoResponse)

	resp, err := p.c.Do(ctx, req, resolve)
	if err != nil {
		return nil, nil, err
	}

	return &resolve.Account, resp, nil
}

// List banks
func (p *PayoutService) ListBanks(ctx context.Context, opts *ListBankOptions) ([]*Bank, *Response, error) {
	s, err := addQueryString("payouts/accounts/banks", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := p.c.NewRequest(http.MethodGet, s, nil)
	if err != nil {
		return nil, nil, err
	}

	var banks = new(listBanksResponse)

	resp, err := p.c.Do(ctx, req, &banks)
	if err != nil {
		return nil, nil, err
	}

	return banks.Banks, resp, nil
}
