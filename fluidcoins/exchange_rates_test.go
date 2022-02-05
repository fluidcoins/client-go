package fluidcoins

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_ExchangeRates(t *testing.T) {
	eTests := []struct {
		Name               string
		ExpectedModels     []*ExchangeRate
		ExpectedMethod     string
		ExpectedRequestUrl string
		ListOptions        *ExchangeRateListOptions
	}{
		{
			Name:               "without-filter",
			ExpectedModels:     make([]*ExchangeRate, 2),
			ExpectedMethod:     http.MethodGet,
			ExpectedRequestUrl: fmt.Sprintf("%s%s?from=&to=", defaultBaseURL, "rates"),
			ListOptions:        &ExchangeRateListOptions{},
		},
		{
			Name:               "with-filter",
			ExpectedModels:     make([]*ExchangeRate, 1),
			ExpectedMethod:     http.MethodGet,
			ExpectedRequestUrl: fmt.Sprintf("%s%s?from=usdc&to=busb", defaultBaseURL, "rates"),
			ListOptions: &ExchangeRateListOptions{
				From: "usdc",
				To:   "busb",
			},
		},
	}

	for _, tt := range eTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(listExchangeRateResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				ExchangeRates: tt.ExpectedModels,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			rates, resp, err := c.ExchangeRate.List(context.TODO(), tt.ListOptions)

			require.NoError(t, err)
			require.Equal(t, rates, tt.ExpectedModels)
			require.Equal(t, resp.Request.Method, http.MethodGet)
			require.Equal(t, resp.Request.URL.String(), tt.ExpectedRequestUrl)
		})

	}
}

func Test_ExchangeRates_Failures(t *testing.T) {
	eFTests := []struct {
		Name            string
		StatusCode      uint
		ApiErrorMessage string
	}{
		{
			Name:            "Bad-Request",
			StatusCode:      400,
			ApiErrorMessage: "Bad Filter",
		},
		{
			Name:            "Unauthroized",
			StatusCode:      401,
			ApiErrorMessage: "Bad Api Key",
		}, {
			Name:            "Internal Server Error",
			StatusCode:      500,
			ApiErrorMessage: "Internal Server Error",
		},
	}

	for _, tt := range eFTests {
		h := &http.Client{
			Timeout:   time.Second * 5,
			Transport: NewRoundTripperWithFailureResponse(tt.StatusCode, tt.ApiErrorMessage),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			rates, resp, err := c.ExchangeRate.List(context.TODO(), &ExchangeRateListOptions{
				From: "busd",
			})

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, rates, []*ExchangeRate(nil))
			require.Equal(t, resp, nilResp)
		})

	}
}
