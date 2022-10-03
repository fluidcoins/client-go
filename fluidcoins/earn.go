package fluidcoins

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type EarnService service

type EarnInterestRate struct {
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	ID        uuid.UUID `json:"id"`
}

type earnInterestRespons struct {
	Rates []EarnInterestRate `json:"rates"`
}

type EarnInterestRateTimeframe string

const (
	TodayInterestRateTimeFrame  EarnInterestRateTimeframe = "today"
	WeeklyInterestRateTimeFrame                           = "week"
)

type FetchearnInterestRateOptions struct {
	Timeframe EarnInterestRateTimeframe
}

func (e *EarnService) InterestRates(ctx context.Context,
	opts *FetchearnInterestRateOptions) ([]EarnInterestRate, *Response, error) {

	if opts == nil {
		opts = &FetchearnInterestRateOptions{
			Timeframe: TodayInterestRateTimeFrame,
		}
	}

	req, err := e.c.NewRequest(http.MethodGet, fmt.Sprintf("earn/rates?timeframe=%s", opts.Timeframe), nil)
	if err != nil {
		return nil, nil, err
	}

	var response = new(earnInterestRespons)

	resp, err := e.c.Do(ctx, req, &response)
	if err != nil {
		return nil, nil, err
	}

	return response.Rates, resp, nil
}
