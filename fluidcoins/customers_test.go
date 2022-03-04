package fluidcoins

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var nilCustomer *Customer

func Test_Customers_Failures_Create(t *testing.T) {
	cFTests := []struct {
		Name            string
		StatusCode      uint
		ApiErrorMessage string
		CreateOptions   CustomerCreateOptions
	}{
		{
			Name:            "Bad-Request",
			StatusCode:      400,
			ApiErrorMessage: "Bad Filter",
			CreateOptions: CustomerCreateOptions{
				Email:    "foo@example.com",
				FullName: "Foo Bar",
				Phone: CustomerCreatePhoneOptions{
					Code:  "+234",
					Phone: "00001111",
				},
			},
		},
		{
			Name:            "Unauthroized",
			StatusCode:      401,
			ApiErrorMessage: "Bad Api Key",
			CreateOptions:   CustomerCreateOptions{},
		}, {
			Name:            "Internal Server Error",
			StatusCode:      500,
			ApiErrorMessage: "Internal Server Error",
			CreateOptions:   CustomerCreateOptions{},
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout:   time.Second * 5,
			Transport: NewRoundTripperWithFailureResponse(tt.StatusCode, tt.ApiErrorMessage),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			addr, resp, err := c.Customer.Create(context.TODO(), &tt.CreateOptions)

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, addr, nilCustomer)
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Customers_Failures_Edit(t *testing.T) {
	cFTests := []struct {
		Name            string
		StatusCode      uint
		ApiErrorMessage string
		Reference       string
		EditOpts        CustomerCreateOptions
	}{
		{
			Name:            "Bad-Request",
			StatusCode:      400,
			ApiErrorMessage: "Bad Filter",
			Reference:       "refernece",
			EditOpts: CustomerCreateOptions{
				Email:    "foo@example.com",
				FullName: "Foo Bar",
				Phone: CustomerCreatePhoneOptions{
					Code:  "+234",
					Phone: "00001111",
				},
			},
		},
		{
			Name:            "Unauthroized",
			StatusCode:      401,
			ApiErrorMessage: "Bad Api Key",
			Reference:       "refernece",
			EditOpts:        CustomerCreateOptions{},
		}, {
			Name:            "Internal Server Error",
			StatusCode:      500,
			ApiErrorMessage: "Internal Server Error",
			Reference:       "refernece",
			EditOpts:        CustomerCreateOptions{},
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout:   time.Second * 5,
			Transport: NewRoundTripperWithFailureResponse(tt.StatusCode, tt.ApiErrorMessage),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			addr, resp, err := c.Customer.Edit(context.TODO(), tt.Reference, &tt.EditOpts)

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, addr, nilCustomer)
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Customers_Failures_Fetch_By_Reference(t *testing.T) {
	cFTests := []struct {
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

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout:   time.Second * 5,
			Transport: NewRoundTripperWithFailureResponse(tt.StatusCode, tt.ApiErrorMessage),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			cust, resp, err := c.Customer.FetchByReference(context.TODO(), tt.Reference)

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, cust, nilCustomer)
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Customers_Failures_Blacklist_By_Reference(t *testing.T) {
	cFTests := []struct {
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

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout:   time.Second * 5,
			Transport: NewRoundTripperWithFailureResponse(tt.StatusCode, tt.ApiErrorMessage),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			cust, resp, err := c.Customer.BlacklistByReference(context.TODO(), tt.Reference)

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, cust, nilCustomer)
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Customers_Failures_Whitelist_By_Reference(t *testing.T) {
	cFTests := []struct {
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

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout:   time.Second * 5,
			Transport: NewRoundTripperWithFailureResponse(tt.StatusCode, tt.ApiErrorMessage),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			cust, resp, err := c.Customer.WhitelistByReference(context.TODO(), tt.Reference)

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, cust, nilCustomer)
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Customers_Create(t *testing.T) {
	cFTests := []struct {
		Name          string
		ExpectedModel Customer
		CreateOptions CustomerCreateOptions
	}{
		{
			Name:          "ok-values",
			ExpectedModel: Customer{},
			CreateOptions: CustomerCreateOptions{
				Email:    "foo@example.com",
				FullName: "Foo Bar",
				Phone: CustomerCreatePhoneOptions{
					Code:  "+234",
					Phone: "00001111",
				},
			},
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(createCustomerResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Customer: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			cust, resp, err := c.Customer.Create(context.TODO(), &tt.CreateOptions)

			require.NoError(t, err)
			require.Equal(t, cust, &tt.ExpectedModel)
			require.Equal(t, resp.Request.Method, http.MethodPost)
			require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%scustomers", defaultBaseURL))
		})

	}
}

func Test_Customers_Edit(t *testing.T) {
	cFTests := []struct {
		Name          string
		Reference     string
		ExpectedModel Customer
		EditOpts      CustomerCreateOptions
		WantErr       bool
	}{
		{
			Name:          "ok-values",
			Reference:     "refernece",
			ExpectedModel: Customer{},
			EditOpts: CustomerCreateOptions{
				Email:    "foo@example.com",
				FullName: "Foo Bar",
				Phone: CustomerCreatePhoneOptions{
					Code:  "+234",
					Phone: "00001111",
				},
			},
		},
		{
			Name:      "error-no-reference",
			Reference: "",
			EditOpts:  CustomerCreateOptions{},
			WantErr:   true,
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(editCustomerResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Customer: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			cust, resp, err := c.Customer.Edit(context.TODO(), tt.Reference, &tt.EditOpts)

			if tt.WantErr {
				require.Error(t, err)

			} else {

				require.NoError(t, err)
				require.Equal(t, cust, &tt.ExpectedModel)
				require.Equal(t, resp.Request.Method, http.MethodPatch)
				require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%scustomers/%s", defaultBaseURL, tt.Reference))
			}
		})

	}
}

func Test_Customers_Fetch_By_Reference(t *testing.T) {
	cFTests := []struct {
		Name          string
		Reference     string
		WantErr       bool
		ExpectedModel Customer
	}{
		{
			Name:          "ok-values",
			Reference:     "refernece",
			ExpectedModel: Customer{},
		},
		{
			Name:      "error-no-reference",
			Reference: "",
			WantErr:   true,
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(createCustomerResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Customer: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			cust, resp, err := c.Customer.FetchByReference(context.TODO(), tt.Reference)

			if tt.WantErr {
				require.Error(t, err)

			} else {

				require.NoError(t, err)
				require.Equal(t, cust, &tt.ExpectedModel)
				require.Equal(t, resp.Request.Method, http.MethodGet)
				require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%scustomers/%s", defaultBaseURL, tt.Reference))
			}
		})

	}
}

func Test_Customers_Blacklist_By_Reference(t *testing.T) {
	cFTests := []struct {
		Name          string
		Reference     string
		WantErr       bool
		ExpectedModel Customer
	}{
		{
			Name:          "ok-values",
			Reference:     "refernece",
			ExpectedModel: Customer{},
		},
		{
			Name:      "error-no-reference",
			Reference: "",
			WantErr:   true,
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(createCustomerResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Customer: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			cust, resp, err := c.Customer.BlacklistByReference(context.TODO(), tt.Reference)

			if tt.WantErr {
				require.Error(t, err)

			} else {

				require.NoError(t, err)
				require.Equal(t, cust, &tt.ExpectedModel)
				require.Equal(t, resp.Request.Method, http.MethodPost)
				require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%scustomers/%s/blacklist", defaultBaseURL, tt.Reference))
			}
		})

	}
}

func Test_Customers_Whitelist_By_Reference(t *testing.T) {
	cFTests := []struct {
		Name          string
		Reference     string
		WantErr       bool
		ExpectedModel Customer
	}{
		{
			Name:          "ok-values",
			Reference:     "refernece",
			ExpectedModel: Customer{},
		},
		{
			Name:      "error-no-reference",
			Reference: "",
			WantErr:   true,
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(createCustomerResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Customer: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			cust, resp, err := c.Customer.WhitelistByReference(context.TODO(), tt.Reference)

			if tt.WantErr {
				require.Error(t, err)

			} else {

				require.NoError(t, err)
				require.Equal(t, cust, &tt.ExpectedModel)
				require.Equal(t, resp.Request.Method, http.MethodDelete)
				require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%scustomers/%s/whitelist", defaultBaseURL, tt.Reference))
			}
		})

	}
}
