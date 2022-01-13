package fluidcoins

import (
	"net/http"
	"time"
)

// Option is a configuration type
type Option func(*Client)

// HTTPClient is an Option type that allows you provide your own HTTP client
func HTTPClient(h *http.Client) Option {
	return func(c *Client) {
		c.c = h
	}
}

// UserAgent is a custom user agent for Cerberus when making requests to
// the Merchants API
func UserAgent(s string) Option {
	return func(c *Client) {
		c.userAgent = s
	}
}

// SecretKey allows you set the key required to communicate with the API
func SecretKey(s string) Option {
	return func(c *Client) {
		c.secretKey = s
	}
}

func setDefaultOptions(c *Client) {
	c.c = &http.Client{
		Timeout: time.Second * 15,
	}

	c.baseURL = defaultBaseURL
	c.userAgent = defaultUA
}
