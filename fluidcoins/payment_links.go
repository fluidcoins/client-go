package fluidcoins

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/fluidcoins/client-go/util"
	"github.com/google/uuid"
)

type PaymentLinkStatus string

const (
	Enabled  PaymentLinkStatus = "enabled"
	Disabled PaymentLinkStatus = "disabled"
)

type PaymentLinkService service

type PaymentLink struct {
	model
	Identifier string    `json:"identifier"`
	MerchantID uuid.UUID `json:"merchant_id"`
	IsEnabled  bool      `json:"is_enabled"`
	// Amount can be zero. If zero, the user has to select the amount to pay
	Amount                int64  `json:"amount"`
	Currency              string `json:"currency"`
	HumanReadableCurrency string `json:"human_readable_currency"`
	Title                 string `json:"title"`
	Description           string `json:"string"`
	Domain                string `json:"domain"`
}

type PaymentLinkMetadata struct {
	CollectPhoneNumber  bool   `json:"collect_phone_number"`
	DisableAfterPayment bool   `json:"disable_after_payment"`
	ShortLink           string `json:"short_link"`
}

type PaymentLinkListOptions struct {
	// filter results by the status of the links Can either be disabled or enabled
	Status PaymentLinkStatus `url:"status"`
	// PerPage is the number to items to return. Defaults to 10 items
	Perpage int64 `url:"per_page"`
	// Page to query data from. Defaults to 1
	Page int64 `url:"page"`
}

type PaymentLinkCreateOptions struct {
	// Amount in kobo/cents
	Amount int64 `json:"amount"`
	// CollectPhoneNumber set to true will request for the phone number of the customer
	CollectPhoneNumber bool `json:"collect_phone_number"`
	// Currency denotes the currency this payment link should be denoted in.
	// Supported: USD, NGN. Will default to NGN
	Currency string `json:"currency"`
	// Description is a longer form text describing this payment link
	Description string `json:"description"`
	// DisableAfterPayment set to true will disable this payment link after the first successful payment
	DisableAfterPayment bool `json:"disable_after_payment"`
	// Title is the name of this payment e.g Payment for XYZ
	Title string `json:"title"`
}

type listPaymentLinksResponse struct {
	apiStatusWithMeta
	Links []*PaymentLink `json:"links"`
}

type listPaymentLinkTransactionsResponse struct {
	apiStatusWithMeta
	Customers []*Customer `json:"customers"`
}

type createPaymentLinkResponse struct {
	apiStatus
	Link PaymentLink `json:"link"`
}

type editPaymentLinkResponse createPaymentLinkResponse

type enablePaymentLinkResponse createPaymentLinkResponse
type disablePaymentLinkResponse createPaymentLinkResponse

// Create a payment link
func (p *PaymentLinkService) Create(ctx context.Context, opts *PaymentLinkCreateOptions) (*PaymentLink, *Response, error) {
	req, err := p.c.NewRequest(http.MethodPost, "links", opts)

	if err != nil {
		return nil, nil, err
	}

	var link = new(createPaymentLinkResponse)

	resp, err := p.c.Do(ctx, req, &link)

	if err != nil {
		return nil, nil, err
	}

	return &link.Link, resp, nil
}

// Edit a payment link
func (p *PaymentLinkService) Edit(ctx context.Context, reference string, opts *PaymentLinkCreateOptions) (*PaymentLink, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := p.c.NewRequest(http.MethodPatch, fmt.Sprintf("links/%s", reference), opts)
	if err != nil {
		return nil, nil, err
	}

	var link = new(editPaymentLinkResponse)

	resp, err := p.c.Do(ctx, req, &link)

	if err != nil {
		return nil, nil, err
	}

	return &link.Link, resp, nil
}

// Enable a payment link for collection by reference
func (p *PaymentLinkService) EnableByReference(ctx context.Context, reference string) (*PaymentLink, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := p.c.NewRequest(http.MethodPost, fmt.Sprintf("links/%s/enable", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var link = new(enablePaymentLinkResponse)

	resp, err := p.c.Do(ctx, req, link)
	if err != nil {
		return nil, nil, err
	}

	return &link.Link, resp, nil
}

// Disable a payment link for collection by reference
func (p *PaymentLinkService) DisableByReference(ctx context.Context, reference string) (*PaymentLink, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := p.c.NewRequest(http.MethodDelete, fmt.Sprintf("links/%s/enable", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var link = new(disablePaymentLinkResponse)

	resp, err := p.c.Do(ctx, req, link)
	if err != nil {
		return nil, nil, err
	}

	return &link.Link, resp, nil
}

func (p *PaymentLinkService) List(ctx context.Context, opts *PaymentLinkListOptions) ([]*PaymentLink, *Response, error) {
	if opts.Perpage <= 0 {
		opts.Perpage = 10
	}

	if opts.Page <= 0 {
		opts.Page = 1
	}

	s, err := addQueryString("links", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := p.c.NewRequest(http.MethodGet, s, nil)
	if err != nil {
		return nil, nil, err
	}

	var links = new(listPaymentLinksResponse)

	resp, err := p.c.Do(ctx, req, &links)
	if err != nil {
		return nil, nil, err
	}

	return links.Links, resp, nil
}

// ListTransactionsByPaymentLinkReference fetches transactions on a payment link (by its reference)
func (p *PaymentLinkService) ListTransactionsByPaymentLinkReference(ctx context.Context, reference string) ([]*Customer, *Response, error) {
	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	var url string = fmt.Sprintf("links/%s/transactions", reference)

	req, err := p.c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	var trans = new(listPaymentLinkTransactionsResponse)

	resp, err := p.c.Do(ctx, req, &trans)
	if err != nil {
		return nil, nil, err
	}

	return trans.Customers, resp, nil
}
