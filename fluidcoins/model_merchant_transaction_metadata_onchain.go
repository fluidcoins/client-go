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

type MerchantTransactionMetadataOnchain struct {
	Coin string `json:"coin,omitempty"`
	Confirmations int32 `json:"confirmations,omitempty"`
	FromAddress string `json:"from_address,omitempty"`
	Hash string `json:"hash,omitempty"`
	ToAddress string `json:"to_address,omitempty"`
}