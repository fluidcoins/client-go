package fluidcoins

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var nilPaymentLink *PaymentLink

func Test_Payment_Link_Failures_Create(t *testing.T) {
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
			link, resp, err := c.PaymentLink.Create(context.TODO(), &PaymentLinkCreateOptions{})

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, link, nilPaymentLink)
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Payment_Link_Failures_Edit(t *testing.T) {
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
			link, resp, err := c.PaymentLink.Edit(context.TODO(), "reference", &PaymentLinkCreateOptions{})

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, link, nilPaymentLink)
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Payment_Link_Failures_EnableByReference(t *testing.T) {
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
			link, resp, err := c.PaymentLink.EnableByReference(context.TODO(), "reference")

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, link, nilPaymentLink)
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Payment_Link_Failures_DisableByReference(t *testing.T) {
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
			link, resp, err := c.PaymentLink.DisableByReference(context.TODO(), "reference")

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, link, nilPaymentLink)
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Payment_Link_Failures_List(t *testing.T) {
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
			link, resp, err := c.PaymentLink.List(context.TODO(), &PaymentLinkListOptions{})

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, link, []*PaymentLink(nil))
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Payment_Link_Failures_ListTransactionsByPaymentLinkReference(t *testing.T) {
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
			link, resp, err := c.PaymentLink.ListTransactionsByPaymentLinkReference(context.TODO(), "reference")

			require.Error(t, err)
			require.Equal(t, err.Error(), tt.ApiErrorMessage)
			require.Equal(t, link, []*Customer(nil))
			require.Equal(t, resp, nilResp)
		})

	}
}

func Test_Payment_Link_Create(t *testing.T) {
	cFTests := []struct {
		Name          string
		ExpectedModel PaymentLink
		CreateOptions PaymentLinkCreateOptions
	}{
		{
			Name:          "ok-values",
			ExpectedModel: PaymentLink{},
			CreateOptions: PaymentLinkCreateOptions{},
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(createPaymentLinkResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Link: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			link, resp, err := c.PaymentLink.Create(context.TODO(), &tt.CreateOptions)

			require.NoError(t, err)
			require.Equal(t, link, &tt.ExpectedModel)
			require.Equal(t, resp.Request.Method, http.MethodPost)
			require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%slinks", defaultBaseURL))
		})

	}
}

func Test_Payment_Link_Edit(t *testing.T) {
	cFTests := []struct {
		Name          string
		Reference     string
		ExpectedModel PaymentLink
		EditOpts      PaymentLinkCreateOptions
		WantErr       bool
	}{
		{
			Name:          "ok-values",
			Reference:     "refernece",
			ExpectedModel: PaymentLink{},
			EditOpts:      PaymentLinkCreateOptions{},
		},
		{
			Name:      "error-no-reference",
			Reference: "",
			EditOpts:  PaymentLinkCreateOptions{},
			WantErr:   true,
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(editPaymentLinkResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Link: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			link, resp, err := c.PaymentLink.Edit(context.TODO(), tt.Reference, &tt.EditOpts)

			if tt.WantErr {
				require.Error(t, err)

			} else {

				require.NoError(t, err)
				require.Equal(t, link, &tt.ExpectedModel)
				require.Equal(t, resp.Request.Method, http.MethodPatch)
				require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%slinks/%s", defaultBaseURL, tt.Reference))
			}
		})

	}
}

func Test_Payment_Link_List(t *testing.T) {
	cFTests := []struct {
		Name           string
		WantErr        bool
		ExpectedModels []*PaymentLink
		ListOptions    *PaymentLinkListOptions
	}{
		{
			Name:           "ok-values-no-filter",
			ExpectedModels: []*PaymentLink{},
			ListOptions:    &PaymentLinkListOptions{},
		},
		{
			Name: "with-enabled-filter",
			ListOptions: &PaymentLinkListOptions{
				Status: Enabled,
			},
			ExpectedModels: []*PaymentLink{},
		},
		{
			Name: "with-disabled-filter",
			ListOptions: &PaymentLinkListOptions{
				Status:  Disabled,
				Perpage: 20,
				Page:    3,
			},
			ExpectedModels: []*PaymentLink{},
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(listPaymentLinksResponse{
				apiStatusWithMeta: apiStatusWithMeta{
					apiStatus: apiStatus{
						Status:  true,
						Message: "Successful",
					},
					Meta: ApiStatusMeta{
						Paging: PagingMeta{
							Perpage: 10,
							Page:    1,
							Total:   100,
						},
					},
				},
				Links: tt.ExpectedModels,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			links, resp, err := c.PaymentLink.List(context.TODO(), tt.ListOptions)

			require.NoError(t, err)
			require.Equal(t, links, tt.ExpectedModels)
			require.Equal(t, resp.Request.Method, http.MethodGet)
			require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%slinks?page=%d&per_page=%d&status=%s", defaultBaseURL, tt.ListOptions.Page, tt.ListOptions.Perpage, tt.ListOptions.Status))
		})

	}
}

func Test_Payment_Link_ListTransactionsByPaymentLinkReference(t *testing.T) {
	cFTests := []struct {
		Name           string
		Reference      string
		WantErr        bool
		ExpectedModels []*Customer
	}{
		{
			Name:           "ok-values",
			Reference:      "refernece",
			ExpectedModels: []*Customer{},
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
			Transport: NewRoundTripperWithResponse(listPaymentLinkTransactionsResponse{
				Customers: tt.ExpectedModels,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			cust, resp, err := c.PaymentLink.ListTransactionsByPaymentLinkReference(context.TODO(), tt.Reference)

			if tt.WantErr {
				require.Error(t, err)

			} else {

				require.NoError(t, err)
				require.Equal(t, cust, tt.ExpectedModels)
				require.Equal(t, resp.Request.Method, http.MethodGet)
				require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%slinks/%s/transactions", defaultBaseURL, tt.Reference))
			}
		})

	}
}

func Test_Payment_Link_EnableByReference(t *testing.T) {
	cFTests := []struct {
		Name          string
		Reference     string
		WantErr       bool
		ExpectedModel PaymentLink
	}{
		{
			Name:          "ok-values",
			Reference:     "refernece",
			ExpectedModel: PaymentLink{},
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
			Transport: NewRoundTripperWithResponse(enablePaymentLinkResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Link: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			cust, resp, err := c.PaymentLink.EnableByReference(context.TODO(), tt.Reference)

			if tt.WantErr {
				require.Error(t, err)

			} else {

				require.NoError(t, err)
				require.Equal(t, cust, &tt.ExpectedModel)
				require.Equal(t, resp.Request.Method, http.MethodPost)
				require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%slinks/%s/enable", defaultBaseURL, tt.Reference))
			}
		})

	}
}

func Test_Payment_Link_DisableByReference(t *testing.T) {
	cFTests := []struct {
		Name          string
		Reference     string
		WantErr       bool
		ExpectedModel PaymentLink
	}{
		{
			Name:          "ok-values",
			Reference:     "refernece",
			ExpectedModel: PaymentLink{},
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
			Transport: NewRoundTripperWithResponse(enablePaymentLinkResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Link: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			cust, resp, err := c.PaymentLink.DisableByReference(context.TODO(), tt.Reference)

			if tt.WantErr {
				require.Error(t, err)

			} else {

				require.NoError(t, err)
				require.Equal(t, cust, &tt.ExpectedModel)
				require.Equal(t, resp.Request.Method, http.MethodDelete)
				require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%slinks/%s/enable", defaultBaseURL, tt.Reference))
			}
		})

	}
}
