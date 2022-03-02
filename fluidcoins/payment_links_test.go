package fluidcoins

import (
	"context"
	"testing"

	"github.com/fluidcoins/client-go/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePaymentLink(t *testing.T) {

	opts := createPaymentLinkOptions{
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

	k, err := LoadAPIKey()
	require.NoError(t, err)

	testCases := []struct {
		name        string
		makeRequest func(*testing.T)
	}{
		{
			name: "OK",
			makeRequest: func(t *testing.T) {
				client, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: client,
				}

				plr, res, err := s.Create(context.Background(), &opts)
				require.NoError(t, err)

				require.NotEmpty(t, *res)
				require.NotEmpty(t, *plr)

				require.Equal(t, opts.Amount, plr.PaymentLink.Amount)
				require.Equal(t, opts.Title, plr.PaymentLink.Title)

				require.Equal(t, true, plr.apiStatus.Status)

			},
		},
		{
			name: "OK",
			makeRequest: func(t *testing.T) {
				client, err := New(SecretKey(k))
				require.NoError(t, err)

				s := PaymentLinkService{
					c: client,
				}

				plr, res, err := s.Create(context.Background(), &opts)
				require.NoError(t, err)

				require.NotEmpty(t, *res)
				require.NotEmpty(t, *plr)

				require.Equal(t, opts.Amount, plr.PaymentLink.Amount)

				require.Equal(t, true, plr.apiStatus.Status)

			},
		},
		{
			name: "No Authentication",
			makeRequest: func(t *testing.T) {
				_, err := New()
				require.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeRequest(t)
		})
	}

}
