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
	responseBody listCurrencyResponse
}

func (rt *roundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	rr, w := io.Pipe()

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(rt.responseBody)
	resp := &http.Response{
		Request:    r,
		StatusCode: 200,
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
		responseBody: resp,
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

			fmt.Printf("resp is %v \n", resp)
			require.NoError(t, err)
			require.Equal(t, currs, tt.ExpectedModels)
			require.Equal(t, resp.Request.Method, http.MethodGet)
			require.Equal(t, resp.Request.URL.String(), tt.ExpectedRequestUrl)
		})

	}
}
