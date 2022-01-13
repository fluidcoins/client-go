package fluidcoins

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOption_HTTPClient(t *testing.T) {

	h := &http.Client{
		Timeout: time.Minute * 20,
	}

	c, err := New(HTTPClient(h), SecretKey("oo"))

	require.NoError(t, err)
	require.Equal(t, h, c.c)
}

func TestOption_UserAgent(t *testing.T) {
	c, err := New(SecretKey("oo"))
	require.NoError(t, err)
	require.Equal(t, defaultUA, c.userAgent)

	s := "lamo-test"

	c, err = New(UserAgent(s), SecretKey("oo"))
	require.NoError(t, err)
	require.Equal(t, s, c.userAgent)
}

func TestOption_SecretKey(t *testing.T) {
	secret := "oo"

	c, err := New(SecretKey(secret))
	require.NoError(t, err)
	require.Equal(t, secret, c.secretKey)
}
