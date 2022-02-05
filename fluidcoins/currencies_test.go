package fluidcoins

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type roundTripper struct {
	successResponseBody listCurrencyResponse
	failureResponseBody apiStatus
	statusCode          uint
	isFailure           bool
}

var nilResp *Response

func (rt *roundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	rr, w := io.Pipe()

	reqBodyBytes := new(bytes.Buffer)

	if rt.isFailure {
		json.NewEncoder(reqBodyBytes).Encode(rt.failureResponseBody)
	} else {
		json.NewEncoder(reqBodyBytes).Encode(rt.successResponseBody)
	}

	resp := &http.Response{
		Request:    r,
		StatusCode: int(rt.statusCode),
		Body:       rr,
	}

	go func() {
		defer w.Close()
		io.Copy(w, reqBodyBytes)
	}()

	return resp, nil
}

func NewRoundTripperWithResponse(resp listCurrencyResponse) *roundTripper {
	return &roundTripper{
		successResponseBody: resp,
		statusCode:          200,
	}
}

func NewRoundTripperWithFailureResponse(statusCode uint, message string) *roundTripper {
	return &roundTripper{
		failureResponseBody: apiStatus{
			Status:  false,
			Message: message,
		},
		statusCode: statusCode,
		isFailure:  true,
	}
}

func Test_Currencies(t *testing.T) {
	cTests := []struct {
		Name               string
		ExpectedModels     []*Currency
		ExpectedMethod     string
		ExpectedRequestUrl string
		ListOptions        *CurrencyListOptions
	}{
		{
			Name:               "without-testnet-network-filter",
			ExpectedModels:     make([]*Currency, 2),
			ExpectedMethod:     http.MethodGet,
			ExpectedRequestUrl: fmt.Sprint("", defaultBaseURL, "currencies"),
			ListOptions:        &CurrencyListOptions{},
		},
		{
			Name:               "with-testnet-network-filter-true",
			ExpectedModels:     make([]*Currency, 1),
			ExpectedMethod:     http.MethodGet,
			ExpectedRequestUrl: fmt.Sprintf("%s%s?FilterTestNet=true&test_net_only=true", defaultBaseURL, "currencies"),
			ListOptions: &CurrencyListOptions{
				FilterTestNet:      true,
				TestNetNetworkOnly: true,
			},
		},
		{
			Name:               "with-testnet-network-filter-false",
			ExpectedModels:     make([]*Currency, 4),
			ExpectedMethod:     http.MethodGet,
			ExpectedRequestUrl: fmt.Sprintf("%s%s?FilterTestNet=true&test_net_only=false", defaultBaseURL, "currencies"),
			ListOptions: &CurrencyListOptions{
				FilterTestNet:      true,
				TestNetNetworkOnly: false,
			},
		},
	}

	for _, tt := range cTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(listCurrencyResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Currencies: tt.ExpectedModels,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			currs, resp, err := c.Currency.List(context.TODO(), tt.ListOptions)

			require.NoError(t, err)
			require.Equal(t, currs, tt.ExpectedModels)
			require.Equal(t, resp.Request.Method, http.MethodGet)
			require.Equal(t, resp.Request.URL.String(), tt.ExpectedRequestUrl)
		})

	}
}

func Test_Currencies_Errors(t *testing.T) {
	cFTests := []struct {
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

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout:   time.Second * 5,
			Transport: NewRoundTripperWithFailureResponse(tt.StatusCode, tt.ApiErrorMessage),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			currs, resp, err := c.Currency.List(context.TODO(), &CurrencyListOptions{})

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, currs, []*Currency(nil))
			require.Equal(t, resp, nilResp)
		})

	}
}
