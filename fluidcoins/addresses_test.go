package fluidcoins

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Addresses_Create(t *testing.T) {
	aTests := []struct {
		Name               string
		ExpectedModel      Address
		ExpectedMethod     string
		ExpectedRequestUrl string
		CreateOptions      AddressCreateOptions
	}{
		{
			Name:               "ok-values",
			ExpectedModel:      Address{},
			ExpectedMethod:     http.MethodPost,
			ExpectedRequestUrl: fmt.Sprintf("%saddress", defaultBaseURL),
		},
	}

	for _, tt := range aTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(fetchAddressResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Address: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			addr, resp, err := c.Address.Create(context.TODO(), &tt.CreateOptions)

			require.NoError(t, err)
			require.Equal(t, addr, &tt.ExpectedModel)
			require.Equal(t, resp.Request.Method, http.MethodPost)
			require.Equal(t, resp.Request.URL.String(), tt.ExpectedRequestUrl)

		})

	}
}

func Test_Addresses_Fetch_By_Reference(t *testing.T) {
	aTests := []struct {
		Name               string
		ExpectedModel      Address
		ExpectedMethod     string
		ExpectedRequestUrl string
		Reference          string
		WantErr            bool
	}{
		{
			Name:               "ok-values",
			ExpectedModel:      Address{},
			ExpectedMethod:     http.MethodGet,
			ExpectedRequestUrl: fmt.Sprintf("%saddress/%s", defaultBaseURL, "reference"),
			Reference:          "reference",
		},
		{
			Name:               "empty-reference",
			ExpectedModel:      Address{},
			ExpectedMethod:     http.MethodGet,
			ExpectedRequestUrl: fmt.Sprintf("%saddress/%s", defaultBaseURL, ""),
			WantErr:            true,
		},
	}

	for _, tt := range aTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(fetchAddressResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Address: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			addr, resp, err := c.Address.FetchByReference(context.TODO(), tt.Reference)

			if tt.WantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, addr, &tt.ExpectedModel)
				require.Equal(t, resp.Request.Method, http.MethodGet)
				require.Equal(t, resp.Request.URL.String(), tt.ExpectedRequestUrl)
			}

		})

	}
}

func Test_Addresses_Fetch_Transactions(t *testing.T) {
	aTests := []struct {
		Name               string
		ExpectedModels     []*AddressTransaction
		ExpectedMethod     string
		ExpectedRequestUrl string
		Reference          string
		FetchOptions       *FetchAddressTransactionsOptions
		WantErr            bool
	}{
		{
			Name:               "with-defaults",
			ExpectedModels:     make([]*AddressTransaction, 2),
			ExpectedMethod:     http.MethodGet,
			Reference:          "reference",
			ExpectedRequestUrl: fmt.Sprintf("%saddress/%s/transactions?page=%d&per_page=%d", defaultBaseURL, "reference", 1, 20),
			FetchOptions:       &FetchAddressTransactionsOptions{},
		},
		{
			Name:               "ok-values",
			ExpectedModels:     make([]*AddressTransaction, 2),
			ExpectedMethod:     http.MethodGet,
			Reference:          "reference",
			ExpectedRequestUrl: fmt.Sprintf("%saddress/%s/transactions?page=%d&per_page=%d", defaultBaseURL, "reference", 2, 10),
			FetchOptions: &FetchAddressTransactionsOptions{
				Perpage: 10,
				Page:    2,
			},
		},
		{
			Name:               "empty-reference",
			ExpectedModels:     make([]*AddressTransaction, 2),
			ExpectedMethod:     http.MethodGet,
			ExpectedRequestUrl: fmt.Sprintf("%saddress/%s/transactions?page=%d&per_page=%d", defaultBaseURL, "", 1, 20),
			FetchOptions:       &FetchAddressTransactionsOptions{},
			WantErr:            true,
		},
	}

	for _, tt := range aTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(fetchAddressTransactionsResponse{
				apiStatusWithMeta: apiStatusWithMeta{
					apiStatus: apiStatus{
						Status:  true,
						Message: "Successful",
					},
					Meta: ApiStatusMeta{
						Paging: PagingMeta{
							Page:    1,
							Perpage: 20,
							Total:   100,
						},
					},
				},
				Transactions: tt.ExpectedModels,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			trans, resp, err := c.Address.FetchTransactionsByAddressReference(context.TODO(), tt.Reference, tt.FetchOptions)

			if tt.WantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, trans, tt.ExpectedModels)
				require.Equal(t, resp.Request.Method, http.MethodGet)
				require.Equal(t, resp.Request.URL.String(), tt.ExpectedRequestUrl)
			}

		})

	}
}

var nilAddress *Address

func Test_Addresses_Failures_Create(t *testing.T) {
	aFTests := []struct {
		Name            string
		StatusCode      uint
		ApiErrorMessage string
		CreateOptions   AddressCreateOptions
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

	for _, tt := range aFTests {
		h := &http.Client{
			Timeout:   time.Second * 5,
			Transport: NewRoundTripperWithFailureResponse(tt.StatusCode, tt.ApiErrorMessage),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			addr, resp, err := c.Address.Create(context.TODO(), &tt.CreateOptions)

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, addr, nilAddress)
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Addresses_Failures_Fetch_By_Reference(t *testing.T) {
	aFTests := []struct {
		Name            string
		StatusCode      uint
		ApiErrorMessage string
		Reference       string
	}{
		{
			Name:            "Bad-Request",
			StatusCode:      400,
			ApiErrorMessage: "Bad Filter",
			Reference:       "refernece",
		},
		{
			Name:            "Unauthroized",
			StatusCode:      401,
			ApiErrorMessage: "Bad Api Key",
			Reference:       "refernece",
		}, {
			Name:            "Internal Server Error",
			StatusCode:      500,
			ApiErrorMessage: "Internal Server Error",
			Reference:       "refernece",
		},
	}

	for _, tt := range aFTests {
		h := &http.Client{
			Timeout:   time.Second * 5,
			Transport: NewRoundTripperWithFailureResponse(tt.StatusCode, tt.ApiErrorMessage),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			addr, resp, err := c.Address.FetchByReference(context.TODO(), tt.Reference)

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, addr, nilAddress)
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Addresses_Failures_Fetch_Transactions(t *testing.T) {
	aFTests := []struct {
		Name            string
		StatusCode      uint
		ApiErrorMessage string
		Reference       string
	}{
		{
			Name:            "Bad-Request",
			StatusCode:      400,
			ApiErrorMessage: "Bad Filter",
			Reference:       "refernece",
		},
		{
			Name:            "Unauthroized",
			StatusCode:      401,
			ApiErrorMessage: "Bad Api Key",
			Reference:       "refernece",
		}, {
			Name:            "Internal Server Error",
			StatusCode:      500,
			ApiErrorMessage: "Internal Server Error",
			Reference:       "refernece",
		},
	}

	for _, tt := range aFTests {
		h := &http.Client{
			Timeout:   time.Second * 5,
			Transport: NewRoundTripperWithFailureResponse(tt.StatusCode, tt.ApiErrorMessage),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			trans, resp, err := c.Address.FetchTransactionsByAddressReference(context.TODO(), tt.Reference, &FetchAddressTransactionsOptions{})

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, trans, []*AddressTransaction(nil))
			require.Equal(t, resp, nilResp)
		})

	}
}
