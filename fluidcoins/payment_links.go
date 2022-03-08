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

// PaymentLinkService is a helper that communicates with the links API
type PaymentLinkService service

type PaymentLinkStatus string

const (
	Enabled  PaymentLinkStatus = "enabled"
	Disabled PaymentLinkStatus = "disabled"
)

type PaymentLink struct {
	ID         uuid.UUID `json:"id"`
	Identifier string    `json:"identifier"`
	MerchantID uuid.UUID `json:"merchant_id"`

	Title                 string `json:"title"`
	Description           string `json:"description"`
	Domain                Domain `json:"domain"`
	Amount                Amount `json:"amount"`
	IsEnabled             bool   `json:"is_enabled"`
	Currency              string `json:"currency"`
	HumanReadableCurrency string `json:"human_readable_currency"`

	Metadata PaymentLinkMetadata `json:"metadata"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaymentLinkMetadata struct {
	CollectPhoneNumber  bool   `json:"collect_phone_number"`
	DisableAfterPayment bool   `json:"disable_after_payment"`
	ShortLink           string `json:"short_link"`
}

type createPaymentLinkResponse struct {
	apiStatus
	PaymentLink PaymentLink `json:"link"`
}

type createPaymentLinkOptions struct {
	Currency            string `json:"currency"`
	Amount              Amount `json:"amount"`
	CollectPhoneNumber  bool   `json:"collect_phone_number"`
	DisableAfterPayment bool   `json:"disable_after_payment"`
	Title               string `json:"title"`
	Description         string `json:"description"`
}

type getPaymentLinkResponse createPaymentLinkResponse

type editPaymentLinkResponse createPaymentLinkResponse

type enablePaymentLinkResponse createPaymentLinkResponse

type disablePaymentLinkResponse createPaymentLinkResponse

type listPaymentLinkOptions struct {
	Status  PaymentLinkStatus `url:"status"`
	Perpage int64             `url:"per_page"`
	Page    int64             `url:"page"`
}

type listPaymentLinkResponse struct {
	apiStatusWithMeta
	PaymentLinks []*PaymentLink `json:"links"`
}

type listPaymentLinkTransactionsResponse struct {
	apiStatusWithMeta
	Customers []*Customer `json:"customers"`
}

// Creates a payment link
func (s *PaymentLinkService) Create(ctx context.Context, opts *createPaymentLinkOptions) (*PaymentLink, *Response, error) {

	req, err := s.c.NewRequest(http.MethodPost, "links", opts)
	if err != nil {
		return nil, nil, err
	}

	var plr = new(createPaymentLinkResponse)

	res, err := s.c.Do(ctx, req, plr)
	if err != nil {
		return nil, nil, err
	}

	return &plr.PaymentLink, res, err
}

// Fetches a payment link by reference
func (s *PaymentLinkService) GetByReference(ctx context.Context, reference string) (*PaymentLink, *Response, error) {

	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := s.c.NewRequest(http.MethodGet, fmt.Sprintf("links/%s", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var glr = new(getPaymentLinkResponse)

	res, err := s.c.Do(ctx, req, glr)
	if err != nil {
		return nil, nil, err
	}

	return &glr.PaymentLink, res, err
}

// List fetches all payment links
func (s *PaymentLinkService) List(ctx context.Context, opts *listPaymentLinkOptions) ([]*PaymentLink, *Response, error) {

	if opts.Perpage <= 0 {
		opts.Perpage = 10
	}

	if opts.Page <= 0 {
		opts.Page = 1
	}

	st, err := addQueryString("links", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.c.NewRequest(http.MethodGet, st, nil)
	if err != nil {
		return nil, nil, err
	}

	var l = new(listPaymentLinkResponse)

	res, err := s.c.Do(ctx, req, l)
	if err != nil {
		return nil, nil, err
	}

	return l.PaymentLinks, res, nil

}

// ListTransactionsByPaymentLinkReference fetches transactions on a payment link (by its reference)
func (s *PaymentLinkService) ListTransactionsByPaymentLinkReference(ctx context.Context, reference string) ([]*Customer, *Response, error) {

	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := s.c.NewRequest(http.MethodGet, fmt.Sprintf("links/%s/transactions", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var t = new(listPaymentLinkTransactionsResponse)

	res, err := s.c.Do(ctx, req, t)
	if err != nil {
		return nil, nil, err
	}

	return t.Customers, res, nil
}

// Edits a payment link by reference and supplied options
func (s *PaymentLinkService) Edit(ctx context.Context, reference string, opts *createPaymentLinkOptions) (*PaymentLink, *Response, error) {

	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := s.c.NewRequest(http.MethodPatch, fmt.Sprintf("links/%s", reference), opts)
	if err != nil {
		return nil, nil, err
	}

	var e = new(editPaymentLinkResponse)

	res, err := s.c.Do(ctx, req, e)
	if err != nil {
		return nil, nil, err
	}

	return &e.PaymentLink, res, nil

}

// Enables a payment link for collection by reference
func (s *PaymentLinkService) EnableByReference(ctx context.Context, reference string) (*PaymentLink, *Response, error) {

	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := s.c.NewRequest(http.MethodPost, fmt.Sprintf("links/%s/enable", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var e = new(enablePaymentLinkResponse)

	res, err := s.c.Do(ctx, req, e)
	if err != nil {
		return nil, nil, err
	}

	return &e.PaymentLink, res, nil
}

// Disables a payment link for collection by reference
func (s *PaymentLinkService) DisableByReference(ctx context.Context, reference string) (*PaymentLink, *Response, error) {

	if util.IsStringEmpty(reference) {
		return nil, nil, errors.New("please provide a reference")
	}

	req, err := s.c.NewRequest(http.MethodDelete, fmt.Sprintf("links/%s/enable", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	var e = new(disablePaymentLinkResponse)

	res, err := s.c.Do(ctx, req, e)
	if err != nil {
		return nil, nil, err
	}

	return &e.PaymentLink, res, nil
}
