/*
 * Fluidcoins Pay api documentation
 *
 * A suite of APIs to create, manage and accept cryptocurrency payments in Africa
 *
 * API version: 0.1.5
 * Contact: developers@fluidcoins.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MerchantTransactionMetadataPayment struct {
	// Amount is what the user actually paid
	Amount int32 `json:"amount,omitempty"`
	// The coin the user used to make payment
	Coin int32 `json:"coin,omitempty"`
	HumanReadableAmount float64 `json:"human_readable_amount,omitempty"`
	RefundAmount int32 `json:"refund_amount,omitempty"`
	RefundStatus int32 `json:"refund_status,omitempty"`
}