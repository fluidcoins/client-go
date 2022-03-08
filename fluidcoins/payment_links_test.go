package fluidcoins

import (
	"context"
	"net/http"
	"testing"

	"github.com/fluidcoins/client-go/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePaymentLink(t *testing.T) {

	opts := createPaymentLinkOptions{
		Currency:            util.RandomCurrency(),
		Amount:              Amount(util.RandomAmount()),
		CollectPhoneNumber:  util.RandomBool(),
		DisableAfterPayment: util.RandomBool(),
	}

	title, err := util.GenerateRandomString(10)
	require.NoError(t, err)

	desc, err := util.GenerateRandomString(15)
	require.NoError(t, err)

	opts.Title = title
	opts.Description = desc

	testCases := []struct {
		name        string
		makeRequest func(*testing.T)
	}{

		{
			name: "OK",
			makeRequest: func(t *testing.T) {
				k, err := LoadAPIKey()
				require.NoError(t, err)

				client, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: client,
				}

				l, res, err := s.Create(context.TODO(), &opts)

				require.Nil(t, err, "expecting nil error")

				require.NotNil(t, res, "expecting non-nil response")
				require.NotNil(t, l, "expecting non-nil payment link")

				require.NotEmpty(t, l)
				require.Equal(t, opts.Amount, l.Amount)
				require.Equal(t, opts.Title, l.Title)
				require.Equal(t, opts.Description, l.Description)
				require.Equal(t, opts.CollectPhoneNumber, l.Metadata.CollectPhoneNumber)
				require.Equal(t, opts.DisableAfterPayment, l.Metadata.DisableAfterPayment)

				require.Equal(t, http.StatusCreated, res.StatusCode)
			},
		},

		{
			name: "No Authentication Key",
			makeRequest: func(t *testing.T) {
				_, err := New()

				require.NotNil(t, err, "expecting an error")
			},
		},
		{
			name: "Expired API Key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("expiredkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.Create(context.TODO(), &opts)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},

		{
			name: "Invalid API key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("invalidkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.Create(context.TODO(), &opts)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeRequest(t)
		})
	}

}

func TestGetPaymentLinkByReference(t *testing.T) {

	reference := "LINK_AWhh5EIROwcXQ0reiMnLq"

	testCases := []struct {
		name        string
		makeRequest func(*testing.T)
	}{

		{
			name: "OK",
			makeRequest: func(t *testing.T) {
				k, err := LoadAPIKey()
				require.NoError(t, err)

				client, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: client,
				}

				l, res, err := s.GetByReference(context.TODO(), reference)

				require.Nil(t, err, "expecting nil error")

				require.NotNil(t, res, "expecting non-nil response")
				require.NotNil(t, l, "expecting non-nil payment link")

				require.NotEmpty(t, l)
				require.Equal(t, reference, l.Identifier)

				require.Equal(t, http.StatusOK, res.StatusCode)
			},
		},

		{
			name: "No Authentication Key",
			makeRequest: func(t *testing.T) {
				_, err := New()

				require.NotNil(t, err, "expecting an error")
			},
		},
		{
			name: "Expired API Key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("expiredkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.GetByReference(context.TODO(), reference)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},

		{
			name: "Invalid API key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("invalidkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.GetByReference(context.TODO(), reference)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeRequest(t)
		})
	}
}

func TestListPaymentLink(t *testing.T) {
	opts := listPaymentLinkOptions{
		Status:  Enabled,
		Perpage: 0,
		Page:    0,
	}

	testCases := []struct {
		name        string
		makeRequest func(t *testing.T)
	}{

		{
			name: "OK",
			makeRequest: func(t *testing.T) {
				k, err := LoadAPIKey()
				require.NoError(t, err)

				client, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: client,
				}

				l, res, err := s.List(context.TODO(), &opts)

				require.Nil(t, err, "expecting nil error")

				require.NotNil(t, res, "expecting non-nil response")
				require.NotNil(t, l, "expecting non-nil payment link")

				require.NotEmpty(t, true, l[0])
				require.Equal(t, true, l[0].IsEnabled)

				require.Equal(t, http.StatusOK, res.StatusCode)
			},
		},

		{
			name: "No Authentication Key",
			makeRequest: func(t *testing.T) {
				_, err := New()

				require.NotNil(t, err, "expecting an error")
			},
		},
		{
			name: "Expired API Key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("expiredkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.List(context.TODO(), &opts)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},

		{
			name: "Invalid API key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("invalidkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.List(context.TODO(), &opts)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeRequest(t)
		})
	}
}

func TestListPaymentLinkTransactions(t *testing.T) {

	reference := "LINK_AWhh5EIROwcXQ0reiMnLq"

	testCases := []struct {
		name        string
		makeRequest func(t *testing.T)
	}{

		{
			name: "OK",
			makeRequest: func(t *testing.T) {
				k, err := LoadAPIKey()
				require.NoError(t, err)

				client, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: client,
				}

				l, res, err := s.ListTransactionsByPaymentLinkReference(context.TODO(), reference)

				require.Nil(t, err, "expecting nil error")

				require.NotNil(t, res, "expecting non-nil response")

				//the used reference has no transactions yet
				require.Nil(t, l, "expecting nil payment link")

				require.Equal(t, http.StatusOK, res.StatusCode)
			},
		},

		{
			name: "No Authentication Key",
			makeRequest: func(t *testing.T) {
				_, err := New()

				require.NotNil(t, err, "expecting an error")
			},
		},
		{
			name: "Expired API Key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("expiredkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.ListTransactionsByPaymentLinkReference(context.TODO(), reference)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},

		{
			name: "Invalid API key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("invalidkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.ListTransactionsByPaymentLinkReference(context.TODO(), reference)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeRequest(t)
		})
	}

}

