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

type MerchantAddressMetadata struct {
	BlockHash string `json:"block_hash,omitempty"`
	BlockHeight int32 `json:"block_height,omitempty"`
	// The sender address. This field will not be available for all blockchains
	From string `json:"from,omitempty"`
}