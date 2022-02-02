package fluidcoins

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// CurrencyService is a helper that communicates with the currencies API
type CurrencyService service

// Currecny represents a currency (FIAT or Coin)
type Currency struct {
	ID                  uuid.UUID        `json:"id"`
	HumanReadableName   string           `json:"human_readable_name"`
	Code                string           `json:"code"`
	IsFiat              bool             `json:"is_fiat"`
	Image               string           `json:"image"`
	IsAvailableOnWidget bool             `json:"is_available_on_widget"`
	IsStableCoin        bool             `json:"is_stable_coin"`
	Multiplier          int64            `json:"multiplier"`
	HasTestNet          bool             `json:"has_test_net"`
	Metadata            CurrencyMetadata `json:"metadata"`
	Networks            []string         `json:"networks"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CurrencyMetadata struct {
	MinDepositAmount int64 `json:"minimum_deposit_amount"`
	MinPayoutAmount  int64 `json:"minimum_payout_amount"`
	PayoutFee        int64 `json:"payout_fee"`
}

type listCurrencyResponse struct {
	apiStatus
	Currencies []*Currency `json:"currencies"`
}

type CurrencyListOptions struct {
	TestNetNetworkOnly bool `url:"test_net_only"`
	FilterTestNet      bool
}

// List fetches all currencies. It uses the CurrencyListOptions to filter
// the results.
func (c *CurrencyService) List(ctx context.Context, opts *CurrencyListOptions) ([]*Currency, *Response, error) {
	var s string
	var err error

	if opts.FilterTestNet {
		s, err = addQueryString("currencies", opts)
	} else {
		s, err = addQueryString("currencies", nil)
	}

	if err != nil {
		return nil, nil, err
	}

	req, err := c.c.NewRequest(http.MethodGet, s, nil)
	if err != nil {
		return nil, nil, err
	}

	var currs = new(listCurrencyResponse)

	resp, err := c.c.Do(ctx, req, &currs)
	if err != nil {
		return nil, nil, err
	}

	return currs.Currencies, resp, nil
}