func TestEditPaymentLink(t *testing.T) {

	reference := "LINK_AWhh5EIROwcXQ0reiMnLq"

	opts := createPaymentLinkOptions{
		// Currency:            "USD",
		Amount:              8000,
		CollectPhoneNumber:  true,
		DisableAfterPayment: true,
		Title:               "Fried Yam",
		Description:         "Fried Stuff",
	}

	testCases := []struct {
		name        string
		makeRequest func(*testing.T)
	}{

		{
			name: "OK",
			makeRequest: func(t *testing.T) {
				k, err := LoadAPIKey()
				require.NoError(t, err)

				client, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: client,
				}

				l, res, err := s.Edit(context.TODO(), reference, &opts)

				require.Nil(t, err, "expecting nil error")

				require.NotNil(t, res, "expecting non-nil response")
				require.NotNil(t, l, "expecting non-nil payment link")

				require.NotEmpty(t, l)
				require.Equal(t, opts.Amount, l.Amount)
				require.Equal(t, opts.Title, l.Title)
				require.Equal(t, opts.Description, l.Description)
				// require.Equal(t, opts.Currency, l.Currency)
				require.Equal(t, opts.CollectPhoneNumber, l.Metadata.CollectPhoneNumber)
				require.Equal(t, opts.DisableAfterPayment, l.Metadata.DisableAfterPayment)

				require.Equal(t, http.StatusOK, res.StatusCode)
			},
		},

		{
			name: "No Authentication Key",
			makeRequest: func(t *testing.T) {
				_, err := New()

				require.NotNil(t, err, "expecting an error")
			},
		},
		{
			name: "Expired API Key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("expiredkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.Edit(context.TODO(), reference, &opts)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},

		{
			name: "Invalid API key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("invalidkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.Edit(context.TODO(), reference, &opts)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeRequest(t)
		})
	}

}

func TestEnablePaymentLink(t *testing.T) {

	reference := "LINK_AWhh5EIROwcXQ0reiMnLq"

	testCases := []struct {
		name        string
		makeRequest func(*testing.T)
	}{

		{
			name: "OK",
			makeRequest: func(t *testing.T) {
				k, err := LoadAPIKey()
				require.NoError(t, err)

				client, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: client,
				}

				l, res, err := s.EnableByReference(context.TODO(), reference)

				require.Nil(t, err, "expecting nil error")

				require.NotNil(t, res, "expecting non-nil response")
				require.NotNil(t, l, "expecting non-nil payment link")

				require.NotEmpty(t, l)
				require.Equal(t, true, l.IsEnabled)

				require.Equal(t, http.StatusOK, res.StatusCode)
			},
		},

		{
			name: "No Authentication Key",
			makeRequest: func(t *testing.T) {
				_, err := New()

				require.NotNil(t, err, "expecting an error")
			},
		},
		{
			name: "Expired API Key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("expiredkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.EnableByReference(context.TODO(), reference)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},

		{
			name: "Invalid API key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("invalidkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.EnableByReference(context.TODO(), reference)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeRequest(t)
		})
	}

}

func TestDisablePaymentLink(t *testing.T) {
	reference := "LINK_AWhh5EIROwcXQ0reiMnLq"

	testCases := []struct {
		name        string
		makeRequest func(*testing.T)
	}{

		{
			name: "OK",
			makeRequest: func(t *testing.T) {
				k, err := LoadAPIKey()
				require.NoError(t, err)

				client, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: client,
				}

				l, res, err := s.DisableByReference(context.TODO(), reference)

				require.Nil(t, err, "expecting nil error")

				require.NotNil(t, res, "expecting non-nil response")
				require.NotNil(t, l, "expecting non-nil payment link")

				require.NotEmpty(t, l)
				require.Equal(t, false, l.IsEnabled)

				require.Equal(t, http.StatusOK, res.StatusCode)
			},
		},

		{
			name: "No Authentication Key",
			makeRequest: func(t *testing.T) {
				_, err := New()

				require.NotNil(t, err, "expecting an error")
			},
		},
		{
			name: "Expired API Key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("expiredkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.DisableByReference(context.TODO(), reference)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},

		{
			name: "Invalid API key",
			makeRequest: func(t *testing.T) {

				k, err := LoadAPIKey(
					SetFileName("invalidkey"),
					SetPath("../"),
					SetExtension("env"),
				)
				require.NoError(t, err)

				c, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: c,
				}

				l, res, err := s.DisableByReference(context.TODO(), reference)

				require.Error(t, err, "expecting an error")

				require.Nil(t, res, "expecting nil response")
				require.Nil(t, l, "expecting nil payment link")

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeRequest(t)
		})
	}
}
