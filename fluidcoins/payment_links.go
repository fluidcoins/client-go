package fluidcoins

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// PaymentLinkService is a helper that communicates with the links API
type PaymentLinkService service

type PaymentLink struct {
	ID         uuid.UUID `json:"id"`
	Identifier string    `json:"identifier"`
	MerchantID uuid.UUID `json:"merchant_id"`

	Title       string              `json:"title"`
	Description string              `json:"decription"`
	Domain      Domain              `json:"domain"`
	Amount      Amount              `json:"amount"`
	IsEnabled   bool                `json:"is_enabled"`
	Metadata    PaymentLinkMetadata `json:"metadata"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaymentLinkMetadata struct {
	CollectPhoneNumber  bool `json:"collect_phone_number"`
	DisableAfterPayment bool `json:"disable_after_payment"`
}

type createPaymentLinkResponse struct {
	apiStatus
	PaymentLink PaymentLink `json:"link"`
}

type createPaymentLinkOptions struct {
	Amount              Amount `json:"amount"`
	CollectPhoneNumber  bool   `json:"collect_phone_number"`
	DisableAfterPayment bool   `json:"disable_after_payment"`
	Title               string `json:"title"`
	Description         string `json:"description"`
}

func (s *PaymentLinkService) Create(ctx context.Context, opts *createPaymentLinkOptions) (*createPaymentLinkResponse, *Response, error) {

	req, err := s.c.NewRequest(http.MethodPost, "links", opts)
	if err != nil {
		return nil, nil, err
	}

	var plr = new(createPaymentLinkResponse)

	res, err := s.c.Do(ctx, req, plr)
	if err != nil {
		return nil, nil, err
	}

	return plr, res, err
}
