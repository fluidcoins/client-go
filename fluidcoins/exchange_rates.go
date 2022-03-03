package fluidcoins

import (
	"context"
	"net/http"
)

type ExchangeRateService service

type ExchangeRate struct {
	Pair                 string                       `json:"pair"`
	From                 ExchangeRateCurrencyMetadata `json:"from"`
	To                   ExchangeRateCurrencyMetadata `json:"to"`
	Rate                 int64                        `json:"rate"`
	BuyRate              int64                        `json:"buy_rate"`
	SaleRate             int64                        `json:"sale_rate"`
	HumanReadableBuyRate int64                        `json:"human_readable_buy_rate"`
	HumanReadableRate    int64                        `json:"human_readable_sale_rate"`
	HumanReadableAmount  int64                        `json:"human_readable_amount"`
}

type ExchangeRateCurrencyMetadata struct {
	Code string `json:"code"`
}

type listExchangeRateResponse struct {
	apiStatus
	ExchangeRates []*ExchangeRate `json:"rates"`
}

type ExchangeRateListOptions struct {
	From string `url:"from"`
	To   string `url:"to"`
}

// List fetches a list of the currenct exchange rates of all supported fiat currencies
// It uses the ExchangeRateListOptions to filter the results.
func (e *ExchangeRateService) List(ctx context.Context, opts *ExchangeRateListOptions) ([]*ExchangeRate, *Response, error) {
	s, err := addQueryString("rates", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := e.c.NewRequest(http.MethodGet, s, nil)
	if err != nil {
		return nil, nil, err
	}

	var rates = new(listExchangeRateResponse)

	resp, err := e.c.Do(ctx, req, &rates)
	if err != nil {
		return nil, nil, err
	}

	return rates.ExchangeRates, resp, nil
}
