package fluidcoins

import (
	"math/rand"
)

// BaseURL allows to set the base url to local server(in tests)
func BaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

func RandomDomain() Domain {
	d := []Domain{LiveDomain, TestDomain}
	n := len(d)
	return d[rand.Intn(n)]
}
