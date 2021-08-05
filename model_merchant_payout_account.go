/*
 * Fluidcoins Pay api documentation
 *
 * A suite of APIs to create, manage and accept cryptocurrency payments in Africa
 *
 * API version: 0.1.0
 * Contact: developers@fluidcoins.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MerchantPayoutAccount struct {
	Bank *MerchantBankPayoutAccount `json:"bank,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Crypto *MerchantCryptoPayoutAccount `json:"crypto,omitempty"`
	Domain string `json:"domain,omitempty"`
	Id string `json:"id,omitempty"`
	MerchantId string `json:"merchant_id,omitempty"`
	PayoutType int32 `json:"payout_type,omitempty"`
	Reference string `json:"reference,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
