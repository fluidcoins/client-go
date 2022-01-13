package fluidcoins

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://api.fluidcoins.com/v1/"
	defaultUA      = "Fluidcoins/Client-go"
)

// Client is the root object of cerberus. Used to interact with Paystack's API
type Client struct {
	c         *http.Client
	baseURL   string
	userAgent string
	secretKey string
}

type service struct {
	c *Client
}

type apiStatus struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// isSuccessful checks that the status of the transaction passes
func (a apiStatus) isSuccessful() bool { return a.Status }

// New creates a new instance of Cerberus
func New(opts ...Option) (*Client, error) {

	c := new(Client)

	setDefaultOptions(c)

	for _, v := range opts {
		v(c)
	}

	c.setUpServices()

	return c, c.validate()
}

func (c *Client) setUpServices() {
}

func (c *Client) validate() error {
	if len(strings.Trim(c.secretKey, " ")) == 0 {
		return errors.New("please provide your secret key")
	}

	return nil
}

// NewRequest creates a new http.Request for the current operation
func (c *Client) NewRequest(method string, url string, v interface{}) (*http.Request, error) {
	if strings.HasPrefix(url, "/") {
		return nil, errors.New("url shouldn't start with /")
	}

	var body = bytes.NewBuffer(nil)

	if v != nil {
		if err := json.NewEncoder(body).Encode(v); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.baseURL+url, body)
	if err != nil {
		return nil, err
	}

	if v != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	req.Header.Add("Authorization", "Bearer "+c.secretKey)
	req.Header.Add("User-Agent", c.userAgent)

	return req, nil
}

// Do makes the HTTP request and writes the body of the response into v
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {

	req = req.WithContext(ctx)

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	if err != nil {
		// check if context has been canceled
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		if e, ok := err.(*url.Error); ok {
			return nil, e
		}

		// Return err as is
		return nil, err
	}

	defer resp.Body.Close()

	if n := resp.StatusCode; n != http.StatusOK && n != http.StatusCreated {
		var d struct {
			Message string `json:"message"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
			return nil, err
		}

		return nil, errors.New(d.Message)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return nil, err
		}
	}

	return resp, err
}
