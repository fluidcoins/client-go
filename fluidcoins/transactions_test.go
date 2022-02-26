package fluidcoins

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fluidcoins/client-go/util"
	"github.com/stretchr/testify/require"
)

func readBody(b io.ReadCloser) ([]byte, error) {
	bt, err := ioutil.ReadAll(b)
	defer b.Close()

	return bt, err
}

func randomTransaction() Transaction {
	c := randomCustomer()

	return Transaction{
		ID:            util.RandomID(),
		Amount:        Amount(util.RandomAmount()),
		Domain:        c.Domain,
		Customer:      &c,
		PaymentLinkID: util.RandomID(),
		Reference:     fmt.Sprintf("TRANS_%s", util.RandomID()),
		Status:        Pending,
	}

}

func TestTransaction_GetByID(t *testing.T) {
	tx := randomTransaction()
	require.NotEmpty(t, tx)

	testCases := []struct {
		name  string
		setup func(*testing.T)
	}{
		{
			name: "OK",
			setup: func(t *testing.T) {
				m := http.NewServeMux()
				s := httptest.NewServer(m)
				defer s.Close()

				c, _ := New(SecretKey("oo"), BaseURL(s.URL+"/"))
				ts := TransactionService{
					c: c,
				}

				m.HandleFunc(fmt.Sprintf("transactions/%s", tx.Reference), func(rw http.ResponseWriter, r *http.Request) {
					rw.Header().Set("Content-Type", "application/json")
					rw.WriteHeader(http.StatusOK)

					tr := transactionResponse{
						apiStatus: apiStatus{
							Status:  true,
							Message: "message",
						},
						Transaction: tx,
					}
					j, err := json.Marshal(&tr)
					require.NoError(t, err)
					rw.Write(j)

				})

				_, body, err := ts.GetByID(context.Background(), tx.Reference)
				require.NoError(t, err)
				data, err := readBody(body.Body)

				require.NoError(t, err)

				tx2 := new(transactionResponse)
				err = json.Unmarshal(data, &tx2)

				require.NoError(t, err)
				require.Equal(t, tx, tx2.Transaction)

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)
		})
	}
}
