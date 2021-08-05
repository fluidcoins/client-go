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

type MerchantAddress struct {
	Address string `json:"address,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Domain string `json:"domain,omitempty"`
	Id string `json:"id,omitempty"`
	MerchantId string `json:"merchant_id,omitempty"`
	Metadata *MerchantCoinMetadata `json:"metadata,omitempty"`
	// Refernece is the unique ID for this address. You can use this id to fetch the address again
	Reference string `json:"reference,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
