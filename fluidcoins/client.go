package fluidcoins

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/fluidcoins/client-go/util"
	"github.com/google/go-querystring/query"
	"github.com/spf13/viper"
)

const (
	defaultBaseURL = "https://devapi.fluidcoins.com/v1/"
	defaultUA      = "Fluidcoins/Client-go"
)

// Domain is the mode of the response.
type Domain string

const (
	// LiveDomain is an operation that occurred on the mainnet
	LiveDomain Domain = "live"
	// TestDomain is an operation that occurred on the testnet
	TestDomain Domain = "test"
)

// Amount is an abstraction to represent the lowest unit of a (crypto)currency
type Amount int64

// Client is the root object of cerberus. Used to interact with Paystack's API
type Client struct {
	c         *http.Client
	baseURL   string
	userAgent string
	secretKey string

	Transaction *TransactionService
}

type service struct {
	c *Client
}

type apiStatus struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type APIAuth struct {
	Apikey string `mapstructure:"API_KEY"`
}

// IsSuccessful checks that the status of the transaction passes
func (a apiStatus) IsSuccessful() bool { return a.Status }

// LoadAPIKey loads generated API key from the environment variables for
// authenticating API calls
func LoadAPIKey() (string, error) {

	viper.AddConfigPath("../")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	a := new(APIAuth)
	err = viper.Unmarshal(a)
	if err != nil {
		return "", err
	}

	if util.IsStringEmpty(a.Apikey) {
		return "", errors.New("environment variable API_KEY is empty")
	}

	return a.Apikey, nil
}

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
	c.Transaction = &TransactionService{c}
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
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {

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

	return &Response{resp}, err
}

// Response embeds the standard response struct and might in the future include
// more detailed metadata about the response
type Response struct {
	*http.Response
}

func addQueryString(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
