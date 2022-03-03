package fluidcoins

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var nilPayoutAccount *PayoutAccount
var nilPayout *Payout

func Test_Payout_Account_Create(t *testing.T) {
	cFTests := []struct {
		Name          string
		ExpectedModel PayoutAccount
		CreateOptions PayoutAccountCreateOptions
	}{
		{
			Name:          "ok-values",
			ExpectedModel: PayoutAccount{},
			CreateOptions: PayoutAccountCreateOptions{},
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(createPayoutAccountResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Payout: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			payoutAccount, resp, err := c.Payout.CreatePayoutAccount(context.TODO(), &tt.CreateOptions)

			require.NoError(t, err)
			require.Equal(t, payoutAccount, &tt.ExpectedModel)
			require.Equal(t, resp.Request.Method, http.MethodPost)
			require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%spayouts/accounts", defaultBaseURL))
		})

	}
}

func Test_Payout_Request(t *testing.T) {
	cFTests := []struct {
		Name          string
		Reference     string
		ExpectedModel Payout
		Opts          RequestPayoutOptions
	}{
		{
			Name:          "ok-values",
			Reference:     "refernece",
			ExpectedModel: Payout{},
			Opts:          RequestPayoutOptions{},
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(requestPayoutResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Payout: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			link, resp, err := c.Payout.RequestPayout(context.TODO(), &tt.Opts)

			require.NoError(t, err)
			require.Equal(t, link, &tt.ExpectedModel)
			require.Equal(t, resp.Request.Method, http.MethodPost)
			require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%spayouts", defaultBaseURL))
		})

	}
}

func Test_Payout_PayoutDetailsByReference(t *testing.T) {
	cFTests := []struct {
		Name          string
		Reference     string
		WantErr       bool
		ExpectedModel Payout
	}{
		{
			Name:          "ok-values-no-filter",
			ExpectedModel: Payout{},
			Reference:     "reference",
		},
		{
			Name:    "error-no-reference",
			WantErr: true,
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(fetchPayoutResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Payout: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			payout, resp, err := c.Payout.PayoutDetailsByReference(context.TODO(), tt.Reference)

			if tt.WantErr {
				require.Error(t, err)

			} else {
				require.NoError(t, err)
				require.Equal(t, payout, &tt.ExpectedModel)
				require.Equal(t, resp.Request.Method, http.MethodGet)
				require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%spayouts/%s", defaultBaseURL, tt.Reference))
			}
		})

	}
}

func Test_Payout_Cancel(t *testing.T) {
	cFTests := []struct {
		Name          string
		Reference     string
		WantErr       bool
		ExpectedModel Payout
	}{
		{
			Name:          "ok-values-no-filter",
			ExpectedModel: Payout{},
			Reference:     "reference",
		},
		{
			Name:    "error-no-reference",
			WantErr: true,
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(fetchPayoutResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Payout: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			payout, resp, err := c.Payout.Cancel(context.TODO(), tt.Reference)

			if tt.WantErr {
				require.Error(t, err)

			} else {
				require.NoError(t, err)
				require.Equal(t, payout, &tt.ExpectedModel)
				require.Equal(t, resp.Request.Method, http.MethodDelete)
				require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%spayouts/%s", defaultBaseURL, tt.Reference))
			}
		})

	}
}

func Test_Payout_ResolveBankAccount(t *testing.T) {
	cFTests := []struct {
		Name          string
		ExpectedModel ResolveBankInfo
		ExpectedUrl   string
		Opts          *ResolveBankInfoOptions
	}{
		{
			Name:          "ok-values-no-filter",
			ExpectedModel: ResolveBankInfo{},
			Opts: &ResolveBankInfoOptions{
				BankCode:      "12344",
				AccountNumber: "92923939393",
			},
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(resolveBankInfoResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Account: tt.ExpectedModel,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			payout, resp, err := c.Payout.ResolveBankAccount(context.TODO(), tt.Opts)

			require.NoError(t, err)
			require.Equal(t, payout, &tt.ExpectedModel)
			require.Equal(t, resp.Request.Method, http.MethodGet)
			require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%spayouts/accounts/banks/resolve?account=%s&bank_code=%s", defaultBaseURL, tt.Opts.AccountNumber, tt.Opts.BankCode))
		})

	}
}

func Test_Payout_ListBanks(t *testing.T) {
	cFTests := []struct {
		Name           string
		ExpectedModels []*Bank
		Opts           *ListBankOptions
	}{
		{
			Name:           "ok-values",
			ExpectedModels: []*Bank{},
		},
	}

	for _, tt := range cFTests {
		h := &http.Client{
			Timeout: time.Second * 5,
			Transport: NewRoundTripperWithResponse(listBanksResponse{
				apiStatus: apiStatus{
					Status:  true,
					Message: "Successful",
				},
				Banks: tt.ExpectedModels,
			}),
		}
		t.Run(tt.Name, func(t *testing.T) {
			c, _ := New(HTTPClient(h), SecretKey("oo"))
			banks, resp, err := c.Payout.ListBanks(context.TODO(), tt.Opts)

			require.NoError(t, err)
			require.Equal(t, banks, tt.ExpectedModels)
			require.Equal(t, resp.Request.Method, http.MethodGet)
			require.Equal(t, resp.Request.URL.String(), fmt.Sprintf("%spayouts/accounts/banks", defaultBaseURL))
		})

	}
}
