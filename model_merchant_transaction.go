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

type MerchantTransaction struct {
	Amount int32 `json:"amount,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Customer *MerchantMerchantCustomer `json:"customer,omitempty"`
	Domain string `json:"domain,omitempty"`
	Id string `json:"id,omitempty"`
	Metadata *MerchantTransactionMetadata `json:"metadata,omitempty"`
	PaymentLinkId string `json:"payment_link_id,omitempty"`
	Reference string `json:"reference,omitempty"`
	Status string `json:"status,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
