package fluidcoins

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fluidcoins/client-go/util"
	"github.com/google/uuid"
)

type CustomerService service

// Customer is used to represent the structure of a created customer
type Customer struct {
	ID            uuid.UUID `json:"id"`
	MerchantID    uuid.UUID `json:"merchant_id"`
	FullName      string    `json:"full_name"`
	Reference     string    `json:"reference"`
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phone_number"`
	IsBlacklisted bool      `json:"is_blacklisted"`
	Domain        Domain    `json:"domain"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomerCreateOptions struct {
	// Email address of the customer
	Email string `json:"email"`
	// Customer's full name
	FullName string                     `json:"full_name"`
	Phone    CustomerCreatePhoneOptions `json:"phone"`
}

type CustomerCreatePhoneOptions struct {
	// Phone code ( e.g. +234 )
	Code string `json:"code"`
	// Remaining part of the phone number
	Phone string `json:"phone"`
}

type createCustomerResponse struct {
	apiStatus
	Customer Customer `json:"customer"`
}

type editCustomerResponse createCustomerResponse

type fetchCustomerResponse struct {
	apiStatus
	Customer Customer `json:"customer"`
}

type blacklistCustomerResponse fetchCustomerResponse
type whitelistCustomerResponse fetchCustomerResponse

// Create a new customer
func (c *CustomerService) Create(ctx context.Context, opts *CustomerCreateOptions) (*Customer, *Response, error) {
	req, err := c.c.NewRequest(http.MethodPost, "customers", opts)

	if err != nil {
		return nil, nil, err
	}

	var cust = new(createCustomerResponse)

	resp, err := c.c.Do(ctx, req, &cust)

	if err != nil {
		return nil, nil, err
	}

	return &cust.Customer, resp, nil
}

// FetchByReference fetches an Customer by the provided reference
func (c *CustomerService) FetchByReference(ctx context.Context, reference string) (*Customer, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := c.c.NewRequest(http.MethodGet, fmt.Sprintf("customers/%s", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var cust = new(fetchCustomerResponse)

	resp, err := c.c.Do(ctx, req, cust)
	if err != nil {
		return nil, nil, err
	}

	return &cust.Customer, resp, nil
}

// BlacklistByReference blacklists a Customer by reference
func (c *CustomerService) BlacklistByReference(ctx context.Context, reference string) (*Customer, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := c.c.NewRequest(http.MethodPost, fmt.Sprintf("customers/%s/blacklist", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var cust = new(blacklistCustomerResponse)

	resp, err := c.c.Do(ctx, req, cust)
	if err != nil {
		return nil, nil, err
	}

	return &cust.Customer, resp, nil
}

// WhitelistByReference whitelists a Customer by reference
func (c *CustomerService) WhitelistByReference(ctx context.Context, reference string) (*Customer, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := c.c.NewRequest(http.MethodDelete, fmt.Sprintf("customers/%s/whitelist", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var cust = new(whitelistCustomerResponse)

	resp, err := c.c.Do(ctx, req, cust)
	if err != nil {
		return nil, nil, err
	}

	return &cust.Customer, resp, nil
}

// Edit a Customer
func (c *CustomerService) Edit(ctx context.Context, reference string, opts *CustomerCreateOptions) (*Customer, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := c.c.NewRequest(http.MethodPatch, fmt.Sprintf("customers/%s", reference), opts)
	if err != nil {
		return nil, nil, err
	}

	var cust = new(editCustomerResponse)

	resp, err := c.c.Do(ctx, req, &cust)

	if err != nil {
		return nil, nil, err
	}

	return &cust.Customer, resp, nil
}
